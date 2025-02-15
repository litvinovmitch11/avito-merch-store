package entities

type UserAuth struct {
	ID       string
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPersonalData struct {
	ID             string
	UserID         string
	HashedPassword string
}

type UserToken struct {
	Token *string `json:"token,omitempty"`
}
