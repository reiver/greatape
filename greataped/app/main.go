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

// @title GreatApe API
// @version 1.0
// @description GreatApe is a free audio and video social-media platform that can be used via an app. GreatApe is a Fediverse technology that supports federation via ActivityPub.
// @BasePath /
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
