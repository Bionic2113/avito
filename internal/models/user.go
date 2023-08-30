package models

type User struct {
	Id        uint64       `json:"id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Active    bool      
}


