package main

import (
	"customer"
)

// Cust forwading to Hello
func Customer(args map[string]interface{}) map[string]interface{} {

	return customer.Customer(args)

	//return login.Customer(args["__ow_headers"].(map[string]interface{}))

}
