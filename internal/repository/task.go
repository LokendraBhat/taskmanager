package repository

import (
	"crud-api/internal/config"
	"crud-api/internal/models"
)

func GetTasksByUserID(userID int) ([]models.Task, error) {
	rows, err := config.DB.Query("SELECT id, title, description, is_completed FROM tasks WHERE user_id = $1 ORDER BY id DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetTaskByIDAndUserID(taskID string, userID int) (*models.Task, error) {
	var task models.Task
	err := config.DB.QueryRow("SELECT id, title, description, is_completed FROM tasks WHERE id = $1 AND user_id = $2", taskID, userID).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func CreateTask(userID int, title, description string) error {
	_, err := config.DB.Exec("INSERT INTO tasks (user_id, title, description) VALUES ($1, $2, $3)", userID, title, description)
	return err
}

func UpdateTask(taskID string, userID int, title, description string) error {
	_, err := config.DB.Exec("UPDATE tasks SET title = $1, description = $2 WHERE id = $3 AND user_id = $4", title, description, taskID, userID)
	return err
}

func ToggleTaskCompletion(taskID string, userID int) error {
	_, err := config.DB.Exec("UPDATE tasks SET is_completed = NOT is_completed WHERE id = $1 AND user_id = $2", taskID, userID)
	return err
}

func DeleteTask(taskID string, userID int) error {
	_, err := config.DB.Exec("DELETE FROM tasks WHERE id = $1 AND user_id = $2", taskID, userID)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}