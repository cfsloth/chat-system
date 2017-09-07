package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name string
}

func LoadFriendsRequest(w http.ResponseWriter, r *http.Request) {
	data := Data{"World!"}
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(js)
}
