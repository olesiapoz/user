package addresses

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
func Addresses(req map[string]interface{}) map[string]interface{} {

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

		adds, err1 := service.GetAddresses(custID)

		if err1 != nil {
			corelog.Print("Customer Post Service Error: " + err1.Error())
		}

		if custID == "" {
			mp["body"] = EmbedStruct{addressesResponse{Addresses: adds}}
			return mp

		}

		if len(adds) == 0 {
			mp["body"] = users.Address{}
			return mp
		}
		mp["body"] = adds[0]
		return mp

	}

	if method == "post" {
		userID := req["userID"].(string)
		address := req["Address"].(users.Address)

		id, e := service.PostAddress(address, userID)
		if e != nil {
			corelog.Print("Address Post Service Error: " + e.Error())
		}
		mp["body"] = postResponse{ID: id}

		return mp

	}
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

type addressPostRequest struct {
	users.Address
	UserID string `json:"userID"`
}

type EmbedStruct struct {
	Embed interface{} `json:"_embedded"`
}
