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

type RecoverReq struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `yaml:"new_password"`
}

type VerifyReq struct {
	Code  string `json:"code"`
	Email string `json:"email"`
	Typo  string `json:"typo"`
}
