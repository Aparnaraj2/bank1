package main

import (
	"fmt"
	"net/http"

	"text/template"

)

var tmp = template.Must(template.ParseGlob("frontend/*"))

type User struct {
	Cid         string
	Name        string
	MobileNo    string
	Gender      string
	Pancard     string
	Nationality string
	Address     string
	Username    string
	Password    string
}

/*
var tmpl template.Template
var db *sql.DB

	func getMySQLDB() *sql.DB {
		db, err := sql.Open("mysql", "root:Akhila@123@tcp(127.0.0.1:3306)/customer?parseTime=true")
		if err != nil {
			log.Fatal(err)
		}
		return db
	}

/*

		if r.FormValue("submit") == "insert" {
			id, _ := strconv.Atoi(userinfo.Cid)
			_, err := db.Exec("insert into user (cid,name,mobileno,gender,pancard,nationality,address,username,password) values(?,?,?,?,?,?,?,?,?)", id, userinfo.Name, userinfo.MobileNo, userinfo.Pancard, userinfo.Nationality, userinfo.Address, userinfo.Gender, userinfo.Nationality, userinfo.Username, userinfo.Password)
			if err != nil {
				tmpl.Execute(w, struct {
					Success bool
					Message string
				}{Success: true, Message: err.Error()})
			} else {
				tmpl.Execute(w, struct {
					Success bool
					Message string
				}{Success: true, Message: "Account created"})
			}

		}
		tmp.ExecuteTemplate(w, "register.html", nil)
	}
*/
func main() {

	http.HandleFunc("/", Bank)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/index", Index)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	var k User

	//v := [] string{"aa","bb"}
	tmp.ExecuteTemplate(w, "index.html", k)
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		v := "This is post request"
		tmp.ExecuteTemplate(w, "login.html", v)
		return
	}
	tmp.ExecuteTemplate(w, "login.html", nil)

}
func Bank(w http.ResponseWriter, r *http.Request) {

	tmp.ExecuteTemplate(w, "bank.html", nil)
}
func Register(w http.ResponseWriter, r *http.Request) {

	var J User
	if r.Method == "POST" {

		J.Cid = r.FormValue("cid")
		J.Name = r.FormValue("name")
		J.MobileNo = r.FormValue("mobileno")
		J.Pancard = r.FormValue("pancard")
		J.Nationality = r.FormValue("nationality")
		J.Address = r.FormValue("address")
		J.Username = r.FormValue("username")
		J.Password = r.FormValue("password")
		fmt.Println(J.Address)
		tmp.ExecuteTemplate(w, "test.html", J)
		return
	}
	tmp.ExecuteTemplate(w, "register.html", nil)

}
