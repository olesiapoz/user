package login

import (
	"aaa"
	"encoding/base64"
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

//login url logic
func Login(req map[string]interface{}) map[string]interface{} {

	authorization64 := req["authorization"].(string)

	username, password, e := parseBasicAuth(authorization64)
	if e != true {
		corelog.Print("ParseBasicAuth Error")
	}

	//db connection
	//password := "tpassword"
	//username := "tuser"

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

func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
