package routes

import (
	"app/models/dto"
	"config"
	. "contracts"
	"db/repos"
	"encoding/json"
	"server/route"
)

var Profile = route.New(HttpGet, "/profile", func(x IContext) error {
	return x.Render("profile", ViewData{
		"Title":    "Profile",
		"Protocol": config.PROTOCOL,
		"Domain":   config.DOMAIN,
	})
})

// GetProfile	godoc
// @Tags     User
// @Accept   json
// @Produce  json
// @Security JWT
// @Success  200 {object} dto.User
// @Router   /api/v1/profile [get]
func _() {}

var GetProfile = route.New(HttpGet, "/api/v1/profile", func(x IContext) error {
	user, err := repos.Default.FindUserById(x.GetUser())
	if err != nil {
		return x.Unauthorized(err)
	}

	actor, _ := json.MarshalIndent(createActor(user), "", "  ")
	webfinger, _ := json.MarshalIndent(createWebfinger(user), "", "  ")
	return x.Json(dto.User{
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
	body := new(dto.ProfileRequest)

	if err := x.ParseBodyAndValidate(body); err != nil {
		return err
	}

	user, err := repos.Default.FindUserById(x.GetUser())
	if err != nil {
		return x.Unauthorized(err)
	}

	access := repos.ACCESS_PUBLIC
	if body.PrivateProfile {
		access = repos.ACCESS_PRIVATE
	}

	if err := repos.Default.UpdateProfile(user.ID,
		map[string]interface{}{
			"display_name": body.DisplayName,
			"bio":          body.Bio,
			"github":       body.Github,
			"avatar":       body.Avatar,
			"banner":       body.Banner,
			"access":       access,
		}); err != nil {
		return err
	}

	return x.Nothing()
})
