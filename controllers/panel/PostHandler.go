package panel

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-simple-crud/helpers"
	"go-simple-crud/models"
	"html/template"
	"net/http"
	"strconv"
)

type Posts struct{}

func (Posts) IndexHandler(w http.ResponseWriter, r *http.Request) {

	view, err := template.ParseFiles(helpers.PageDirectories("panel/pages/posts/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	view.ExecuteTemplate(w, "posts", data)

}

func (Posts) CreateHandler(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles(helpers.PageDirectories("panel/pages/posts/create")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "postCreate", nil)
}

func (Posts) StoreHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	t := r.FormValue("title")
	c := r.FormValue("content")
	d := r.FormValue("description")

	models.Post{}.Add(t, c, d)

	http.Redirect(w, r, "/panel/posts", http.StatusSeeOther)
}

func (Posts) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	models.Post{}.Delete(vars["id"])

	http.Redirect(w, r, "/panel/posts", http.StatusSeeOther)
}

func (Posts) EditHandler(w http.ResponseWriter, r *http.Request)  {

	view, err := template.ParseFiles(helpers.PageDirectories("panel/pages/posts/edit")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	params := mux.Vars(r)
	data := models.Post{}.Get(params["id"])

	view.ExecuteTemplate(w, "postEdit", data)
}

func (Posts) UpdateHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()

	params := mux.Vars(r)

	t := r.FormValue("title")
	c := r.FormValue("content")
	d := r.FormValue("description")
	i,_ := strconv.Atoi(params["id"])

	models.Post{}.Update(t, c, d, uint(i))

	http.Redirect(w, r, "/panel/posts", http.StatusSeeOther)
}
