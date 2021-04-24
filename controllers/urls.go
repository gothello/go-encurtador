package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/mgo.v2/bson"

	//	"strings"
	//	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gothello/go-encurtador/models"
	"github.com/gothello/go-encurtador/utils"
)

var (
	codes = map[string]string{}

	seeders = []rune("abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ123456789")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GenerateCode(lenght int, url string) string {
	h := make([]rune, lenght)

	for {

		for i := range h {
			rand.Seed(time.Now().UTC().UnixNano())
			h[i] = seeders[rand.Intn(len(seeders))]
		}

		if _, ok := codes[string(h)]; !ok {

			codes[string(h)] = url

			return string(h)
		}
	}

	return string(h)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	hash := keys.Get("code")
	log.Println(hash)

	/*	doc, err := models.FindOne(hash)
		if err != nil {
			utils.ToError(w, err.Error(), http.StatusNotFound)
		}
	*/

	for code, url := range codes {
		if hash == code {
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		}
	}
}

func Urls(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ToError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type Url struct {
		Url string `json:"url"`
	}

	var j Url

	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		utils.ToError(w, err.Error(), http.StatusInternalServerError)
	}

	uri, err := url.ParseRequestURI(j.Url)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
	}

	_ = uri

	code := GenerateCode(7, j.Url)

	p := models.Link{
		ID:        bson.NewObjectId(),
		Code:      "TESTE",
		Url:       "google.com",
		CreatedAt: time.Now(),
	}

	err = models.Create(p)
	if err != nil {
		log.Println("tem erro na funcao create burro.", err)
	}

	o, err := models.GetByCode("TESTE")
	if err != nil {
		log.Println("erro na funcao getbycode", err)
	}

	fmt.Println("result", o)

	l, err := models.GetAll()
	if err != nil {
		log.Println("Erro na funcao get all", err)

	}

	fmt.Printf("%+v\n", l)

	/*

		link := models.Link{
			ID: primitive.NewObjectID(),
			Code: code,
			Url: uri.String(),
			CreatedAt: time.Now(),
		}

		if err := models.Insert(link); err != nil {
			utils.ToError(w, err.Error(), http.StatusBadRequest)
		}
	*/

	utils.ToJson(w, map[string]string{
		"url": "http://localhost:3000/api/" + code,
	}, http.
		StatusCreated)
}
