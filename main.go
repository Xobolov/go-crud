package main

import (
	"go-simple-crud/routes"
	"net/http"
)

func main()  {

	port := ":8000"
	http.ListenAndServe(port, routes.Routes())
}

