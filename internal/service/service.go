package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zapirus/testwbapis/internal/model"
	"strconv"
	"strings"

	"net/http"
)

var (
	users []model.User
	shops []model.Shop
)

// Функция, которая режет URL
func strip(url string) string {
	var (
		res     string
		counter int
	)

	for _, i2 := range url {
		if counter == 2 {
			break
		} else if i2 == 47 {
			counter += 1
		}
		res += string(i2)
	}
	return res
}

// UniversalFunc Универсальная функция, которая работает непосредственно с записями.
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
			w.Write([]byte("Нет такой записи"))
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
			w.Write([]byte("Нет такой записи"))
			return nil, nil
		}
		shops = append(shops[:id], shops[id+1:]...)
		return nil, shops

	}

	w.Write([]byte("Ничего не нашлось. Сорян\t"))
	return nil, nil
}

//Вынес функции методов GET, потому что они не модифицируют записи никак.
//Ибо UniversalFunc становится огромным

// GetAll Функция, которая получает все записи той, или иной сущности. В зависимости от URL-a
func GetAll(w http.ResponseWriter, r *http.Request) ([]model.User, []model.Shop) {
	if r.URL.RequestURI() == "/getallusers" {
		return users, nil
	} else if r.URL.RequestURI() == "/getallshops" {
		return nil, shops

	}
	return nil, nil
}

// GetOneTable Функция, которая получает одну запись по имени той, или иной сущности. В зависимости от URL-a
func GetOneTable(w http.ResponseWriter, r *http.Request) (model.User, model.Shop) {
	if strip(r.URL.RequestURI()) == "/getoneuser/" {
		var reqId = mux.Vars(r)["title"]
		result := strings.ToLower(reqId)

		for i, user := range users {
			if strings.ToLower(user.Name) == result {
				return users[i], model.Shop{}
			}
		}

	} else if strip(r.URL.RequestURI()) == "/getoneshop/" {
		var reqId = mux.Vars(r)["title"]
		result := strings.ToLower(reqId)

		for i, shop := range shops {
			if strings.ToLower(shop.Title) == result {
				return model.User{}, shops[i]
			}
		}

	}

	return model.User{}, model.Shop{}

}

// GetOneField Возвращает одно поле по имени
func GetOneField(w http.ResponseWriter, r *http.Request) (model.User, model.Shop) {
	if strip(r.URL.RequestURI()) == "/getfielduser/" {
		var reqId = mux.Vars(r)["title"]
		result := strings.ToLower(reqId)
		mp := make(map[string]interface{})

		for _, user := range users {
			if strings.ToLower(user.Name) == result {
				mp["Name"] = user.Name
			} else if strings.ToLower(user.Otch) == result {
				mp["Otch"] = user.Otch
			} else if strings.ToLower(user.Family) == result {
				mp["Family"] = user.Family
			} else if strings.ToLower(user.Registration) == result {
				mp["Registration"] = user.Registration
			}
		}
		jsonbody, err := json.Marshal(mp)
		fmt.Println(jsonbody)
		if err != nil {
			fmt.Println(err)
		}
		us := model.User{}
		if err = json.Unmarshal(jsonbody, &us); err != nil {
			fmt.Println(err)
			return model.User{}, model.Shop{}
		}
		return us, model.Shop{}

	} else if strip(r.URL.RequestURI()) == "/getfieldshop/" {
		var reqId = mux.Vars(r)["title"]
		result := strings.ToLower(reqId)
		mp := make(map[string]interface{})

		for _, shop := range shops {
			if strings.ToLower(shop.Title) == result {
				mp["Title"] = shop.Title
			} else if strings.ToLower(shop.Working) == result {
				mp["Working"] = shop.Working
			} else if strings.ToLower(shop.Address) == result {
				mp["Address"] = shop.Address
			}
		}
		jsonbody, err := json.Marshal(mp)
		fmt.Println(jsonbody)
		if err != nil {
			fmt.Println(err)
		}
		sh := model.Shop{}
		if err = json.Unmarshal(jsonbody, &sh); err != nil {
			fmt.Println(err)
			return model.User{}, model.Shop{}
		}
		return model.User{}, sh
	}
	return model.User{}, model.Shop{}

}
