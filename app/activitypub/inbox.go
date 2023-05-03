package activitypub

import "encoding/json"

type Inbox struct {
	Context      string      `json:"@context"`
	ID           string      `json:"id,omitempty"`
	Type         string      `json:"type,omitempty"`
	TotalItems   int         `json:"totalItems"`
	OrderedItems interface{} `json:"orderedItems,omitempty"`
}

func UnmarshalInbox(data []byte) (Inbox, error) {
	var inbox Inbox
	err := json.Unmarshal(data, &inbox)
	return inbox, err
}

func (inbox *Inbox) Marshal() ([]byte, error) {
	return json.Marshal(inbox)
}
