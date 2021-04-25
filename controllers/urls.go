package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gothello/go-encurtador/models"
	"github.com/gothello/go-encurtador/utils"
	"gopkg.in/mgo.v2/bson"
)

var (
	codes = map[string]string{}

	seeders = []rune("abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ123456789")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GenerateCode(lenght int, oldUrl string) (models.Link, error) {
	result := make([]rune, lenght)
	var l models.Link
	uri, err := url.ParseRequestURI(oldUrl)
	if err != nil {
		return l, err
	}

	for {

		for i := range result {
			rand.Seed(time.Now().UTC().UnixNano())
			result[i] = seeders[rand.Intn(len(seeders))]
		}

		if _, err := models.GetByCode(string(result)); err != nil {

			l = models.Link{
				ID:        bson.NewObjectId(),
				Code:      string(result),
				Url:       uri.String(),
				CreatedAt: time.Now(),
			}

			err := models.Create(l)
			if err != nil {
				return l, err
			}

			break
		}

		continue
	}

	return l, nil
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	hash := keys.Get("code")
	log.Println(hash)

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
	var l models.Link
	var err error

	if err = json.NewDecoder(r.Body).Decode(&j); err != nil {
		utils.ToError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	l, err = GenerateCode(7, j.Url)
	fmt.Println(l)

	utils.ToJson(w, map[string]string{
		"code": l.Code,
	}, http.StatusCreated)
}
