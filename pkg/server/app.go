package server

import (
	"context"
	"edittapi/pkg/admin"
	adminhttp "edittapi/pkg/admin/delivery"
	"edittapi/pkg/admin/delivery/auth"
	adminuc "edittapi/pkg/admin/usecase"
	"edittapi/pkg/feedback"
	fbhttp "edittapi/pkg/feedback/delivery/http"
	fbmongo "edittapi/pkg/feedback/repo/mongo"
	fbuc "edittapi/pkg/feedback/usecase"
	"edittapi/pkg/metrics/collector"
	metricsmgo "edittapi/pkg/metrics/repository/mongo"
	metricsuc "edittapi/pkg/metrics/usecase"
	"edittapi/pkg/publication/upload"
	"edittapi/pkg/publication/usecase"
	"edittapi/sidecar/filestorage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	limit "github.com/yangxikun/gin-limit-by-key"
	"github.com/zhashkevych/scheduler"
	"golang.org/x/time/rate"

	_ "edittapi/docs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"edittapi/pkg/publication"
	pubhttp "edittapi/pkg/publication/delivery/http"
	pubmongo "edittapi/pkg/publication/repository/mongo"
)

// *** SWAGGER COMMENTS ***

// @title editt API
// @version 0.1
// @description editt back-end API

// @BasePath /

type App struct {
	httpServer *http.Server

	publicationUseCase publication.UseCase
	adminUseCase       admin.UseCase
	metricsCollector   *collector.MetricsCollector

	fileStorage   *filestorage.FileStorage
	imageUploader publication.Uploader

	feedbackUseCase feedback.UseCase
}

func NewApp(accessKey, secretKey, env string) *App {
	db := initDB()

	publicationRepo := pubmongo.NewPublicationRepository(db, viper.GetString("mongo.publications_collection"))
	publicationUseCase := usecase.NewPublicationUseCase(publicationRepo)

	metricsRepo := metricsmgo.NewMetricsRepository(db, viper.GetString("mongo.metrics_collection"))
	metricsUseCase := metricsuc.NewMetricsUseCase(metricsRepo, publicationUseCase)
	metricsCollector := collector.NewMetricsCollector(metricsUseCase)

	feedbackRepo := fbmongo.NewFeedbackRepository(db, viper.GetString("mongo.feedback_collection"))
	feedbackUseCase := fbuc.NewFeedbackUseCase(feedbackRepo)

	adminUseCase := adminuc.NewAdminUseCase(metricsUseCase, publicationUseCase, feedbackUseCase)

	// Initiate an S3 compatible client
	client, err := minio.New(viper.GetString("storage.endpoint"), accessKey, secretKey, false)
	if err != nil {
		log.Fatal(err)
	}
	fileStorage := filestorage.NewFileStorage(client, viper.GetString("storage.bucket"), viper.GetString("storage.endpoint"), env)

	return &App{
		publicationUseCase: publicationUseCase,
		adminUseCase:       adminUseCase,
		metricsCollector:   metricsCollector,
		fileStorage:        fileStorage,
		imageUploader:      upload.NewUploader(fileStorage),
		feedbackUseCase:    feedbackUseCase,
	}
}

func (a *App) Run(port string) error {
	// Init scheduler
	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, a.metricsCollector.Flush, viper.GetDuration("metrics.interval")*time.Minute)

	// Init gin handler
	router := gin.Default()

	// Config and add CORS. Configuration can be found in privacyapi.yml. Endpoint is specific to env.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:8080", "http://localhost", "http://editt.network"}
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

	// HTTP Routes
	router.Use(
		cors.New(corsConfig),
		gin.Recovery(),
		gin.Logger(),
		rateLimiterMiddleware,
		a.metricsCollector.Middleware,
	)

	// swagger router
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API endpoints
	api := router.Group("/api")
	pubhttp.RegisterHTTPEndpoints(api, a.publicationUseCase, a.imageUploader)
	fbhttp.RegisterHTTPHandlers(api, a.feedbackUseCase)

	// Admin Panel Endpoints
	admin := router.Group("/admin")
	authorizer := auth.NewAuthorizer(
		viper.GetString("authorizer.username"),
		viper.GetString("authorizer.password_hash"),
		viper.GetString("authorizer.hash_salt"),
		[]byte(viper.GetString("authorizer.signing_key")),
		viper.GetDuration("authorizer.expire_duration"),
	)
	adminhttp.RegisterHTTPEndpoints(admin, a.adminUseCase, authorizer)

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

	worker.Stop()

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
