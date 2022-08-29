package db

import (
	"contracts"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresqlStorage struct {
	// Underlying database connection
	database *gorm.DB
	dsn      string
}

func NewPostgreSQLStorage() contracts.IStorage {
	return &postgresqlStorage{}
}

// Connect initiate the database connection and migrate all the tables
func (storage *postgresqlStorage) Connect(dsn string) {
	storage.dsn = dsn

	database, err := gorm.Open(postgres.Open(storage.dsn), &gorm.Config{
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
func (storage *postgresqlStorage) Migrate(tables ...interface{}) error {
	return storage.database.AutoMigrate(tables...)
}

func (storage *postgresqlStorage) Prepare(string) contracts.IQuery {
	return nil
}
