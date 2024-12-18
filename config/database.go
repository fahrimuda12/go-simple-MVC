package config

import (
	"go-simple-MVC/app/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init databse tapi konfigurasi menggukanan env
func DBInit() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	// dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	

	// dsn := dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	dsn := "host="+dbHost+" user="+dbUserName+" password="+dbPassword+" dbname="+dbName+" port="+dbPort+" sslmode=disable TimeZone=Asia/Jakarta"
	print(dsn)
	// db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open(dbConnection, dbUser+":"+dbPass+"@tcp("+dbHost+":"+port+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(models.Wallets{})
	return db
}