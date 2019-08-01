package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func DataBaseConn() (db *sql.DB) {
	driverDB := "mysql"
	userDB := "braxton"
	passDB := ""
	nameDB := "goblog"
	db, err := sql.Open(driverDB, userDB+":"+passDB+"@/"+nameDB)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

//Index func
func Index(w http.ResponseWriter, r *http.Request) {
	db := DataBaseConn()
	selectDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	response := []Employee{}
	for selectDB.Next() {
		var id int
		var name, city string
		err = selectDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.City = city
		response = append(response, employee)
	}
	tmpl.ExecuteTemplate(w, "Index", response)
	defer db.Close()
}

//Show func
func Show(w http.ResponseWriter, r *http.Request) {
	db := DataBaseConn()
	// this will get the id from the url, (i guess)
	nID := r.URL.Query().Get("id")
	fmt.Println(nID)
	selectDB, err := db.Query("SELECT * FROM Employee where id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	for selectDB.Next() {
		var id int
		var name, city string
		err = selectDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.City = city
	}
	tmpl.ExecuteTemplate(w, "Show", employee)
	db.Close()
}

// New Employee func
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Edit Employee func
func Edit(w http.ResponseWriter, r *http.Request) {
	db := DataBaseConn()
	nID := r.URL.Query().Get("id")
	selectDB, err := db.Query("SELECT * FROM Employee where id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	for selectDB.Next() {
		var id int
		var name, city string
		err = selectDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.City = city
	}
	tmpl.ExecuteTemplate(w, "Edit", employee)
	defer db.Close()
}

func main() {
	log.Println("Listening server on: http://localhost:3500")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.ListenAndServe(":3500", nil)
}
