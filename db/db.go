package db

import (

	"log"
	"time"
	"context"
	"go.mongo.org/mongo-driver/mongo"
	"go.mongo.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	conf := config.Load()

	mongoUser := conf.Get("mongouser")
	mongoPass := conf.Get("mongopass")
	mongoHost := conf.Get("mongoHost")

	credentials := options.Credentials{
		Username: mongoUser,
		password: mongoPass,
	}

	ctx, cancel := context.WithTimeout(context.Background, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(context.TODO(), options.Client().SetAuth(credentials).ApplyURI(mongoHost))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
} 