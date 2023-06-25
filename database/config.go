package database

import (
	"fmt"

	"github.com/sudrajad02/ecommerce/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConncetionDatabase() {
	envFile := godotenv.Load()

	if envFile != nil {
		fmt.Print(envFile)
	}

	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	// dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open("admin:password@tcp(localhost:3306)/ecommerce_db"))

	if err != nil {
		panic(err)
	}
	// println(db)

	db.AutoMigrate(&models.Account{})

	DB = db

}
