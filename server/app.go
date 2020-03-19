package server

import (
	"context"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"edittapi/publication"
	pubhttp "edittapi/publication/delivery/http"
	pubmongo "edittapi/publication/repository/mongo"
	pubuc "edittapi/publication/usecase"
)

type App struct {
	httpServer *http.Server

	publicationUseCase publication.UseCase
}

func NewApp() *App {
	db := initDB()

	publicationRepo := pubmongo.NewPublicationRepository(db, viper.GetString("mongo.publications_collection"))
	publicationUseCase := pubuc.NewPublicationUseCase(publicationRepo)

	return &App{
		publicationUseCase: publicationUseCase,
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()

	// Config and add CORS. Configuration can be found in privacyapi.yml. Endpoint is specific to env.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:8080"}
	corsConfig.AllowMethods = []string{"*"}
	corsConfig.AllowHeaders = []string{"Origin", "X-Client-Version", "X-User-Identity", "X-Mode", "Idempotency-Key", "Authorization", "Content-Type", "Accept", "Referer", "User-Agent", "Access-Control-Allow-Origin", "Accept-Version"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour
	corsConfig.AllowBrowserExtensions = true

	rateLimiterMiddleware := limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(100*time.Millisecond), 10), time.Hour
	}, func(c *gin.Context) {
		c.AbortWithStatus(429)
	})

	router.Use(
		cors.New(corsConfig),
		gin.Recovery(),
		gin.Logger(),
		rateLimiterMiddleware,
	)

	// API endpoints
	api := router.Group("/api")
	pubhttp.RegisterHTTPEndpoints(api, a.publicationUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}
