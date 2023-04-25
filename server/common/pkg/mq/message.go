package mq

type Message struct {
	ID int64 `json:"id"` // ID of the message
}

const (
	HEADER_ETL_INFO = "ETL-INFO"
)

type MessageETLInfoRequest struct {
	ID int64 `json:"id"` // ID of the message
}

type MessageETLInfoResponse struct {
	ID    int64    `json:"id"`    // ID of the message
	Owner string   `json:"owner"` // Owner of the message
	Data  []string `json:"data"`  // Data of the message
}
