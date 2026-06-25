package app

import (
	"log"
	"net/http"
	"task_manager/constants"
)

func Server(port string) {
	log.Println("Server is running on Port:", constants.PORT)

	RegisterRouter()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func EnsureRequestMethod(method string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != method {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}
