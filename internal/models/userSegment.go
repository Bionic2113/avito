package models

type UserSegment struct {
	Id           uint64 `json:"id,omitempty"`
	User_id      uint64 `json:"user_id,omitempty"`
	Segment_id   uint64 `json:"segment_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	DeletionTime int64  `json:"deletion_time,omitempty"`
	Duration     int64  `json:"duration,omitempty"`
	Active       bool   `json:"active,omitempty"`
}
