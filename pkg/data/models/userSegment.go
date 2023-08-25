package models

type UserSegment struct {
	Id      int     `json:"id,omitempty"`
	User    User    `json:"user,omitempty"`
	Segment Segment `json:"segment,omitempty"`
}
