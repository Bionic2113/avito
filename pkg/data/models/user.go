package models

type User struct {
	Id        int       `json:"id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Segments  []Segment `json:"segments,omitempty"`
	Active    bool      `json:"active,omitempty"`
}


