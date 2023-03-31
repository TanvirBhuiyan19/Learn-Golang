package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hosting_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	// defer db.Close()

	fmt.Printf("Connection is established")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
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

func contact(w http.ResponseWriter, r *http.Request) {

	//Method-1
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// fmt.Printf(" Successfully submitted!\r\n\r\n Your name is: %s\r\n Your company is: %s\r\n Your email is: %s", name, company, email)
	// fmt.Fprintf(w, " Successfully submitted!\r\n\r\n Your name is: %s\r\n Your company is: %s\r\n Your email is: %s", name, company, email)

	//Method-2 (Recommended)
	// r.ParseForm()
	// for key, value := range r.Form {
	// 	fmt.Printf("Your %s is: %s\r\n", key, value[0])
	// }
	// fmt.Fprintf(w, "Successfully submitted!")

	// perform a db.Query insert
	qs := "INSERT INTO `contact` (`id`, `name`, `email`, `company`, `status`, `created_at`, `updated_at`) VALUES (NULL, '%s', '%s', '%s', '1', '%s', NULL);"
	sql := fmt.Sprintf(qs, name, email, company, currentTime)

	insert, err := db.Query(sql)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	fmt.Fprintf(w, "Successfully submitted!")

}
