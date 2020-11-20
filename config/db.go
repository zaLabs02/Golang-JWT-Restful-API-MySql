package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Database is the object uses by the models for accessing
// database tables and executing queries.
var Database *gorm.DB

func init() {
	var err error
	godotenv.Load(".env")
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	Database, err = gorm.Open(os.Getenv("DB_DRIVER"), DBURL)

	if err != nil {
		panic(err)
	}

	// set this to 'true' to see sql logs
	Database.LogMode(false)

	// fmt.Println("Database connection successful.")
}
