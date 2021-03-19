package db

import (

	
	mgo "gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"github.com/gothello/go-encurtador/config"
)

func Connect() (*mgo.Database, error) {
	conf, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("Error to load config databases: %s", err)
	}

//	mongoUser := conf.Get("mongouser")
	db := conf.GetString("database")
	mongohost := conf.GetString("mongohost")

	session, err := mgo.Dial(mongohost)
	if err != nil {
		return nil, err
	}


	return session.DB(db), nil
} 