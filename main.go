package main

import (
	"auth-micro/config"
	"auth-micro/jwt"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger *zap.Logger
var userDbConnector *gorm.DB
var JwtManager *jwt.JWTManager
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

	// Create a new jwt manager
	JwtManager =  jwt.NewJWTManager(os.Getenv("SECRET_KEY"), 5 * time.Hour);


	// configuration of the http server.
	httpServer := gin.Default();
	
	//? Method : @POST
	// ? Endpoint Route : /save-user
	httpServer.POST("/save-user",AddUser)
	httpServer.POST("/authenticate-user",AuthenticateUser)

    httpServer.Use(jwt.AuthorizeJwtToken());
	httpServer.GET("/hello", func(ctx *gin.Context) {
		fmt.Println(ctx.GetString("usermail"))
		ctx.JSON(http.StatusOK, gin.H{"message":"hello"});
	})
	// running the server
	httpServer.Run(":8080");  // Infinite loop
	
}