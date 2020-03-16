package routes

import "github.com/gorilla/mux"

// NewRouter ...
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeGetHandler).Methods("GET")
	return r
}
