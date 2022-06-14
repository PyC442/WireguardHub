package router

import (
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles("./templates/index.html"))

func rootRoute(w http.ResponseWriter, r *http.Request) {
	templates.Execute(w, "index.html")
}


func RouterSetup(port string) {
	http.HandleFunc("/", rootRoute)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(port, nil)
}
