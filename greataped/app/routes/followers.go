package routes

import (
	"activitypub"
	"app/models/domain"
	"app/models/dto"
	"app/models/repos"
	"config"
	. "contracts"
	"encoding/json"
	"server/route"
	"strconv"
)

var Followers = route.New(HttpGet, "/u/:username/followers", func(x IContext) error {
	username := domain.Username(x.Request().Params("username"))
	if username.IsEmpty() {
		return x.BadRequest("username required.")
	}

	if username.IsFederated() {
		webfinger, err := x.GetWebFinger(username)
		if err != nil {
			return err
		}

		actor, err := x.GetActor(webfinger)
		if err != nil {
			return err
		}

		followers, err := x.GetOrderedCollection(actor.Followers)
		if err != nil {
			return err
		}

		return x.Activity(followers)
	} else {
		actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username)
		id := x.StringUtil().Format("%s://%s/u/%s/followers", config.PROTOCOL, config.DOMAIN, username)

		followers := &[]dto.FollowerResponse{}
		err := repos.FindFollowers(followers, actor).Error
		if err != nil {
			return err
		}

		items := []string{}
		for _, follower := range *followers {
			items = append(items, follower.Handle)
		}

		result := &activitypub.Followers{
			Context:      activitypub.ActivityStreams,
			ID:           id,
			Type:         activitypub.TypeOrderedCollection,
			TotalItems:   len(items),
			OrderedItems: items,
		}

		return x.Activity(result)
	}
})

var AcceptFollowRequest = route.New(HttpPut, "/u/:username/followers/:id/accept", func(x IContext) error {
	username := x.Request().Params("username")
	id := x.Request().Params("id")
	followerId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return x.BadRequest("invalid_id")
	}

	follower := &repos.Follower{}
	if err := repos.FindFollowerById(follower, followerId).Error; err != nil {
		return err
	}

	data, _ := json.Marshal(&activitypub.Activity{
		Context: activitypub.ActivityStreams,
		ID:      x.StringUtil().Format("%s://%s/%s", config.PROTOCOL, config.DOMAIN, x.GUID()),
		Type:    activitypub.TypeAccept,
		Actor:   x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
		Object:  follower.Activity,
	})

	user, err := repos.FindUserByUsername(username)
	if err != nil {
		return err
	}

	keyId := x.StringUtil().Format("%s://%s/u/%s#main-key", config.PROTOCOL, config.DOMAIN, username)

	if err := x.PostActivityStreamSigned(follower.HandleInbox, keyId, user.PrivateKey, data, nil); err != nil {
		return err
	}

	if err := repos.AcceptFollower(follower.ID).Error; err != nil {
		return err
	}

	return x.Nothing()
})
