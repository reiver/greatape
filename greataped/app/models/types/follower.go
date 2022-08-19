package types

type FollowerDTO struct {
	Target   string `json:"target" validate:"required"`
	Handle   string `json:"handle" validate:"required"`
	Accepted bool   `json:"accepted"`
}

type FollowerResponse struct {
	Target   string `json:"target"`
	Handle   string `json:"handle"`
	Accepted bool   `json:"accepted"`
}
