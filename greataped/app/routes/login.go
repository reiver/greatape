package routes

import (
	"app/models/repos"
	"app/models/types"
	. "contracts"
	"encoding/json"
	"server/route"
	"utility/jwt"
	"utility/password"
)

var Login = route.New(HttpPost, "/api/v1/login", func(x IContext) error {
	body := new(types.LoginDTO)

	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	user, err := repos.FindUserByEmail(body.Email)
	if err != nil {
		return x.Unauthorized("invalid email or password")
	}

	if err := password.Verify(user.Password, body.Password); err != nil {
		return x.Unauthorized("invalid email or password")
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	actor, _ := json.MarshalIndent(createActor(user), "", "  ")
	webfinger, _ := json.MarshalIndent(createWebfinger(user), "", "  ")
	return x.Json(&types.AuthResponse{
		User: &types.UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			Bio:         user.Bio,
			Github:      user.Github,
			Avatar:      user.Avatar,
			Banner:      user.Banner,
			ApiKey:      user.ApiKey,
			PublicKey:   user.PublicKey,
			Actor:       string(actor),
			Webfinger:   string(webfinger),
		},
		Auth: &types.AccessResponse{
			Token: token,
		},
	})
})
