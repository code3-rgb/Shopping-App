package routes

import (
	"html/template"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func Server() {
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Server Error: ", err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	p := person{"John"}

	t, _ := template.ParseFiles("./static/index.htm")
	t.Execute(w, p)
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./static/hello.htm")
	t.Execute(w, nil)
}

func Router() {
	router := http.HandleFunc

	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	router("/", home)
	router("/hello", hello)

}
