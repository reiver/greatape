package db

import (
	"contracts"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlStorage struct {
	// Underlying database connection
	database *gorm.DB
	dsn      string
}

func NewMySQLStorage() contracts.IStorage {
	return &mysqlStorage{}
}

// Connect initiate the database connection and migrate all the tables
func (storage *mysqlStorage) Connect(dsn string) {
	storage.dsn = dsn

	database, err := gorm.Open(mysql.Open(storage.dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	storage.database = database
	Executor = storage.database

	log.Println("[DATABASE]::CONNECTED")
}

// Migrate migrates all the database tables
func (storage *mysqlStorage) Migrate(tables ...interface{}) error {
	return storage.database.AutoMigrate(tables...)
}

func (storage *mysqlStorage) Prepare(string) contracts.IQuery {
	return nil
}
