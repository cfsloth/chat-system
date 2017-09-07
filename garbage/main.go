package main

import (
	"Chat-System/chat-system/model"
)

func main() {
	session := model.InitializeDB()
	model.InsertFriendRequest(session, "filip", "claudio", "claudio@gmail.com")
}
