package db

import (
	"contracts"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type sqliteStorage struct {
	// Underlying database connection
	database *gorm.DB
	path     string
}

func NewSqliteStorage() contracts.IStorage {
	return &sqliteStorage{}
}

// Connect initiate the database connection and migrate all the tables
func (storage *sqliteStorage) Connect(path string) {
	storage.path = path

	database, err := gorm.Open(sqlite.Open(storage.path), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	storage.database = database
	DB = storage.database

	log.Println("[DATABASE]::CONNECTED")
}

// Migrate migrates all the database tables
func (storage *sqliteStorage) Migrate(tables ...interface{}) error {
	return storage.database.AutoMigrate(tables...)
}

func (storage *sqliteStorage) Prepare(string) contracts.IQuery {
	return nil
}
