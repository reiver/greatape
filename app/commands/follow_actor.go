package commands

import (
	"io"
	"net/http"
	"strings"

	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func FollowActor(x IDispatcher, username string, acct string) (IFollowActorResult, error) {
	parts := strings.Split(acct, "@")
	x.Assert(len(parts) == 2).Or(ERROR_INVALID_PARAMETERS)

	webfingerUrl := x.Format("https://%s/.well-known/webfinger?resource=acct:%s", parts[1], acct)
	resp, err := http.Get(webfingerUrl)
	x.AssertNoError(err)

	data, err := io.ReadAll(resp.Body)
	x.AssertNoError(err)

	webfinger, err := activitypub.UnmarshalWebfinger(data)
	x.AssertNoError(err)

	subject := ""
	for _, link := range webfinger.Links {
		if link.Rel == "self" {
			subject = *link.Href
			break
		}
	}

	if x.IsEmpty(subject) {
		return nil, ERROR_INVALID_PARAMETERS
	}

	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	follower := x.GetActorId()

	followee := &activitypub.Actor{}
	if err := x.GetActivityStreamSigned(subject, nil, followee); err != nil {
		return nil, err
	}

	uniqueIdentifier := x.GenerateUUID()
	follow := activitypub.NewFollow(follower, followee.ID, uniqueIdentifier)

	x.Atomic(func() error {
		activity := x.MarshalJson(follow)

		x.AddActivityPubOutgoingActivity(
			identity.Id(),
			uniqueIdentifier,
			x.UnixNano(),
			follower,
			followee.ID,
			activitypub.TypeFollow,
			activity,
		)

		x.AddActivityPubFollower(
			follower,
			followee.Inbox,
			followee.ID,
			activity,
			false,
		)

		if err := x.PostActivityStreamSigned(followee.Inbox, follow, nil); err != nil {
			return err
		}

		return nil
	})

	return x.NewFollowActorResult(follow.Id), nil
}
