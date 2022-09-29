package routes

import (
	"app/models/repos"
	"app/models/types"
	. "contracts"
	"server/route"
	"utility"
	"utility/jwt"
	"utility/password"
)

var Signup = route.New(HttpPost, "/api/v1/signup", func(x IContext) error {
	body := new(types.SignupDTO)
	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	if _, err := repos.FindUserByEmail(body.Email); err == nil {
		return x.Conflict("email already exists")
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

	code := utility.GenerateConfirmationCode()
	x.Cache().Put(user.Email, &struct {
		user *repos.User
		code string
	}{
		user: user,
		code: code,
	})

	return x.Json(struct{ Code string }{
		Code: code, // TODO: Remove and send with email
	})
})

var Verify = route.New(HttpPost, "/api/v1/verify", func(x IContext) error {
	body := new(types.VerificationDTO)
	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	item := x.Cache().Get(body.Email)
	if item == nil {
		return x.BadRequest("not found")
	}

	registration := item.(*struct {
		user *repos.User
		code string
	})

	if registration.code != body.Code {
		return x.Unauthorized("invalid code")
	}

	user := registration.user
	if err := repos.CreateUser(user); err != nil {
		return x.Conflict(err)
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return x.Json(&types.AuthResponse{
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
