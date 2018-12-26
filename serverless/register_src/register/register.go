package register

import (
	"aaa"
	corelog "log"

	"github.com/microservices-demo/user/api"
	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/db/mongodb"
)

func init() {
	aaa.Load()
}

func init() {
	db.Register("mongodb", &mongodb.Mongo{})
}

//register url logic
func Register(req map[string]interface{}) map[string]interface{} {

	email := req["email"].(string)
	firstName := req["firstName"].(string)
	lastName := req["lastName"].(string)
	password := req["password"].(string)
	username := req["username"].(string)

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

	// Service domain.
	var service api.Service
	{
		service = api.NewFixedService()
	}

	userID, err1 := service.Register(username, password, email, firstName, lastName)
	if err1 != nil {
		corelog.Print("Register Service Error: " + err1.Error())
	}

	mp := make(map[string]interface{})
	mp["id"] = userID

	return mp
}
