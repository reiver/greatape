package routes

import (
	"app/models/repos"
	"app/models/types"
	. "contracts"
	"errors"
	"server/route"
	"utility"
	"utility/jwt"
	"utility/password"

	"gorm.io/gorm"
)

var Signup = route.New(HttpPost, "/api/v1/signup", func(x IContext) error {
	body := new(types.SignupDTO)
	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	err := repos.FindUserByEmail(&struct{ ID string }{}, body.Email).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return x.Conflict("Email already exists")
	}

	privateKey, publicKey, err := utility.GenerateRSAKeyPair()
	if err != nil {
		return err
	}

	apiKey, err := createApiKey()
	if err != nil {
		return err
	}

	user := &repos.User{
		Username:   body.Username,
		Password:   password.Generate(body.Password),
		Email:      body.Email,
		ApiKey:     apiKey,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	if err := repos.CreateUser(user); err.Error != nil {
		return x.Conflict(err.Error.Error())
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return x.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:          user.ID,
			DisplayName: user.Username,
			Email:       user.Email,
		},
		Auth: &types.AccessResponse{
			Token: token,
		},
	})
})
