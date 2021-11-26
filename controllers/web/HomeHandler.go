package web

import (
	"fmt"
	"go-simple-crud/helpers"
	"go-simple-crud/models"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	view, err := template.ParseFiles(helpers.PageDirectories("web")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	posts := models.Post{}.GetAll()


	data := make(map[string]interface{})
	data["Posts"] = posts

	view.ExecuteTemplate(w, "index", data)

}














