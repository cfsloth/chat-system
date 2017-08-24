package main

import ( 
	"fmt"
	"net/http"
	"html/template"
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
	//Serve css and js and sass
	http.FileServer(http.Dir("./template"))
	http.HandleFunc("/",handler);
	http.ListenAndServe(":8080",nil)
}


