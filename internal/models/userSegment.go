package models

type UserSegment struct {
	Id           uint64 `json:"id,omitempty"`
	User_id      uint64 `json:"user_id,omitempty"`
	Segment_name string `json:"segment_id,omitempty"`
	CreationTime uint64 `json:"creation_time,omitempty"`
	DeletionTime uint64 `json:"deletion_time,omitempty"`
	Duration     uint64 `json:"duration,omitempty"`
	Active       bool   `json:"active,omitempty"`
}
