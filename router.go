package main

import ( 
	"fmt"
	"net/http"
	"html/template"
	"Chat-System/chat-system/controller"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_ := template.ParseFiles("./template/index.gtpl")
		t.Execute(w,nil)
	}else{
		r.ParseForm()
//		Acess to Database
		if r.PostFormValue("username") == "claudio" && r.PostFormValue("password") == "admin" {
			fmt.Println("claudio");
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./template/index.gtpl");
}

func chatService(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./template/chat.gtpl")
}

func main(){
	r := mux.NewRouter()
	cssHandler := http.FileServer(http.Dir("./template/css/"))
	jsHandler := http.FileServer(http.Dir("./template/js/"))
	sassHandler := http.FileServer(http.Dir("./template/sass/"))
	http.Handle("/css/",http.StripPrefix("/css/", cssHandler))
	http.Handle("/js/",http.StripPrefix("/js/",jsHandler))
	http.Handle("/sass/",http.StripPrefix("/sass/",sassHandler))
	r.HandleFunc("/",user.LoginUser)
	r.HandleFunc("/chatMessage",chatService)
	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)
}
