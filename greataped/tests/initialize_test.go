package tests

import (
	"app/models/repos"
	"app/routes"
	"bytes"
	"caching"
	"config"
	"db"
	"encoding/json"
	"fmt"
	"logging"
	"net/http"
	"net/url"
	"os"
	"server"
	"server/mime"
	"testing"
)

type Payload map[string]interface{}

const (
	DOMAIN = "domain.social"
	ROOT   = "http://localhost"
)

func TestMain(m *testing.M) {
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

	go func() {
		app.Listen(fmt.Sprintf(":%s", config.PORT))
	}()

	os.Exit(m.Run())
}

func Get(path string) (*http.Response, error) {
	var err error
	path, err = url.JoinPath(ROOT, path)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Get(path)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func Post(path string, payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	path, err = url.JoinPath(ROOT, path)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Post(path, mime.ActivityJsonUtf8, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return resp, err
}
