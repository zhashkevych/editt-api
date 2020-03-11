package server

import (
	"context"
	"edittapi/application/profile/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"edittapi/auth"

	authhttp "edittapi/auth/delivery/http"
	authmongo "edittapi/auth/repository/mongo"
	authusecase "edittapi/auth/usecase"

	profilemongo "edittapi/application/profile/repository/mongo"
	profilehttp "edittapi/application/profile/delivery/http"

	"edittapi/application/feed"
	"edittapi/application/profile"
	"edittapi/application/publication"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase

	profileUC     profile.UseCase
	publicationUC publication.PublicationUseCase
	commentUC     publication.CommentUseCase
	likeUC        publication.LikeUseCase
	feedUC        feed.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	profileRepo := profilemongo.NewProfileRepository(db, viper.GetString("mongo.profile_collection"))

	profileUC := usecase.NewProfileUseCase(profileRepo)
	authUC := authusecase.NewAuthUseCase(
		userRepo,
		profileUC,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl"),
	)

	return &App{
		profileUC: profileUC,
		authUC:    authUC,
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	profilehttp.RegisterHTTPEndpoints(api, a.profileUC)

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
