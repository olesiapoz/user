package main

import (
	"addresses"
)

// Cust forwading to Hello
func Addresses(args map[string]interface{}) map[string]interface{} {

	return addresses.Addresses(args)

	//return login.Customer(args["__ow_headers"].(map[string]interface{}))

}
