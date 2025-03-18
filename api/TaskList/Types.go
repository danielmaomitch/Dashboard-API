package TaskList

type Task struct {
	UserID    string `json:"UserID"`
	SK        string `json:"SK"`
	Body      string `json:"Body"`
	CreatedAt string `json:"CreatedAt"`
}
