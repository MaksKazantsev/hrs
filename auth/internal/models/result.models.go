package models

type RegRes struct {
	UUID  string `json:"uuid"`
	Token string `json:"-"`
}

type LoginRes struct {
	Token string `json:"-"`
}
