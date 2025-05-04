package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/vkhangstack/dlt/internal/logger"
	"os"
)

//	type DB struct {
//		db    *gorm.DB
//		cache *cache.RedisCache
//	}

// new database
//func NewDB(db *gorm.DB, cache *cache.RedisCache) *DB {
//	return &DB{
//		db:    db,
//		cache: cache,
//	}
//}

type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}

func ConnectDatabase() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	logger.SetupLogger()

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", conn)
	db.LogMode(gin.Mode() == "debug")

	return db, err
}
