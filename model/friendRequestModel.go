package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FriendRequest struct {
	//FROM and TO are emails
	FROMNAME  string
	FROMEMAIL string
	TOEMAIL   string
}

func InsertFriendRequest(session *mgo.Session, fromName string, fromEmail string, toEmail string) {
	c := session.DB("chat-service").C("friendRequests")
	err := c.Insert(&FriendRequest{fromName, fromEmail, toEmail})
	if err != nil {
		panic(err)
	}
}

func FindFriendRequests(session *mgo.Session, email string) []FriendRequest {
	var result []FriendRequest
	c := session.DB("chat-service").C("friendRequests")
	c.Find(bson.M{"toemail": email}).All(&result)
	return result
}
