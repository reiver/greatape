package dto

type FollowingRequest struct {
	Target string `json:"target" validate:"required"`
	Handle string `json:"handle" validate:"required"`
}

type FollowingResponse struct {
	Target string `json:"target"`
	Handle string `json:"handle"`
}
