package routes

import (
	"activitypub"
	"app/models/repos"
	"config"
	"encoding/hex"
	"fmt"
	"server/mime"
	"time"

	"github.com/mazen160/go-random"
)

func createApiKey() (string, error) {
	data, err := random.Bytes(16)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

func createActor(user *repos.User) *activitypub.Actor {
	id := fmt.Sprintf("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, user.Username)

	return &activitypub.Actor{
		Context: []interface{}{
			activitypub.ActivityStreams,
			"https://w3id.org/security/v1",
		},
		Followers:         id + "/followers",
		Following:         id + "/following",
		ID:                id,
		Inbox:             id + "/inbox",
		Name:              user.DisplayName,
		Outbox:            id + "/outbox",
		PreferredUsername: user.Username,
		Type:              "Person",
		Url:               id,
		Icon: activitypub.Icon{
			Height:    120,
			MediaType: "image/jpeg",
			Type:      "Image",
			URL:       user.Avatar,
			Width:     120,
		},
		Image: activitypub.Icon{
			Height:    317,
			MediaType: "image/jpeg",
			URL:       user.Banner,
			Width:     1920,
		},
		PublicKey: activitypub.PublicKey{
			ID:           id + "#main-key",
			Owner:        id,
			PublicKeyPem: user.PublicKey,
		},
		Summary:   user.Bio,
		Published: time.Now(),
	}
}

func createWebfinger(user *repos.User) *activitypub.Webfinger {
	subject := fmt.Sprintf("acct:%s@%s", user.Username, config.DOMAIN)
	href := fmt.Sprintf("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, user.Username)
	_type := mime.ActivityJson
	template := fmt.Sprintf("%s://%s/authorize_interaction?uri={uri}", config.PROTOCOL, config.DOMAIN)

	return &activitypub.Webfinger{
		Aliases: []string{
			href,
		},
		Links: []activitypub.Link{
			{
				Href: &href,
				Rel:  "self",
				Type: &_type,
			},
			{
				Rel:      OStatusSubscription,
				Template: &template,
			},
		},
		Subject: subject,
	}
}
