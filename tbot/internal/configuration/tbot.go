package configuration

// Group Telegram groups
type Group struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type Settings struct {
	Groups *[]Group `json:"groups"`
}
