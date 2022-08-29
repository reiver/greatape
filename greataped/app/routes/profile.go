package routes

import (
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"encoding/json"
	"errors"
	"server/route"

	"gorm.io/gorm"
)

var Profile = route.New(HttpGet, "/profile", func(x IContext) error {
	return x.Render("profile", ViewData{
		"Title":    "Profile",
		"Protocol": config.PROTOCOL,
		"Domain":   config.DOMAIN,
	})
})

var GetProfile = route.New(HttpGet, "/api/v1/profile", func(x IContext) error {
	userId := x.GetUser()
	user := &repos.User{}

	if err := repos.FindUserById(user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return x.Unauthorized("not_found")
		} else {
			return x.InternalServerError(err.Error())
		}
	}

	actor, _ := json.MarshalIndent(createActor(user), "", "  ")
	webfinger, _ := json.MarshalIndent(createWebfinger(user), "", "  ")
	return x.Json(&types.UserResponse{
		Username:       user.Username,
		DisplayName:    user.DisplayName,
		Bio:            user.Bio,
		Github:         user.Github,
		Avatar:         user.Avatar,
		Banner:         user.Banner,
		ApiKey:         user.ApiKey,
		Actor:          string(actor),
		Webfinger:      string(webfinger),
		PrivateProfile: user.Access == repos.ACCESS_PRIVATE,
	})
})

var UpdateProfile = route.New(HttpPost, "/api/v1/profile", func(x IContext) error {
	body := new(types.ProfileDTO)

	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	userId := x.GetUser()
	user := &repos.User{}

	err := repos.FindUserById(user, userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.Unauthorized("not_found")
	}

	access := repos.ACCESS_PUBLIC
	if body.PrivateProfile {
		access = repos.ACCESS_PRIVATE
	}

	if err := repos.UpdateProfile(user.ID, map[string]interface{}{
		"display_name": body.DisplayName,
		"bio":          body.Bio,
		"github":       body.Github,
		"avatar":       body.Avatar,
		"banner":       body.Banner,
		"access":       access,
	}).Error; err != nil {
		return x.InternalServerError("update_failed")
	}

	return x.Nothing()
})
