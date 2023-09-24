package model

type DeliveredAddress struct {
	Country *string `json:"country"`
	Region  *string `json:"region"`
	Okrug   *string `json:"okrug"`
}
