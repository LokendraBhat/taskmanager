package models

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

type DashboardData struct {
	Tasks     []Task
	EditTask  *Task
	CSRFToken string
}

type User struct {
	ID           int
	Username     string
	PasswordHash string
	CreatedAt    string
}