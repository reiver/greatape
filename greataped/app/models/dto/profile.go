package dto

type ProfileRequest struct {
	DisplayName    string `json:"display_name"`
	Bio            string `json:"bio"`
	Github         string `json:"github"`
	Avatar         string `json:"avatar"`
	Banner         string `json:"banner"`
	PrivateProfile bool   `json:"private_profile"`
}
