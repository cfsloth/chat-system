package main

import (
	"Chat-System/chat-system/model"
)

func main() {
	session := model.InitializeDB()
	model.InsertFriendRequest(session, "claudio", "964122904", "claudiofilipesilvagoncalves@gmail.com")
	model.InsertFriendRequest(session, "tiago", "919227679", "claudiofilipesilvagoncalves@gmail.com")
	model.InsertFriendRequest(session, "filipe", "967793796", "claudiofilipesilvagoncalves@gmail.com")
}
