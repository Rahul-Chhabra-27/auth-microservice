package config

import (
	"auth-micro/model"
	"fmt"
	"os"

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

func ConnectDB() *gorm.DB {
	userdb, err := gorm.Open(mysql.Open(DatabaseDsn()), &gorm.Config{});

	if err != nil {
		panic("Failed to connect DB");
	}
	userdb.AutoMigrate(&model.User{});
	return userdb
}

func ValidatingFieldsOfUser(user model.User) bool {
	// if (user.Name != "" && user.Email.contain("@"))
	return true;
}