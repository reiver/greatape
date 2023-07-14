package activitypub

import (
	"encoding/json"
	"fmt"
)

type Follow struct {
	Context string `json:"@context" validate:"activitystream"`
	Id      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

func NewFollow(follower, followee, uuid string) *Follow {
	return &Follow{
		Context: ActivityStreams,
		Id:      fmt.Sprintf("%s#follow/%s", follower, uuid),
		Type:    TypeFollow,
		Actor:   follower,
		Object:  followee,
	}
}

func UnmarshalFollow(data []byte) (Follow, error) {
	var follow Follow
	err := json.Unmarshal(data, &follow)
	return follow, err
}

func (follow *Follow) Marshal() ([]byte, error) {
	return json.Marshal(follow)
}
