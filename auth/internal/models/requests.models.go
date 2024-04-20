package models

type RegReq struct {
	UUID     string `json:"uuid"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Token       string `json:"token"`
}
