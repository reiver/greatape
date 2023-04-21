package spi

import (
	"time"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func GetActor(x IDispatcher, username string) (IGetActorResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	id := x.Format("%s/u/%s", x.PublicUrl(), identity.Username())

	context := []string{
		ACTIVITY_STREAMS,
		W3ID_SECURITY_V1,
	}

	var icon IActivityPubMedia
	if x.IsNotEmpty(identity.Avatar()) {
		icon, _ = x.NewActivityPubMedia()
		icon.SetType(ACTIVITY_PUB_IMAGE)
		icon.SetUrl(identity.Avatar())
	}

	var image IActivityPubMedia
	if x.IsNotEmpty(identity.Banner()) {
		image, _ = x.NewActivityPubMedia()
		image.SetType(ACTIVITY_PUB_IMAGE)
		image.SetUrl(identity.Banner())
	}

	publicKey, _ := x.NewActivityPubPublicKey()
	publicKey.SetId(x.Format("%s#main-key", id))
	publicKey.SetOwner(id)
	publicKey.SetPublicKeyPem(identity.PublicKey())

	published := time.Now().Format("2006-01-02T15:04:05Z")

	return x.NewGetActorResult(
		context,                      // context
		id,                           // id
		x.Format("%s/followers", id), // followers
		x.Format("%s/following", id), // following
		x.Format("%s/inbox", id),     // inbox
		x.Format("%s/outbox", id),    // outbox
		identity.DisplayName(),       // name
		identity.Username(),          // preferredUsername
		ACTIVITY_PUB_PERSON,          // type
		id,                           // url
		icon,                         // icon
		image,                        // image
		publicKey,                    // publicKey
		identity.Summary(),           // summary
		published,                    // published
	), nil
}
