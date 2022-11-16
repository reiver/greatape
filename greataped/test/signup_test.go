package test

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

func signup(api IApi) error {

	rand.Seed(time.Now().UnixNano())

	var (
		token, code string
		id          = 100000 + rand.Intn(899999)
		username    = fmt.Sprintf("u%d", id)
		email       = fmt.Sprintf("%s@domain.com", username)
		password    = "AaBbCc1$"
		displayName = fmt.Sprintf("n%d", id)
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

	// GetProfileByUser
	{
		input := &GetProfileByUserRequest{}

		output, err := api.GetProfileByUser(input)
		if err != nil {
			return err
		}

		if output.Username != username {
			return errors.New("get_profile_by_user_failed")
		}
	}

	// UpdateProfileByUser
	{
		input := &UpdateProfileByUserRequest{
			DisplayName: displayName,
			Avatar:      "Avatar",
			Banner:      "Banner",
			Summary:     "Summary",
			Github:      "Github",
		}

		output, err := api.UpdateProfileByUser(input)
		if err != nil {
			return err
		}

		if output.DisplayName != displayName {
			return errors.New("update_profile_by_user_failed")
		}
	}

	// GetProfileByUser
	{
		input := &GetProfileByUserRequest{}

		output, err := api.GetProfileByUser(input)
		if err != nil {
			return err
		}

		if output.DisplayName != displayName {
			return errors.New("get_profile_by_user_failed")
		}
	}

	return nil
}
