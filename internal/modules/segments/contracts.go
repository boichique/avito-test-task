package segments

type CreateSegmentRequest struct {
	Slug string `json:"slug"`
}

type DeleteSegmentRequest struct {
	Slug string `json:"slug"`
}

type GetUserSegmentsRequest struct {
	UserID int `param:"userID"`
}

type GetUserSegmentsResponse struct {
	Segments []Segment `json:"segments"`
}

type UpdateUserSegmentsRequest struct {
	UserID         int      `param:"userID"`
	AddSegments    []string `json:"add_segments,omitempty"`
	DeleteSegments []string `json:"delete_segments,omitempty"`
}
