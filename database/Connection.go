package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err = godotenv.Load()
	if err != nil {
		fmt.Printf("godotenv.Load error: %s\n", err)
	}

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWD"),
		os.Getenv("MYSQL_ADDR"),
		os.Getenv("MYSQL_DB"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("gorm.Open error: %s\n", err)
	}
}
