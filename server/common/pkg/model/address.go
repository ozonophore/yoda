package model

import "time"

type DeliveredAddress struct {
	Country *string   `json:"country"`
	Region  *string   `json:"region"`
	Okrug   *string   `json:"okrug"`
	Date    time.Time `json:"date"`
}
