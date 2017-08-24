package user

import (
	"fmt"
	"net/http"
	"html/template"
)

//Probable turn this into a login function
func LoginUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		t,_ := template.ParseFiles("template/index.gtpl")
		t.Execute(w,nil)
	}else{
		r.ParseForm()
		if r.PostFormValue("username") == "claudio" && r.PostFormValue("Password") == "admin" {
			//Insert model method to acess database
			fmt.Println("Login sucessFull")
		}
	}
}

//Incomplete implement database first
func RegisterUser(w http.ResponseWriter, r *http.Request){
	bool registered = false;
	if !registered {
		r.ParseForm()
		string name = r.PostFormValue("name")
		string password = r.PostFormValue("password")
		string email = r.PostFormValue("email")
		//Send to database
	}else {
		panic("Error user already exists")
	}
}
