package main

import (
	"app/models/repos"
	"app/routes"
	"config"
	"db"
	"fmt"
	"logging"
	"server"
)

func main() {
	logger := logging.CreateLogger(logging.StdIOLogger)

	storage := db.CreateStorage(db.SqliteStorage)
	storage.Connect(config.SQLITE_DB)
	storage.Migrate(repos.All...)

	app := server.New()
	app.SetStorageProvider(storage)
	app.SetLogger(logger)
	app.Bind(routes.All...)

	app.Listen(fmt.Sprintf(":%s", config.PORT))
}
