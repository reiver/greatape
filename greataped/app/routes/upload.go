package routes

import (
	"config"
	"contracts"
	"path"
	"server/route"

	"github.com/google/uuid"
)

var Upload = route.New(contracts.HttpPost, "/upload", func(x contracts.IContext) error {
	file, err := x.Request().FormFile("file")
	if err != nil {
		return err
	}

	uuid := uuid.New().String()
	extension := path.Ext(file.Filename)
	fileName := x.StringUtil().Format("%s%s", uuid, extension)

	filePath := path.Join(config.UPLOAD_PATH, fileName)
	if err = x.SaveFile(file, filePath); err != nil {
		return err
	}

	return x.Json(struct {
		Url string `json:"url"`
	}{
		Url: x.StringUtil().Format("%s://%s/media/%s", config.PROTOCOL, config.DOMAIN, fileName),
	})
})
