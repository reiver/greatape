package activitypub

import "encoding/json"

type Followers struct {
	Context      string      `json:"@context"`
	ID           string      `json:"id,omitempty"`
	Type         string      `json:"type,omitempty"`
	TotalItems   int         `json:"totalItems"`
	OrderedItems interface{} `json:"orderedItems,omitempty"`
}

func UnmarshalFollowers(data []byte) (Followers, error) {
	var o Followers
	err := json.Unmarshal(data, &o)
	return o, err
}

func (o *Followers) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
