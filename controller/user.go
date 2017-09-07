package controller

import (
	"Chat-System/chat-system/model"
	"fmt"
	"net/http"
	"time"
)

func LoadPageAndMethods(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var cookie, err = r.Cookie("loginC")
		//Here use some encriptation/desincriptation function for cookies
		if err == nil && cookie.Value == "true" {
			http.Redirect(w, r, "/chatMessage", http.StatusSeeOther)
		} else {
			http.ServeFile(w, r, "template/index.gtpl")
			fmt.Println("GET /")
		}
	} else {
		r.ParseForm()
		if r.PostFormValue("name") != "" {
			fmt.Println("POST /register")
			RegisterUser(w, r)
		} else {
			fmt.Println("POST /login")
			LoginUser(w, r)
		}
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	session := model.InitializeDB()
	user := model.FindUserByEmail(session, r.PostFormValue("email"))
	if r.PostFormValue("email") == user.EMAIL && r.PostFormValue("password") == user.PASSWORD {
		//Insert model method to acess database
		fmt.Println("Login sucessFull")
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "loginC", Value: "true", Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/chatMessage", http.StatusSeeOther)
	} else {
		//Handling error
		var html = []byte(`
			<script> alert('Error : Wrong login')
			window.location.href = "/"
			</script>`)
		w.Write(html)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	session := model.InitializeDB()
	user := model.FindUserByEmail(session, r.PostFormValue("email"))
	if user.EMAIL == "" {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		//Send to the database
		model.InsertUser(session, email, password, name)
	} else {
		//Handling error
		var html = []byte(`
		<script>
		alert("Error: Email already exists")
		window.location.href = "/"
		</script>`)
		w.Write(html)
	}
}
