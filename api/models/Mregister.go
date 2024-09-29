package models

type Usuario struct {
	UID          string `json:"uid"`
	DisplayName  string `json:"displayName"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	Email        string `json:"email"`
	Verified     bool   `json:"verified"`
}
