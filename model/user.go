package userModel

import (
	"gopkg.in/mgo.v2"
	"time"
)

const(
	MongoDBHosts = "ds159953.mlab.com:59953"
	AuthDatabase = "chat-service"
	AuthUserName = "claudiofilipe21"
	AuthPassword = "claudiofilipe21"
)

//Attributes must be upper case or it will not save on the database
type Users struct {
	EMAIL string
	PASSWORD string
	NAME string
}

func InitializeDB(){

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs: []string{MongoDBHosts},
		Timeout: 600 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}

	c := session.DB("chat-service").C("users")
	err = c.Insert(&Users{"claudio@gmail.com","123456","claudio"})
	if err != nil {
		panic(err)
	}
}

