package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	//_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName))

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}
