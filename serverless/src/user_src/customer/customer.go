package customer

import (
	"aaa"
	corelog "log"
	"strings"

	"github.com/microservices-demo/user/api"
	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/db/mongodb"
	"github.com/microservices-demo/user/users"
)

type userResponse struct {
	User users.User `json:"user"`
}

func init() {
	aaa.Load()
}

func init() {
	db.Register("mongodb", &mongodb.Mongo{})
}

//customers url logic
func Customer(req map[string]interface{}) map[string]interface{} {

	method := req["__ow_method"].(string)

	dbconn := false

	for !dbconn {
		err := db.Init()
		if err != nil {
			if err == db.ErrNoDatabaseSelected {
				corelog.Fatal(err)
			}
		} else {
			dbconn = true
		}
	}

	mp := make(map[string]interface{})

	// Service domain.
	var service api.Service
	{
		service = api.NewFixedService()
	}

	attr := ""
	custID := ""

	if method == "get" {
		args := req["__ow_path"].(string)
		urlPart := strings.Split(args, "/")
		corelog.Print("args and url path" + args)
		corelog.Print("Len : " + string(len(urlPart)))

		if len(urlPart) > 2 {
			custID = urlPart[2]
			corelog.Print("CustID : " + custID)
		} else {
			custID = ""
		}

		if len(urlPart) > 3 {
			attr = urlPart[3]
			corelog.Print("urlid" + attr)
		} else {
			attr = ""
		}

		usrs, err1 := service.GetUsers(custID)
		corelog.Print("usrs")
		//	corelog.Print(usrs)

		if err1 != nil {
			corelog.Print("Customer Post Service Error: " + err1.Error())
		}

		if custID == "" {
			mp["body"] = EmbedStruct{usersResponse{Users: usrs}}
			corelog.Print("Empty Custid")
			//		corelog.Print(mp)
			return mp

		}

		if len(usrs) == 0 {
			if attr == "addresses" {
				mp["body"] = EmbedStruct{addressesResponse{Addresses: make([]users.Address, 0)}}
				corelog.Print("Addresses")
				corelog.Print(mp)
				return mp

			}
			if attr == "cards" {
				mp["body"] = EmbedStruct{cardsResponse{Cards: make([]users.Card, 0)}}
				//			corelog.Print(mp)
				return mp
			}
			mp["body"] = EmbedStruct{users.User{}}
			corelog.Print("user")
			//		corelog.Print(mp)
			return mp
		}

		user := usrs[0]

		db.GetUserAttributes(&user)

		if attr == "addresses" {
			mp["body"] = EmbedStruct{addressesResponse{Addresses: user.Addresses}}
			corelog.Print("attr Addresses")
			//		corelog.Print(mp)
			return mp
		}

		if attr == "cards" {
			mp["body"] = EmbedStruct{cardsResponse{Cards: user.Cards}}
			//		corelog.Print(mp)
			return mp
		}

		mp["body"] = user
		corelog.Print("user   !")
		//	corelog.Print(mp)
		return mp

	}

	if method == "post" {
		args := req["user"].(users.User)

		id, e := service.PostUser(args)
		if e != nil {
			corelog.Print("Customer Post Service Error: " + e.Error())
		}
		mp["body"] = postResponse{ID: id}
		corelog.Print("ID")
		//	corelog.Print(mp)
		return mp

	}
	corelog.Print(mp)
	corelog.Print("NO slashas")
	return mp

}

type usersResponse struct {
	Users []users.User `json:"customer"`
}

type addressesResponse struct {
	Addresses []users.Address `json:"address"`
}

type cardsResponse struct {
	Cards []users.Card `json:"card"`
}

type postResponse struct {
	ID string `json:"id"`
}

type EmbedStruct struct {
	Embed interface{} `json:"_embedded"`
}
