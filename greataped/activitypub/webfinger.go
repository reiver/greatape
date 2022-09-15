package activitypub

import "encoding/json"

func UnmarshalWebfinger(data []byte) (Webfinger, error) {
	var r Webfinger
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Webfinger) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

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
