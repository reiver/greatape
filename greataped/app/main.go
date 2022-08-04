package main

import (
	"app/models/repos"
	"app/routes"
	"config"
	"db"
	"fmt"
	"server"
)

func main() {
	storage := db.NewSqliteStorage()
	storage.Connect(config.SQLITE_DB)
	storage.Migrate(
		&repos.User{},
		&repos.IncomingActivity{},
		&repos.OutgoingActivity{},
		&repos.Follower{},
		&repos.Following{},
	)

	app := server.New()
	app.SetStorageProvider(storage)

	app.Bind(
		routes.Root,
		routes.Profile,
		routes.Signup,
		routes.Login,
		routes.GetProfile,
		routes.UpdateProfile,
		routes.WebFinger,
		routes.User,
		routes.Message,
		routes.InboxPost,
		routes.InboxGet,
		routes.OutboxPost,
		routes.OutboxGet,
		routes.Followers,
		routes.Follow,
		routes.Authorize,
	)

	app.Listen(fmt.Sprintf(":%s", config.PORT))
}
