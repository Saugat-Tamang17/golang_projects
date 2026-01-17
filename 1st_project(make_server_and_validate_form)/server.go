package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "1st_project(make_server_and_validate_form)/index.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "1st_project(make_server_and_validate_form)/form.html")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	name := r.Form.Get("name")
	age := r.Form.Get("age")

	if name == "" || age == "" {
		http.Error(w, "all fields are supposed to be filled", http.StatusBadRequest)
		return
	}

	fmt.Printf("New signup: %s (%s)\n", name, age)

	// Redirecting to /success route
	http.Redirect(w, r, "/success", http.StatusSeeOther)
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "1st_project(make_server_and_validate_form)/sucess.html")
}

func main() {
	fmt.Println("Server chai port 8080 ma start gardim ig")

	http.HandleFunc("/", homeHandler)         // home page
	http.HandleFunc("/form", formHandler)     // form page
	http.HandleFunc("/submit", submitHandler) // form POST
	http.HandleFunc("/success", successHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
