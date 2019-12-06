package main

import (
	"net/http"

	"./controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	cssHandler := http.FileServer(http.Dir("./template/css/"))
	jsHandler := http.FileServer(http.Dir("./template/js/"))
	sassHandler := http.FileServer(http.Dir("./template/sass/"))
	images := http.FileServer(http.Dir("./template/images/"))
	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandler))
	http.Handle("/sass/", http.StripPrefix("/sass/", sassHandler))
	http.Handle("/images/", http.StripPrefix("/images/", images))
	r.HandleFunc("/", controller.LoadPageAndMethods)
	r.HandleFunc("/chatMessage", controller.LoadPageAndMethodsChat)
	r.HandleFunc("/findFriendsRequests", controller.LoadFriendsRequest)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
