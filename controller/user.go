package controller

import (
	"fmt"
	"net/http"
	"Chat-System/chat-system/model"
)

//Controls the page
func LoadPageAndMethods(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		http.ServeFile(w,r,"template/index.gtpl")
		fmt.Println("GET /")
	}else{
		r.ParseForm()
		if r.PostFormValue("name") != "" {
			fmt.Println("POST /registe")
			RegisterUser(w,r)
		}else{
			fmt.Println("POST /login")
			LoginUser(w,r)
		}
	}
}

//Does the login - incomplete
func LoginUser(w http.ResponseWriter, r *http.Request){
		if r.PostFormValue("email") == "claudio" && r.PostFormValue("password") == "admin" {
			//Insert model method to acess database
			fmt.Println("Login sucessFull")
			http.Redirect(w, r, "/chatMessage", http.StatusSeeOther)
		}else{
			panic("Wrong login")
		}
}

//Registe user - incomplete
func RegisterUser(w http.ResponseWriter, r *http.Request){
	var(
		registered bool = false;
	)
	if !registered {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		session := model.InitializeDB()
		model.InsertUser(session,email,password,name)
		//Send to database
	}else {
		panic("Error user already exists")
	}
}
