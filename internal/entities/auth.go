package entities

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserToken struct {
	Token *string `json:"token,omitempty"`
}
