package main

import (
	"auth-micro/config"
	"auth-micro/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthenticateUser(ctx *gin.Context) {
	var user model.User;
 	ctx.ShouldBindJSON(&user);

	userEmail := user.Email
	userPassword := user.Password

	logger.Info("Recieved Authenticate User Request", zap.String("user email", userEmail));

	// ? validation logic
	if userEmail == "" || userPassword == "" || !strings.Contains(userEmail,"@") || !strings.Contains(userEmail,".") || len(userPassword) < 6 {
		logger.Warn("Invalid Request", zap.String("user email", userEmail));
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"The request contains missing or invalid fields"})
		return;
	} 

	// ? After fields are validated
	 var existingUser model.User;
	//   userNotFoundError := userDbConnector.Where("email = ?,", user.Email).First(&existingUser).Error;
	userNotFoundError := userDbConnector.Where("email = ?", user.Email).First(&existingUser).Error;
		// ? If the user is already exist -> userNotFoundError = nil
		// ? If the user does not exist -> userNotFoundError = error

	if userNotFoundError != nil {
		logger.Warn("Authentication failed", zap.Error(userNotFoundError))
		ctx.JSON(http.StatusNotFound, gin.H{"message":"User not found"})
		return;
	}
	// ? config.ComparePassword(existingUser.Password, userPassword) is nil success passwords are matching. 
	// ? config.ComparePassword(existingUser.Password, userPassword) is not nil failed passwords are not matching.
	//? compare the passwords
	if config.ComparePassword(existingUser.Password, userPassword) != nil {
		// ? success -> Passwwords are matching
		logger.Warn("Authentication failed due to wrong password");
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"Authentication failed due to wrong password"})
		return;
	}
	logger.Info("User Authenticated Successfully", zap.String("username", existingUser.Name), zap.String("useremail", userEmail) );
	ctx.JSON(http.StatusOK, gin.H{"message":"User Authenticated Successfully"})
}