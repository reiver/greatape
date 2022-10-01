package routes

import (
	"app/models/dto"
	. "contracts"
	"db/repos"
	"encoding/json"
	"server/route"
	"utility"
	"utility/jwt"
	"utility/password"
)

// Signup	godoc
// @Tags	Authentication
// @Accept	json
// @Produce	json
// @Param	payload body dto.SignupRequest true "Payload"
// @Success	200 {object} dto.SignupResponse
// @Router	/api/v1/signup [post]
func _() {}

var Signup = route.New(HttpPost, "/api/v1/signup", func(x IContext) error {
	body := new(dto.SignupRequest)
	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	if _, err := repos.Default.FindUserByEmail(body.Email); err == nil {
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

	return x.Json(dto.SignupResponse{
		Code: code, // TODO: Remove and send with email
	})
})

// Verify	godoc
// @Tags	Authentication
// @Accept	json
// @Produce	json
// @Param	payload body dto.VerifyRequest true "Payload"
// @Success	200 {object} dto.VerifyResponse
// @Router	/api/v1/verify [post]
func _() {}

var Verify = route.New(HttpPost, "/api/v1/verify", func(x IContext) error {
	body := new(dto.VerifyRequest)
	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	item := x.Cache().Get(body.Email)
	if item == nil {
		return x.BadRequest("not found")
	}

	x.Cache().Remove(body.Email)

	registration := item.(*struct {
		user *repos.User
		code string
	})

	if registration.code != body.Code {
		return x.Unauthorized("invalid code")
	}

	user := registration.user
	if err := repos.Default.CreateUser(user); err != nil {
		return x.Conflict(err)
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	actor, _ := json.MarshalIndent(createActor(user), "", "  ")
	webfinger, _ := json.MarshalIndent(createWebfinger(user), "", "  ")
	return x.Json(dto.VerifyResponse{
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
