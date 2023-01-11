package service

import (
	"encoding/json"
	"github.com/zapirus/testwbapis/internal/model"

	"net/http"
)

var shops []model.Shop
var users []model.User

func UniversalFunc(w http.ResponseWriter, r *http.Request) ([]model.User, []model.Shop) {

	if r.Method == "POST" && r.URL.RequestURI() == "/user" {
		var newUser model.User
		json.NewDecoder(r.Body).Decode(&newUser)

		users = append(users, newUser)
		return users, nil

	} else if r.Method == "POST" && r.URL.RequestURI() == "/shop" {
		var newShop model.Shop
		json.NewDecoder(r.Body).Decode(&newShop)

		shops = append(shops, newShop)

		return nil, shops
	}
	return nil, nil

}
