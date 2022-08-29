package main

import (
	"app/models/repos"
	"app/routes"
	"caching"
	"config"
	"db"
	"fmt"
	"logging"
	"server"
)

func main() {
	logger := logging.CreateLogger(logging.StdIOLogger)
	cache := caching.CreateCache(caching.InProcCache)

	storage := db.CreateStorage(db.SqliteStorage)
	storage.Connect(config.SQLITE_DB)
	storage.Migrate(repos.All...)

	app := server.New()
	app.SetStorage(storage)
	app.SetLogger(logger)
	app.SetCache(cache)
	app.Bind(routes.All...)

	app.Listen(fmt.Sprintf(":%s", config.PORT))
}
