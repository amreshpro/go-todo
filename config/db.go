package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "strconv"
    "time"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("SSL_MODE"),
    )

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to open DB:", err)
    }

    // Connection pooling
    maxOpenConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
    maxIdleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
    connMaxLifetimeMin, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME_MIN"))

    DB.SetMaxOpenConns(maxOpenConns)                                           // Max total connections
    DB.SetMaxIdleConns(maxIdleConns)                                           // Idle connections
    DB.SetConnMaxLifetime(time.Minute * time.Duration(connMaxLifetimeMin))     // Lifetime of each conn

    // Ping to test connection
    if err := DB.Ping(); err != nil {
        log.Fatal("Failed to ping DB:", err)
    }

    log.Println("✅ Connected to PostgreSQL with connection pooling!")
}
