package routes

import (
	"app/models/repos"
	"app/models/types"
	. "contracts"
	"encoding/json"
	"errors"
	"server/route"
	"utility/jwt"
	"utility/password"

	"gorm.io/gorm"
)

var Login = route.New(HttpPost, "/api/v1/login", func(x IContext) error {
	body := new(types.LoginDTO)

	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	u := &repos.User{}

	err := repos.FindUserByEmail(u, body.Email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.Unauthorized("Invalid email or password")
	}

	if err := password.Verify(u.Password, body.Password); err != nil {
		return x.Unauthorized("Invalid email or password")
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	})

	actor, _ := json.MarshalIndent(createActor(u), "", "  ")
	webfinger, _ := json.MarshalIndent(createWebfinger(u), "", "  ")
	return x.Json(&types.AuthResponse{
		User: &types.UserResponse{
			ID:          u.ID,
			Username:    u.Username,
			DisplayName: u.DisplayName,
			Email:       u.Email,
			Bio:         u.Bio,
			Github:      u.Github,
			Avatar:      u.Avatar,
			Banner:      u.Banner,
			ApiKey:      u.ApiKey,
			PublicKey:   u.PublicKey,
			Actor:       string(actor),
			Webfinger:   string(webfinger),
		},
		Auth: &types.AccessResponse{
			Token: token,
		},
	})
})
