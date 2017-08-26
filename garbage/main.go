package main

import(
	"Chat-System/chat-system/model"
)

func main(){
	session := model.InitializeDB()
	model.InsertUser(session,"123","1321","13123")
}
