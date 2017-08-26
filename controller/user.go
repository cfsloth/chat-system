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

//Does the login - need to do session variables
func LoginUser(w http.ResponseWriter, r *http.Request){
	session := model.InitializeDB()
	user := model.FindUserByEmail(session,r.PostFormValue("email"))
	if r.PostFormValue("email") == user.EMAIL && r.PostFormValue("password") == user.PASSWORD {
			//Insert model method to acess database
			fmt.Println("Login sucessFull")
			http.Redirect(w, r, "/chatMessage", http.StatusSeeOther)
		}else{
			//Handling error
			var html = []byte(`
			<script> alert('Error : Wrong login') 
			window.location.href = "/"
			</script>`)
			w.Write(html)
		}
}

//Registe user
func RegisterUser(w http.ResponseWriter, r *http.Request){
	session := model.InitializeDB()
	user := model.FindUserByEmail(session,r.PostFormValue("email"))
	if user.EMAIL == "" {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		//Send to the database
		model.InsertUser(session,email,password,name)
	}else {
		//Handling error
		var html = []byte(`
		<script>
		alert("Error: Email already exists")
		window.location.href = "/"
		</script>`)
		w.Write(html)
	}
}

