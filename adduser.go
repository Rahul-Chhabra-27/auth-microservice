package main

import (
	"auth-micro/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AddUser(ctx *gin.Context) {
  var user model.User;
  ctx.ShouldBindJSON(&user);

  logger.Info("Recieved User Request", zap.String("useremail", user.Email), zap.String("username", user.Name));

  // 1. write the validation logic.
  // 2. Check for Existing User.
  var existingUser model.User;
//   userNotFoundError := userDbConnector.Where("email = ?,", user.Email).First(&existingUser).Error;
 userNotFoundError := userDbConnector.Where("email = ?", user.Email).First(&existingUser).Error;
//  fmt.Println(userNotFoundError);
  // user doesn't exist
  if userNotFoundError == gorm.ErrRecordNotFound {
	// 3. Creating a hashed password
	hashedPassword := user.Password;

	newUser := &model.User{ Name: user.Name, Email: user.Email, Password: hashedPassword, Address: user.Address, City: user.City, Phone: user.Phone }

	primaryKey := userDbConnector.Create(newUser);

	if primaryKey.Error != nil {
		logger.Error("Failed to Create user", zap.String("userPhone ", user.Phone), zap.Error(primaryKey.Error))
		ctx.JSON(http.StatusConflict, gin.H{ "message" : "The Phone is already registered"})
		return
	}
	logger.Info(fmt.Sprintf("User %s created successfully", user.Name));
	ctx.JSON(http.StatusCreated, gin.H{ "message" : "User created successfully" })
  } else {
	logger.Warn("User Email Already Exist", zap.String("usermail", user.Email))
	ctx.JSON(http.StatusConflict, gin.H{"message" : "User Email Already Exist"})
  }
}
