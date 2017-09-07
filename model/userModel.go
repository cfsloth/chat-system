package model

import (
	//	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	EMAIL     string
	PASSWORD  string
	NAME      string
	CELLPHONE string
}

func InsertUser(session *mgo.Session, email string, password string, name string, cellPhone string) {
	c := session.DB("chat-service").C("users")
	err := c.Insert(&User{email, password, name, cellPhone})
	if err != nil {
		panic(err)
	}
}

func FindUserByEmail(session *mgo.Session, searchTerm string) User {
	c := session.DB("chat-service").C("users")
	user := User{}
	c.Find(bson.M{"email": searchTerm}).One(&user)
	return user
}
