package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/master.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	tmpl.Execute(w, nil)
}

func features(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/master.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	tmpl, err = tmpl.ParseFiles("pages/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	tmpl.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/master.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	tmpl, err = tmpl.ParseFiles("pages/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	tmpl.Execute(w, nil)
}
