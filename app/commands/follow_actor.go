package commands

import (
	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/contracts"
)

func FollowActor(x IDispatcher, username string, account string) (IFollowActorResult, error) {
	webfinger, err := x.ResolveWebfinger(account)
	x.AssertNoError(err)

	identity := x.GetIdentityByUsername(username)
	follower := x.GetActorId(identity)
	followee := &activitypub.Actor{}

	if err := x.GetSignedActivityStream(webfinger.Self(), followee, identity); err != nil {
		return nil, err
	}

	followers := x.FilterActivityPubFollowers(func(follow IActivityPubFollower) bool {
		return follow.Handle() == follower && follow.Subject() == followee.Id
	})

	if followers.HasAtLeastOneItem() && followers.First().Accepted() {
		return x.NewFollowActorResult(), nil
	}

	follow := activitypub.NewFollow(follower, followee.Id)

	x.Atomic(func() error {
		if followers.IsEmpty() {
			activity := x.MarshalJson(follow)

			x.AddActivityPubOutgoingActivity(
				identity.Id(),
				follow.UniqueIdentifier,
				x.UnixNano(),
				follower,
				followee.Id,
				activitypub.TypeFollow,
				activity,
			)

			x.AddActivityPubFollower(
				follower,
				followee.Inbox,
				followee.Id,
				activity,
				false,
			)
		}

		if err := x.PostSignedActivityStream(followee.Inbox, follow, identity); err != nil {
			return err
		}

		return nil
	})

	return x.NewFollowActorResult(), nil
}
