package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	// DefaultMaxOpenConns - max open connections for database
	DefaultMaxOpenConns = 20
	// DefaultMaxIdleConns - max idle connections for database
	DefaultMaxIdleConns = 5
)

// Handler for the db operations
type Handler struct {
	db *gorm.DB
}

type Config struct {
	ServerAddr string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPasswd   string
	DBName     string
}

func New() *Handler {
	settings := Config{
		ServerAddr: ":8080",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPasswd:   "docker",
		DBName:     "postgres",
	}

	if value := os.Getenv("SERVER_ADDR"); value != "" {
		settings.ServerAddr = value
	}
	if value := os.Getenv("DB_HOST"); value != "" {
		settings.DBHost = value
	}
	if value := os.Getenv("DB_PORT"); value != "" {
		settings.DBPort = value
	}
	if value := os.Getenv("DB_USER"); value != "" {
		settings.DBUser = value
	}
	if value := os.Getenv("DB_PASSWD"); value != "" {
		settings.DBPasswd = value
	}
	if value := os.Getenv("DB_NAME"); value != "" {
		settings.DBName = value
	}
	handler, err := NewHandler(settings)
	if err != nil {
		panic("Error connecting to DB " + err.Error())
	}
	return handler
}

func NewHandler(settings Config) (*Handler, error) {
	s := &Handler{}

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings.DBHost,
		settings.DBPort,
		settings.DBUser,
		settings.DBPasswd,
		settings.DBName,
	)

	log.Println("Connecting to DB: " + databaseURL)

	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}

	log.Println("DB connected")
	db.LogMode(false)
	db.DB().SetMaxOpenConns(DefaultMaxOpenConns)
	db.DB().SetMaxIdleConns(DefaultMaxIdleConns)
	if err := db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}
	s.db = db
	return s, nil
}
