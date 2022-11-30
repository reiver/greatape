package activitypub

type Create struct {
	Context string      `json:"@context"`
	Type    string      `json:"type"`
	Id      string      `json:"id"`
	To      []string    `json:"to"`
	Actor   string      `json:"actor"`
	Object  interface{} `json:"object"`
}
