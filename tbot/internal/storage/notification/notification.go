package notification

func (*Notification) TableName() string {
	return `ml."notification"`
}

type Notification struct {
	ID      int64   `json:"id"`
	Type    string  `json:"type"`
	Message *string `json:"message"`
	IsSent  bool    `json:"is_sent"`
	Sender  string  `json:"sender"`
}
