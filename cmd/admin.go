package main

import (
	"fmt"
	"html/template"
	"main/pkg/services"
	"net/http"
)

type AdminHanlder struct {
	Service *services.Service
}

func NewAdminandler(Service *services.Service) *AdminHanlder {
	return &AdminHanlder{
		Service: Service,
	}
}

func (a *AdminHanlder) ViewUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.Service.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}

	file := "./ui/templates/adminPage.html"
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		http.Error(w, "Error parsing templates", 500)
		return
	}

	err = tmpl.Execute(w, users)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error executing template", 500)
		return
	}
}

func (a *AdminHanlder) PromoteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.FormValue("id")
	err = a.Service.UserService.UpdateUserRole(userID, "teacher")
	if err != nil {
		http.Error(w, "Error promoting user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (a *AdminHanlder) DemoteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.FormValue("id")
	err = a.Service.UserService.UpdateUserRole(userID, "student")
	if err != nil {
		http.Error(w, "Error demoting user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (a *AdminHanlder) Ban(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.FormValue("id")
	err = a.Service.UserService.UpdateUserRole(userID, "banned")
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (a *AdminHanlder) UNBan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.FormValue("id")
	err = a.Service.UserService.UpdateUserRole(userID, "student")
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
