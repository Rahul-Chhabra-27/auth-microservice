package config

import (
	"auth-micro/model"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
}
// ? password -> abcd --> generateEncryptedPassword(password(1234)) --> ddbhcbdhcbdh3232 , comparePassword(plainOldPassword, encrypted password) 

func GenerateHashedPassword(password string) string {
	hashedPassword, err :=  bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost);
	if err != nil {
		log.Fatalf("Couldn't hash password and the error is %s", err);
	}
	return string(hashedPassword)
}
func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ConnectDB() *gorm.DB {
	userdb, err := gorm.Open(mysql.Open(DatabaseDsn()), &gorm.Config{});

	if err != nil {
		panic("Failed to connect DB");
	}
	userdb.AutoMigrate(&model.User{});
	return userdb
}
func ValidatePhone(phone string) bool {
	if len(phone) != 10 {
		return false
	}
	for _, digit := range phone {
		if digit < '0' && digit > '9' {
			return false
		}
	}
	return true
}
func ValidatingFieldsOfUser(user model.User) bool {
	if (user.Name != "" || user.Password != "" || user.Address != "" || user.City != "") {
		return false
	}
	if (!strings.Contains(user.Email,"@") || !strings.Contains(user.Email,".")) {
		return false
	}
	return len(user.Password) >= 6 && ValidatePhone(user.Phone);
}
// func ValidatingFieldsOfUser(user model.User) bool { 
// 		user.Name == "" { 		return true	} 	)	
// 		phonePattern := `^\d{10,15}$`	
// 		matched, _ := regexp.MatchString(phonePattern, user.Name) 	
// 		if matched { 		return false	}
// func ValidatingFieldsOfUser(user model.User) bool { 	
// 	if user.Name == "" { 		return false	} 	
// 	if user.Password == "" || len(user.Password) < 8 { 		return false	} 	
// 	if user.Email == "" || !contains(user.Email, '@') || !contains(user.Email, '.') { 		return false	} 	
// 	if user.Phone == "" || !isNumeric(user.Phone) || len(user.Phone) < 10 || len(user.Phone) > 15 { 		return false	} 
// 	if user.Address == "" { 		return false	} 	
// 	if user.City == "" { 		return false	} 	return true}
// func ValidatingFieldsOfUser(user model.User) bool string {
// 	if len(user.Name) < 5 {
// 		return false, "Name must be at least 5 characters long"
// 	}
// 	if isNumeric(user.Name) {
// 		return false, "Name must not be a numeric value"
// 	}
// 	if len(user.Password) < 8 {
// 		return false, "Password must be at least 8 characters long"
// 	}
 
// 	if user.Address == "" {
// 		return false, "Address is required"
// 	}
 
// 	if user.City == "" {
// 		return false, "City is required"
// 	}
 
	
// 	return true, "Validation successful"
// } 