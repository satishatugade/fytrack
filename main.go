package main

import (
	"fmt"
	"fytrack/config"
	"fytrack/controller"
	"fytrack/entity"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	logDir := os.Getenv("LOG_DIRECTORY")
	logFile := os.Getenv("LOG_FILE")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Error creating logs directory: %v", err)
	}
	logPath := logDir + logFile
	file, LogFileError := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if LogFileError != nil {
		log.Fatal("Error opening log file: ", LogFileError)
	}
	defer file.Close()
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(file)
	log.Println("FYTRACK Application Started.............................")

	KeycloakClientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	if KeycloakClientSecret != "" {
		fmt.Println("keycloak client")
	}
	config.InitKeycloak()
	config.ConnectDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
	}))
	router.GET("/", func(ctx *gin.Context) {
		entity.RespondSuccess(ctx, "Fytrack server running...", nil)
	})
	member := router.Group("/v1/member")
	{
		member.POST("/add-member", controller.AddMemberData)
		member.POST("/get-member-info", controller.GetMemberInfo)
		member.PUT("/edit-member", controller.UpdateMemberInfo)
		member.POST("/delete-member", controller.DeleteMemberInfo)
	}
	router.Run(":" + os.Getenv("GO_SERVER_PORT"))

}