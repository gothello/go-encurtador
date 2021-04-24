package db

import (
	"fmt"

	"github.com/gothello/go-encurtador/config"
	mgo "gopkg.in/mgo.v2"
)

func Connect() (*mgo.Database, error) {
	conf, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("Error to load config databases: %s", err)
	}

	db := conf.GetString("database")
	mongohost := conf.GetString("mongohost")

	session, err := mgo.Dial(mongohost)
	if err != nil {
		return nil, err
	}

	return session.DB(db), nil
}
