package handlers

import (
	"crud-api/internal/middleware"
	"crud-api/internal/models"
	"crud-api/internal/repository"
	"crud-api/internal/services"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

func LoginPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"CSRFToken": middleware.GetOrSetCSRFToken(w, r),
	}
	if r.URL.Query().Get("error") == "1" {
		data["Error"] = "Invalid credentials"
	}
	tmpl.ExecuteTemplate(w, "login.html", data)
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	_, err := services.AuthenticateUser(username, password)
	if err != nil {
		http.Redirect(w, r, "/?error=1", http.StatusSeeOther)
		return
	}

	// login success, set cookie
	http.SetCookie(w, &http.Cookie{Name: "user", Value: username, Path: "/", HttpOnly: true})
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "user", Value: "", Path: "/", MaxAge: -1, HttpOnly: true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DashboardPage(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tasks, err := repository.GetTasksByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := models.DashboardData{Tasks: tasks, CSRFToken: middleware.GetOrSetCSRFToken(w, r)}
	editID := r.URL.Query().Get("edit")
	if editID != "" {
		task, err := repository.GetTaskByIDAndUserID(editID, userID)
		if err == nil {
			data.EditTask = task
		}
	}

	tmpl.ExecuteTemplate(w, "dashboard.html", data)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		title := r.FormValue("title")
		desc := r.FormValue("description")
		err := repository.CreateTask(userID, title, desc)
		if err != nil {
			http.Error(w, "Failed to create", 500)
			return
		}
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		id := r.FormValue("id")
		title := r.FormValue("title")
		desc := r.FormValue("description")
		err := repository.UpdateTask(id, userID, title, desc)
		if err != nil {
			http.Error(w, "Failed to update", 500)
			return
		}
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func ToggleTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userID, err := middleware.GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id := r.FormValue("id")
	err = repository.ToggleTaskCompletion(id, userID)
	if err != nil {
		http.Error(w, "Update failed", 500)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userID, err := middleware.GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id := r.FormValue("id")
	err = repository.DeleteTask(id, userID)
	if err != nil {
		http.Error(w, "Delete failed", 500)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}