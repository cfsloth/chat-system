package main

import ( 
	"fmt"
	"net/http"
	"html/template"
	"Chat-System/chat-system/controller"
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
	fmt.Fprintf(w,"Hello Go");
}

func main(){
	//When in root file redirect to login or chat page
	//Serve css and js and sass
	http.Handle("./template", http.FileServer(http.Dir("./template")))
	http.HandleFunc("/login",user.LoginUser);
	http.ListenAndServe(":8080",nil)
}
