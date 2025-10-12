package main

import (
	"fmt"
	"github.com/turman17/mini-rest/internal/users"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", users.HandleRoot)

	mux.HandleFunc("POST /users", users.CreateUser)
	mux.HandleFunc("GET /users/{id}", users.GetUser)
	mux.HandleFunc("DELETE /users/{id}", users.DeleteUser)

	fmt.Println("Server listening to :8080")
	http.ListenAndServe(":8080", mux)
}
