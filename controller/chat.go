package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		var data Data
		data.Name = r.FormValue("email")
		data.Email = "claudio"
		w.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Write(js)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
