package types

import "time"

type TimeEntry struct {
	ID          int64     `json:"id"`
	WorkspaceID int       `json:"workspace_id"`
	ProjectID   int       `json:"project_id"`
	Billable    bool      `json:"billable"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
	Duration    int       `json:"duration"`
	Description string    `json:"description"`
	Duronly     bool      `json:"duronly"`
	At          time.Time `json:"at"`
	UserID      int       `json:"user_id"`
}

type Project struct {
	ID            int       `json:"id"`
	WorkspaceID   int       `json:"workspace_id"`
	Name          string    `json:"name"`
	IsPrivate     bool      `json:"is_private"`
	Active        bool      `json:"active"`
	At            time.Time `json:"at"`
	CreatedAt     time.Time `json:"created_at"`
	Color         string    `json:"color"`
	Billable      bool      `json:"billable"`
	Recurring     bool      `json:"recurring"`
	ActualHours   int       `json:"actual_hours"`
	ActualSeconds int       `json:"actual_seconds"`
	TotalCount    int       `json:"total_count"`
	CanTrackTime  bool      `json:"can_track_time"`
	StartDate     string    `json:"start_date"`
	Status        string    `json:"status"`
	Pinned        bool      `json:"pinned"`
}
