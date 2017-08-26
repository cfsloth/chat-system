package main

import(
	"fmt"
	"Chat-System/chat-system/model"
)

func main(){
	session := model.InitializeDB()
	model.InsertUser(session,"123","1321","13123")
	u := model.FindUser(session,"claudio@gmail.com")
	fmt.Println(u.NAME)
	fmt.Println(u.PASSWORD)
}
