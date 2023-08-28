package models

type UserSegment struct {
	Id           int     `json:"id,omitempty"`
	User         User    `json:"user,omitempty"`
	Segment      Segment `json:"segment,omitempty"`
	CreationTime int64   `json:"creation_time,omitempty"`
	DeletionTime int64   `json:"deletion_time,omitempty"`
	Duration     int64   `json:"duration,omitempty"`
	Active       bool    `json:"active,omitempty"`
}
