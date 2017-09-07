package model

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FriendRequest struct {
	//FROM and TO are emails
	FROM string
	TO string
}

func InsertFriendRequest(session *mgo.Session,from string,to string){
	c := session.DB("chat-service").C("friendRequests")
	err := c.Insert(&FriendRequest{from,to})
	if err != nil {
		panic(err)
	}
}

func FindFriendRequests(session *mgo.Session, email string) *mgo.Query{
	c := session.DB("chat-service").C("friendRequests")
	array := c.Find(bson.M{"to":email})
	return array
}

