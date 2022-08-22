package tests

import (
	"app/models/repos"
	"app/routes"
	"config"
	"db"
	"fmt"
	"logging"
	"os"
	"server"
	"testing"
)

const Root = "http://localhost"

func TestMain(m *testing.M) {
	logger := logging.CreateLogger(logging.StdIOLogger)

	storage := db.CreateStorage(db.SqliteStorage)
	storage.Connect(config.SQLITE_DB)
	storage.Migrate(repos.All...)

	app := server.New()
	app.SetStorageProvider(storage)
	app.SetLogger(logger)
	app.Bind(routes.All...)

	go func() {
		app.Listen(fmt.Sprintf(":%s", config.PORT))
	}()

	os.Exit(m.Run())
}
