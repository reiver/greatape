package test

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
)

func logout(api IApi) error {

	rand.Seed(time.Now().UnixNano())

	var (
		token, code string
		id          = 100000 + rand.Intn(899999)
		username    = fmt.Sprintf("u%d", id)
		email       = fmt.Sprintf("%s@domain.com", username)
		password    = "AaBbCc1$"
	)

	// Signup
	{
		input := &SignupRequest{
			Username: username,
			Email:    email,
			Password: password,
		}

		output, err := api.Signup(input)
		if err != nil {
			return err
		}

		token = output.Token
		code = output.Code
	}

	// Verify
	{
		input := &VerifyRequest{
			Email: email,
			Token: token,
			Code:  code,
		}

		output, err := api.Verify(input)
		if err != nil {
			return err
		}

		_ = output

		// api.SetToken(output.Token + "asdf")
	}

	// Logout
	{
		input := &LogoutRequest{}

		output, err := api.Logout(input)
		if err != nil {
			return err
		}

		_ = output
	}

	// Login
	{
		input := &LoginRequest{
			Email:    email,
			Password: password,
		}

		output, err := api.Login(input)
		if err != nil {
			return err
		}

		api.SetToken(output.Token)
	}

	// Logout
	{
		input := &LogoutRequest{}

		output, err := api.Logout(input)
		if err != nil {
			return err
		}

		_ = output
	}

	return nil
}
