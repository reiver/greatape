package routes

import (
	"app/activitypub"
	"config"
	. "contracts"
	"io/ioutil"
	"net/http"
	"net/url"
	"server/route"
	"strings"
)

var Follow = route.New(HttpGet, "/u/:name/follow", func(x IContext) error {
	username := x.Request().Params("name")
	follower := x.Request().Query("acct")

	parts := strings.Split(follower, "@")
	if len(parts) != 2 {
		return x.BadRequest("bad_request")
	}

	webfingerUrl := x.StringUtil().Format("https://%s/.well-known/webfinger?resource=acct:%s", parts[1], follower)
	resp, err := http.Get(webfingerUrl)
	if err != nil {
		x.InternalServerError(err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		x.InternalServerError(err.Error())
	}

	webfinger, err := activitypub.UnmarshalWebfinger(data)
	if err != nil {
		x.InternalServerError(err.Error())
	}

	template := ""
	for _, link := range webfinger.Links {
		if link.Rel == "http://ostatus.org/schema/1.0/subscribe" {
			template = *link.Template
			break
		}
	}

	if template == "" {
		x.BadRequest("There was an error looking up the remote account")
	}

	uri := url.QueryEscape(x.StringUtil().Format("https://%s/u/%s", config.DOMAIN, username))
	template = x.StringUtil().Replace(template, "{uri}", uri, 01)

	return x.Redirect(template)
})
