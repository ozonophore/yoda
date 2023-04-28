package dto

import "time"

type RefreshAction struct {
	Type    string    `json:"type"`
	NextRun time.Time `json:"nextRun"`
}
