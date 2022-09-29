package routes

import (
	"app/models/dto"
	"app/models/repos"
	. "contracts"
	"encoding/json"
	"server/route"
	"utility/jwt"
	"utility/password"
)

// Login godoc
// @Tags	Authentication
// @Accept	json
// @Produce	json
// @Param	payload	body	dto.LoginRequest	true	"Payload"
// @Success	200	{object}	dto.LoginResponse
// @Router	/api/v1/login	[post]
func _() {}

var Login = route.New(HttpPost, "/api/v1/login", func(x IContext) error {
	body := new(dto.LoginRequest)

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
	return x.Json(dto.LoginResponse{
		User: dto.User{
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
		Auth: dto.Auth{
			Token: token,
		},
	})
})
