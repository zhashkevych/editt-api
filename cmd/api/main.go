package main

import (
	"edittapi/pkg/config"
	"edittapi/pkg/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const ENV_PROD = "prod"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)


	if os.Getenv("HOST") == ENV_PROD {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	app := server.NewApp(accessKey, secretKey)

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
