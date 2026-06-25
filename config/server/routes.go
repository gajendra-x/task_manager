package app

import (
	"fmt"
	"net/http"
	"task_manager/modules/todo"
	"task_manager/modules/user"
)

func RegisterRouter() {
	// health
	http.Handle("/api/health", EnsureRequestMethod("GET")(http.HandlerFunc(healthCheck)))

	// user
	http.Handle("/api/user/add", EnsureRequestMethod("POST")(http.HandlerFunc(user.CreateUser)))
	http.Handle("/api/user/{email}", EnsureRequestMethod("GET")(http.HandlerFunc(user.GetUser)))

	// todo
	http.Handle("/api/todo/add", EnsureRequestMethod("POST")(http.HandlerFunc(todo.CreateTodo)))
	http.Handle("/api/todo/{user_id}", EnsureRequestMethod("GET")(http.HandlerFunc(todo.GetUsersTodo)))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Looks good :)")
}
