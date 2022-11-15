package routes

import (
	"fmt"
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

type people struct {
	Name            string `json:"Name"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	PasswordConfirm string `json:"PasswordConfirm"`
}

func form(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "<h1>Welcome to you!!!</h1>", r.Form.Get("Name"))

	var p = people{
		Name:            r.FormValue("Name"),
		Email:           r.FormValue("Email"),
		Password:        r.FormValue("Password"),
		PasswordConfirm: r.FormValue("PasswordConfirm"),
	}

	fmt.Println(p.Name)

}

func Router() {
	router := http.HandleFunc

	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	router("/", home)
	router("/hello", hello)
	router("/form", form)

}
