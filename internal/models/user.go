package models

type User struct {
	Id        uint64 `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Active    bool   `json:"active,omitempty"`
}


