package models

type Segment struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Time int64  `json:"time,omitempty"`
}

const (
	AVITO_VOICE_MESSAGES  string = "AVITO_VOICE_MESSAGES"
	AVITO_PERFORMANCE_VAS string = "AVITO_PERFORMANCE_VAS"
	AVITO_DISCOUNT_30     string = "AVITO_DISCOUNT_30"
	AVITO_DISCOUNT_50     string = "AVITO_DISCOUNT_50"
)
