package models

import (


//	"log"
	"time"
//	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/gothello/go-encurtador/db"
)


type Link struct {
	ID        bson.ObjectId 	 `bson:"_id"`
	Code      string             `bson:"code"`
	Url       string             `bson:"url"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func Create(link Link) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}


	err = db.C("links").Insert(&link)
	if err  != nil {
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

/*
func Insert(link Link) error {
	var linksCollection *LinksCollection

	_, err := linksCollection.links.InsertOne(context.TODO(), link)
	if err != nil {
		return err
	}

	return nil
}

func FindOne(code string) (Link, error){
	var doc Link
	var linksCollection *LinksCollection

	err := linksCollection.links.FindOne(context.TODO(), bson.M{"code": code}).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments{

			return doc, errors.New("No documents")
		}

		return doc, err
	}

	return doc, nil
}
*/