package service

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zapirus/testwbapis/internal/model"
	"strconv"
	"strings"
	"sync"
)

var (
	users []model.User
	shops []model.Shop
)

// UniversalFunc Универсальная функция, которая работает непосредственно с записями.
func UniversalFunc(met, url, id string, newUser model.User, newShop model.Shop) ([]model.User, []model.Shop) {
	var mx sync.Mutex

	if url == "/user" && met == "POST" {
		mx.Lock()
		users = append(users, newUser)
		mx.Unlock()
		return users, nil
	} else if url == "/shop" && met == "POST" {
		mx.Lock()
		shops = append(shops, newShop)
		mx.Unlock()
		return nil, shops
	} else if url == "/changeuser/" && met == "PUT" {
		ids, err := strconv.Atoi(id)
		if err != nil {
			logrus.Printf("Не удалось сконвертировать. %s", err)
			return nil, nil
		}
		if ids >= len(users) {
			logrus.Printf("Вы выбрали элемент, которого нет. %s", err)
			return nil, nil
		}
		users[ids] = newUser
		return users, nil

	} else if url == "/changeshop/" && met == "PUT" {
		ids, err := strconv.Atoi(id)
		if err != nil {
			logrus.Printf("Не удалось сконвертировать. %s", err)
			return nil, nil
		}
		if ids >= len(shops) {
			logrus.Printf("Вы выбрали элемент, которого нет. %s", err)
			return nil, nil
		}
		shops[ids] = newShop
		return nil, shops
	} else if url == "/changeuser/" && met == "DELETE" {
		ids, err := strconv.Atoi(id)
		if err != nil {
			logrus.Printf("Не удалось сконвертировать. %s", err)
			return nil, nil
		}
		if ids >= len(users) {
			logrus.Printf("Вы выбрали элемент, которого нет. %s", err)
			return nil, nil
		}
		users = append(users[:ids], users[ids+1:]...)
		return users, nil
	} else if url == "/changeshop/" && met == "DELETE" {
		ids, err := strconv.Atoi(id)
		if err != nil {
			logrus.Printf("Не удалось сконвертировать. %s", err)
			return nil, nil
		}
		if ids >= len(shops) {
			logrus.Printf("Вы выбрали элемент, которого нет. %s", err)
			return nil, nil
		}
		shops = append(shops[:ids], shops[ids+1:]...)
		return nil, shops
	} else if url == "/getallusers" && met == "GET" {
		return users, nil
	} else if url == "/getallshops" && met == "GET" {
		return nil, shops

	}
	logrus.Println("Извините, ничего не нашлось")
	return nil, nil
}

// GetOneEntity Функция, которая получает одну запись по имени той, или иной сущности. В зависимости от URL-a
func GetOneEntity(url, name string) (model.User, model.Shop) {
	if url == "/getoneuser/" {
		result := strings.ToLower(name)
		for i, user := range users {
			if strings.ToLower(user.Name) == result {
				return users[i], model.Shop{}
			}
		}

	} else if (url) == "/getoneshop/" {
		result := strings.ToLower(name)

		for i, shop := range shops {
			if strings.ToLower(shop.Title) == result {
				return model.User{}, shops[i]
			}
		}
	}
	return model.User{}, model.Shop{}

}

// GetOneField Возвращает одно поле по имени
func GetOneField(urlField, reqId string) (model.User, model.Shop) {
	if urlField == "/getfielduser/" {

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
		if err != nil {
			logrus.Fatalln(err)
		}
		us := model.User{}
		if err = json.Unmarshal(jsonbody, &us); err != nil {
			logrus.Printf("Не удалось преаброзовать: %s", err)
			return model.User{}, model.Shop{}
		}
		return us, model.Shop{}

	} else if urlField == "/getfieldshop/" {
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
		if err != nil {
			logrus.Printf("Не удалось преаброзовать: %s", err)
		}
		sh := model.Shop{}
		if err = json.Unmarshal(jsonbody, &sh); err != nil {
			logrus.Printf("Не удалось преаброзовать: %s", err)
			return model.User{}, model.Shop{}
		}
		return model.User{}, sh
	}
	return model.User{}, model.Shop{}

}
