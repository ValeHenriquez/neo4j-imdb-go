package models

type Actor struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character"`
}
