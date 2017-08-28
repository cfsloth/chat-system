package main

import ( 
	"net/http"
	"Chat-System/chat-system/controller"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./template/index.gtpl");
}

func chatService(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./template/chat-box.html")
}

func main(){
	r := mux.NewRouter()
	cssHandler := http.FileServer(http.Dir("./template/css/"))
	jsHandler := http.FileServer(http.Dir("./template/js/"))
	sassHandler := http.FileServer(http.Dir("./template/sass/"))
	http.Handle("/css/",http.StripPrefix("/css/", cssHandler))
	http.Handle("/js/",http.StripPrefix("/js/",jsHandler))
	http.Handle("/sass/",http.StripPrefix("/sass/",sassHandler))
	r.HandleFunc("/",controller.LoadPageAndMethods)
	r.HandleFunc("/chatMessage",chatService)
	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)
}
