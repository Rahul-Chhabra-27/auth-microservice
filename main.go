package main

import (
	"auth-micro/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger *zap.Logger
var userDbConnector *gorm.DB
func init() {
	var err error;
	logger, err = zap.NewDevelopment();

	if err != nil {
		panic(err)
	}
	defer logger.Sync();
}
// entry point
func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("No .env file found")
	}
	userDbConnector = config.ConnectDB();

	// configuration of the http server.
	httpServer := gin.Default();
	//? Method : @POST
	// ? Endpoint Route : /save-user
	httpServer.POST("/save-user",AddUser)

	// running the server
	httpServer.Run(":8080");  // Infinite loop
	
}