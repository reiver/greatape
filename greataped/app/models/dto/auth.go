package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginResponse struct {
	User User `json:"user"`
	Auth Auth `json:"auth"`
}

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
}

type SignupResponse struct {
	Code string `json:"code"`
}

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type VerifyResponse struct {
	User User `json:"user"`
	Auth Auth `json:"auth"`
}

type User struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	DisplayName    string `json:"display_name,omitempty"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	Bio            string `json:"bio,omitempty"`
	Github         string `json:"github,omitempty"`
	Avatar         string `json:"avatar,omitempty"`
	Banner         string `json:"banner,omitempty"`
	ApiKey         string `json:"api_key,omitempty"`
	PublicKey      string `json:"public_key,omitempty"`
	Actor          string `json:"actor,omitempty"`
	Webfinger      string `json:"webfinger,omitempty"`
	PrivateProfile bool   `json:"private_profile"`
}

type Auth struct {
	Token string `json:"token"`
}
