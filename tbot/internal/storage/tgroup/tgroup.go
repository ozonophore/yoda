package tgroup

func (*Group) TableName() string {
	return `ml."tg_group""`
}

type Group struct {
	UserName  string `json:"user_name"`
	GroupName string `json:"group_name"`
	ChatID    int64  `json:"chat_id"`
}
