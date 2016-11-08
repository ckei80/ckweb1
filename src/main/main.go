package main

import (
	"html/template"
	"net/http"
	"os"
)

//Compile templates on start
var templates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/main.html", "templates/about.html"))

//A Page structure
type Page struct {
	Title string
}

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	fs := justFilesFilesystem{http.Dir("content/css/")}
	http.Handle("/content/css/", http.StripPrefix("/content/css/", http.FileServer(fs)))
	templates.ExecuteTemplate(w, tmpl, data)
}

//The handlers.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "main", &Page{Title: "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "about", &Page{Title: "About"})
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about", aboutHandler)

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
