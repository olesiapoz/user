package muser

import (
	"aaa"
	//	"crypto/sha1"
	corelog "log"
	//	"encoding/json"
	//	"os"
	//    "log"
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

/*
func main(){
        var input = make(map[string]interface{})
	input["user"] = "test"
        input["password"] = 12345
        output := Main(input)
	json, _ := json.Marshal(output)
	fmt.Printf("%s", json)
}
*/

func Muser(req map[string]interface{}) map[string]interface{} {

	password := req["password"].(string)
	username := req["user"].(string)

	//db connection
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

	var user users.User

	user, err1 := service.Login(username, password)
	if err1 != nil {
		corelog.Print("Login Service Error: " + err1.Error())
	}

	mp := make(map[string]interface{})
	mp["user"] = user

	return mp
}
