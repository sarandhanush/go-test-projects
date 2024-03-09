package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

// User represents a simple user structure
type User struct {
	Username string
	Password string
}

// In-memory storage for user credentials
var users = map[string]User{
	"user1": {"user1", "password1"},
	"user2": {"user2", "password2"},
}

// Handler for the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if the provided credentials are valid
		if user, ok := users[username]; ok && user.Password == password {
			// Successful login, you can set a session or cookie here
			fmt.Fprintf(w, "Welcome, %s!", username)
			return
		}

		// Invalid credentials, show an error message
		fmt.Fprintln(w, "Invalid username or password")
		return
	}

	// Display the login form
	tmpl := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Login</title>
		</head>
		<body>
			<h1>Login</h1>
			<form method="post" action="/login">
				<label for="username">Username:</label>
				<input type="text" id="username" name="username" required><br>
				<label for="password">Password:</label>
				<input type="password" id="password" name="password" required><br>
				<button type="submit">Login</button>
			</form>
		</body>
		</html>
	`

	// Render the HTML login form
	tmplParsed, err := template.New("login").Parse(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmplParsed.Execute(w, nil)
}
