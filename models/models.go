package models

import (

	//	"log"
	"time"
	//	"errors"
	"github.com/gothello/go-encurtador/db"
	"gopkg.in/mgo.v2/bson"
)

type Link struct {
	ID        bson.ObjectId `bson:"_id"`
	Code      string        `bson:"code"`
	Url       string        `bson:"url"`
	CreatedAt time.Time     `bson:"createdAt"`
}

func Create(link Link) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	err = db.C("links").Insert(&link)
	if err != nil {
		return err
	}

	return nil
}

func GetByCode(code string) (Link, error) {
	var link Link

	db, err := db.Connect()
	if err != nil {
		return link, err
	}

	err = db.C("links").Find(bson.M{"code": code}).One(&link)
	if err != nil {
		return link, err
	}

	return link, err
}

func GetAll() ([]Link, error) {
	var link []Link

	db, err := db.Connect()
	if err != nil {
		return link, err
	}

	err = db.C("links").Find(bson.M{}).All(&link)
	if err != nil {
		return link, err
	}

	return link, nil
}
