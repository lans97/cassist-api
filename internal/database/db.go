package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lans97/cassist-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
}

var DB *gorm.DB

func InitDB() {
	config := getDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Mexico_City",
        config.host,
        config.user,
        config.password,
        config.dbname,
        config.port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Could not open database conection: %v", err)
	}

    DB.AutoMigrate(&models.User{}, &models.User{}, &models.User{})
}

func getDBConfig() DBConfig {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
	return DBConfig{
        host,
        user,
        password,
        dbname,
        port,
    }
}

func Pagination(c echo.Context) func(db *gorm.DB) *gorm.DB {
    return func (db *gorm.DB) *gorm.DB {
        limit, err := strconv.Atoi(c.QueryParam("limit"))

        if err != nil {
            log.Errorf("Error converting 'limit' query parameter: %v", err)
            return db
        }
        page, err := strconv.Atoi(c.QueryParam("page"))
        if err != nil {
            log.Errorf("Error converting 'page' query parameter: %v", err)
            return db
        }

        offset := page*limit

        return db.Limit(limit).Offset(offset)
    }
}

