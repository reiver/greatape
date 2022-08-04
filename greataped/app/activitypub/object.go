package activitypub

import "encoding/json"

const ActivityStreams = "https://www.w3.org/ns/activitystreams"

type Object struct {
	Context string `json:"@context" validate:"activitystream"`
	Type    string `json:"type"`
}

func UnmarshalObject(data []byte) (Object, error) {
	var object Object
	err := json.Unmarshal(data, &object)
	return object, err
}

func (object *Object) Marshal() ([]byte, error) {
	return json.Marshal(object)
}
