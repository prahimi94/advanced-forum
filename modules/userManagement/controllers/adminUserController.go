package controller

import (
	"fmt"
	errorManagementControllers "forum/modules/errorManagement/controllers"
	"forum/modules/userManagement/models"
	"forum/utils"
	"net/http"
	"text/template"
)

// AdminReadAllUsers retrieves and displays all users for admin purposes.
func AdminReadAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorManagementControllers.HandleErrorPage(w, r, errorManagementControllers.MethodNotAllowedError)
		return
	}

	users, err := models.ReadAllUsers()
	if err != nil {
		errorManagementControllers.HandleErrorPage(w, r, errorManagementControllers.InternalServerError)
		return
	}

	// Create a template with a function map
	tmpl, err := template.New("users.html").Funcs(template.FuncMap{
		"formatDate": utils.FormatDate, // Register function globally
	}).ParseFiles(
		publicUrl + "users.html",
	)
	if err != nil {
		errorManagementControllers.HandleErrorPage(w, r, errorManagementControllers.InternalServerError)
		return
	}

	err = tmpl.Execute(w, users)
	if err != nil {
		errorManagementControllers.HandleErrorPage(w, r, errorManagementControllers.InternalServerError)
		return
	}
}

// AdminUpdateUser handles updating a user's status or details.
func AdminUpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extract parameters, validate admin rights, and update using the User model.
	fmt.Fprintln(w, "Admin: User updated")
}

// AdminDeleteUser handles user deletion.
func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	// Use the model to delete the user.
	fmt.Fprintln(w, "Admin: User deleted")
}

func RedirectToAdminIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin/", http.StatusFound)
}
