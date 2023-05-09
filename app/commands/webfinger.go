package commands

import (
	"strings"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func Webfinger(x IDispatcher, resource string) (IWebfingerResult, error) {
	username := strings.Replace(resource[5:], x.Format("@%s", x.FQDN()), "", -1)
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	subject := x.Format("acct:%s@%s", identity.Username(), x.FQDN())
	href := x.Format("%s/u/%s", x.PublicUrl(), identity.Username())
	template := x.Format("%s/authorize_interaction?uri={uri}", x.PublicUrl())

	aliases := []string{
		href,
	}

	self, _ := x.NewActivityPubLink()
	self.SetHref(href)
	self.SetRel("self")
	self.SetType("application/activity+json")

	ostatus, _ := x.NewActivityPubLink()
	ostatus.SetRel("http://ostatus.org/schema/1.0/subscribe")
	ostatus.SetTemplate(template)

	links := []IActivityPubLink{
		self, ostatus,
	}

	return x.NewWebfingerResult(aliases, links, subject), nil
}
