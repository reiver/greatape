package dto

type MessageRequest struct {
	Guid    string `json:"guid" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type MessageResponse struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}
