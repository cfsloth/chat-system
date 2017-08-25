package user

import (
	"fmt"
	"net/http"
)

//Probable turn this into a login function
func LoginUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		http.ServeFile(w,r,"template/index.gtpl")
		fmt.Println("GET /")
	}else{
		fmt.Println("POST /")
		r.ParseForm()
		if r.PostFormValue("username") == "claudio" && r.PostFormValue("password") == "admin" {
			//Insert model method to acess database
			fmt.Println("Login sucessFull")
			http.Redirect(w, r, "/chatMessage", http.StatusSeeOther)
		}
	}
}

//Incomplete implement database first
func RegisterUser(w http.ResponseWriter, r *http.Request){
	var(
		registered bool = false;
	)
	if !registered {
		r.ParseForm()
	//	string name = r.PostFormValue("name")
	//	string password = r.PostFormValue("password")
	//	string email = r.PostFormValue("email")
		//Send to database
	}else {
		panic("Error user already exists")
	}
}
