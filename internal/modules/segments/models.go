package segments

import "time"

type Segment struct {
	Slug      string     `json:"slug"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
