package db

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

    config := mysql.Config{
        User:            username,
        Passwd:          password,
        Net:             "tcp",
        Addr:            host + ":" + port,
        DBName:          name,
        ParseTime:       true,                
        AllowNativePasswords: true,            
    }

    db, err := sqlx.Open("mysql", config.FormatDSN())
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    DB = db
}
