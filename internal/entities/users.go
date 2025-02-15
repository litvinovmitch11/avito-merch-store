package entities

type User struct {
	ID       string
	Username string `json:"username"`
}

type UserAuth struct {
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
