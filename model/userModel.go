package model

import(
	"gopkg.in/mgo.v2"
)

type User struct {
	EMAIL string
	PASSWORD string
	NAME string
}

func InsertUser(session *mgo.Session, email string, password string, name string){
	c := session.DB("chat-service").C("users")
	err := c.Insert(&User{email,password,name})
	if err != nil {
		panic(err)
	}
}
