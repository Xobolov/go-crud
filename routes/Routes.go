package routes

import (
	"github.com/gorilla/mux"
	"go-simple-crud/controllers/panel"
	"go-simple-crud/controllers/web"
	"net/http"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	// Home
	r.HandleFunc("/", web.HomeHandler).Methods("GET")

	// Dashboard
	admin := r.PathPrefix("/panel").Subrouter()
	admin.HandleFunc("/dashboard", panel.Dashboard{}.DashboardHandler).Methods("GET")

	// Posts
	posts := admin.PathPrefix("/posts").Subrouter()
	posts.HandleFunc("", panel.Posts{}.IndexHandler).Methods("GET")
	posts.HandleFunc("/create", panel.Posts{}.CreateHandler).Methods("GET")
	posts.HandleFunc("/store", panel.Posts{}.StoreHandler).Methods("POST")
	posts.HandleFunc("/edit/{id:[0-9]+}", panel.Posts{}.EditHandler).Methods("GET")
	posts.HandleFunc("/update/{id:[0-9]+}", panel.Posts{}.UpdateHandler).Methods("POST")
	posts.HandleFunc("/delete/{id}", panel.Posts{}.DeleteHandler).Methods("POST")



	//r.Use(helpers.LoggingMiddleware)



	fileServer := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", fileServer))

	return r
}



