package main

import ( 
	"fmt"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("./template/index.html");
	t.Execute(w,"");
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Go");
}

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/chat",hello);
	http.ListenAndServe(":8080",nil)
}


