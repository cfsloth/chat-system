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
		fmt.Println("Acess denied")
	}
}
