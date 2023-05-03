package activitypub

import "encoding/json"

type Webfinger struct {
	Aliases []string `json:"aliases"`
	Links   []Link   `json:"links"`
	Subject string   `json:"subject"`
}

type Link struct {
	Href     *string `json:"href,omitempty"`
	Rel      string  `json:"rel"`
	Type     *string `json:"type,omitempty"`
	Template *string `json:"template,omitempty"`
}

func (webfinger *Webfinger) Self() string {
	self := ""
	for _, link := range webfinger.Links {
		if link.Rel == "self" && link.Type != nil && *link.Type == "application/activity+json" {
			self = *link.Href
			break
		}
	}

	return self
}

func UnmarshalWebfinger(data []byte) (Webfinger, error) {
	var webfinger Webfinger
	err := json.Unmarshal(data, &webfinger)
	return webfinger, err
}

func (webfinger *Webfinger) Marshal() ([]byte, error) {
	return json.Marshal(webfinger)
}
