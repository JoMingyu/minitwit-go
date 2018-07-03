package config

import (
	"log"

	"github.com/globalsign/mgo"
)

var mongodbSession *mgo.Session
var dbName = "minitwit-go"
var (
	UserCol   = DB().C("users")
	FollowCol = DB().C("follows")
)

// DBSession returns the current db session.
func DBSession() *mgo.Session {
	if mongodbSession != nil {
		return mongodbSession
	}

	uri := "mongodb://localhost"

	di, err := mgo.ParseURL(uri)
	if err != nil {
		log.Fatalf("Can't parse mongo uri, go error %v\n", err)
	}

	mongodbSession, err = mgo.DialWithInfo(di)
	if mongodbSession == nil || err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
	}

	mongodbSession.SetSafe(&mgo.Safe{})

	return mongodbSession
}

// DB returns a database given a name.
func DB() *mgo.Database {
	return DBSession().DB(dbName)
}

// AddBasicIndex add a ascending index given a list of `keys`. The index is always built in background.
func AddBasicIndex(collection *mgo.Collection, keys ...string) {
	collection.EnsureIndex(mgo.Index{
		Key:        keys,
		Background: true,
	})
}
