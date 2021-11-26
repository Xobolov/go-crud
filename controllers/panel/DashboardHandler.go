package panel

import (
	"fmt"
	"go-simple-crud/helpers"
	"html/template"
	"net/http"
)

type Dashboard struct {}

func (Dashboard) DashboardHandler(w http.ResponseWriter, r *http.Request)  {

	view, err := template.ParseFiles(helpers.PageDirectories("panel/pages/dashboard")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "dashboard", nil)

}