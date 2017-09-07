package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FriendRequest struct {
	//FROM and TO are emails
	FROMNAME  string
	CELLPHONE string
	TOEMAIL   string
}

func InsertFriendRequest(session *mgo.Session, fromName string, cellPhone string, toEmail string) {
	c := session.DB("chat-service").C("friendRequests")
	err := c.Insert(&FriendRequest{fromName, cellPhone, toEmail})
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
