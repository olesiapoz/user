package muser

import (
  "aaa"
// "github.com/subosito/gotenv"
	"crypto/sha1"
//	"reflect"
	corelog "log"
	"fmt"
	"io"
//	"encoding/json"
//	"os"
  //    "log"
	"github.com/microservices-demo/user/api"
	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/users"
	"github.com/microservices-demo/user/db/mongodb"


//	"github.com/olesiapoz/Serverless-microservices/user/db/mongodb"

)

type userResponse struct {
        User users.User `json:"user"`
}

func init() {
aaa.Load()
//fmt.Println(os.Getenv("USER_DATABASE"))
}

func init() {
	db.Register("mongodb", &mongodb.Mongo{})
//corelog.Print("DbInit")
//	session := *mongodb.Mongo{}
//	fmt.Println("Created DB")
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
//        os.Setenv("MONGO_HOST","192.168.0.3:27017")
//        os.Setenv("USER_DATABASE", "mongodb")
	fmt.Println(req) 
 
	password := req["password"].(string)
	//password := string(passw[0:len(passw)-1])
	username := req["user"].(string)
    
    dbconn := false
    //    fmt.Println("Done")
	for !dbconn {
		err := db.Init()
//		fmt.Println(db.DBTypes)
		if err != nil {
			if err == db.ErrNoDatabaseSelected {
				corelog.Fatal(err)
			}
//			corelog.Print("DB : " + err.Error())
		} else {
			dbconn = true
  //                      fmt.Println("Connected to DB")
		}
	}
	
        // Service domain.
	var service api.Service
	{
		service = api.NewFixedService()
	}
//      corelog.Print("service" + service.Login(username, password).user.id)

	u, err := db.GetUserByName(username)
	if err != nil {
        corelog.Print("E1: " + err.Error())
        fmt.Println("E2 : " + err.Error())
        }

	
	salt := u.Salt

        fmt.Println("Salt: " + salt)
        fmt.Println("PSW: " + u.Password)
	fmt.Println("user: " + username + " pwd: " + password)
	fmt.Println(calculatePassHash(password, salt))

	var user users.User
	//corelog.Print("User" + user)	
	user, err1 := service.Login(username, password)
	if err1 != nil {
	corelog.Print("R1: " + err1.Error())
	fmt.Println("R2 : " + err1.Error())
	}


	
//	json_st, _ := json.Marshal(userResponse{User: user})
//	mp =: map[string]string(userResponse{User: user})
//      json_st, _ := json.Marshal(mp)

  //      fmt.Println(userResponse{User: user})
     mp := make(map[string]interface{})
     mp ["user"]= user
  // log in stdout or in stderr
//     log.Printf("name=%s\n", username)
//	fmt.Println(string(json_st)) 
//	fmt.Println(reflect.TypeOf(json))
//	fmt.Println(userResponse{User: user})	
// json_st, _ := json.Marshal(mp)

//	fmt.Println("%s\n", json_st)
        return mp
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
	


