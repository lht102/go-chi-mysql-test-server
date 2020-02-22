package common

import (
    "fmt"
    "os"
    "log"

    "github.com/joho/godotenv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    dbDriver := os.Getenv("DB_DRIVER")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbPort := os.Getenv("DB_PORT")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
    db, err := sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatal("Connection fail")
    }
    DB = db
}
