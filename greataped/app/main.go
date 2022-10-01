package main

import (
	"app/docs"
	"app/routes"
	"caching"
	"config"
	"db"
	"db/repos"
	"fmt"
	"logging"
	"server"
)

// @title GreatApe API
// @version 1.0
// @description GreatApe is a free audio and video social-media platform that can be used via an app.
// @description It is a Fediverse technology that supports federation via ActivityPub.
// @BasePath /
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @description Example: Bearer {Your JWT Token}
func main() {
	if config.IsProduction() {
		docs.SwaggerInfo.Host = config.DOMAIN
	}

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
