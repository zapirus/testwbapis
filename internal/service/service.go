package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zapirus/testwbapis/internal/model"
	"strconv"

	"net/http"
)

var shops []model.Shop
var users []model.User

func strip(url string) string {
	var res string
	var item int

	for _, i2 := range url {
		if item == 2 {
			break
		} else if i2 == 47 {
			item += 1
		}
		res += string(i2)
	}
	return res
}

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

	} else if r.Method == "PUT" && strip(r.URL.RequestURI()) == "/changeuser/" {
		var reqId = mux.Vars(r)["id"]
		id, err := strconv.Atoi(reqId)

		if err != nil {
			w.Write([]byte("Не удалось сконвертировать"))
			return nil, nil
		}
		if id >= len(users) {
			w.Write([]byte("Нет такого поста"))
			return nil, nil
		}
		var changeUser model.User
		json.NewDecoder(r.Body).Decode(&changeUser)
		users[id] = changeUser
		return users, nil

	} else if r.Method == "PUT" && strip(r.URL.RequestURI()) == "/changeshop/" {
		var reqId = mux.Vars(r)["id"]
		id, err := strconv.Atoi(reqId)
		if err != nil {
			w.Write([]byte("Не удалось сконвертировать"))
			return nil, nil
		}
		if id >= len(users) {
			w.Write([]byte("Нет такого поста"))
			return nil, nil
		}
		var changeUser model.User
		json.NewDecoder(r.Body).Decode(&changeUser)
		users[id] = changeUser
		return users, nil

	} else if r.Method == "DELETE" && strip(r.URL.RequestURI()) == "/changeuser/" {
		var reqId = mux.Vars(r)["id"]
		id, err := strconv.Atoi(reqId)
		if err != nil {
			w.Write([]byte("Не удалось сконвертировать"))
			return nil, nil
		}
		if id >= len(users) {
			w.Write([]byte("Нет такого поста"))
			return nil, nil
		}
		users = append(users[:id], users[id+1:]...)
		return users, nil

	} else if r.Method == "DELETE" && strip(r.URL.RequestURI()) == "/changeshop/" {
		var reqId = mux.Vars(r)["id"]
		id, err := strconv.Atoi(reqId)
		if err != nil {
			w.Write([]byte("Не удалось сконвертировать"))
			return nil, nil
		}
		if id >= len(users) {
			w.Write([]byte("Нет такого поста"))
			return nil, nil
		}
		shops = append(shops[:id], shops[id+1:]...)
		return nil, shops

	}

	w.Write([]byte("Ничего не нашлось. Сорян\n"))
	return nil, nil
}
