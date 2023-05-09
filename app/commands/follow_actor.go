package commands

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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

	template := ""
	for _, link := range webfinger.Links {
		if link.Rel == OSTATUS_SUBSCRIPTION {
			template = *link.Template
			break
		}
	}

	if template == "" {
		return nil, fmt.Errorf("remote_account_lookup_failed")
	}

	uri := url.QueryEscape(x.Format("%s/u/%s", x.PublicUrl(), username))
	template = strings.Replace(template, "{uri}", uri, -1)

	return x.NewFollowActorResult(template), nil
}
