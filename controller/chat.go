package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
)

type Data struct {
	Name  string
	Email string
}

func LoadPageAndMethodsChat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var cookie, err = r.Cookie("loginC")
		//Here use some encriptation/desincriptation function for cookies
		if err == nil && cookie.Value == "true" {
			http.ServeFile(w, r, "template/chat.gtpl")
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Println("GET /chatMessage")
		}
	}
}

func LoadFriendsRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		session := model.InitializeDB()
		listOfRequests := model.FindFriendRequests(session, email)
		w.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(listOfRequests)
		if err != nil {
			panic(err)
		}
		w.Write(js)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
