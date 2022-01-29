package main

import (
	"fmt"

	bank "github.com/marcovargas74/m74-bank-api/src/api-bank"
)

var isProduction = false

func init() {
	bank.SetIsProduction(isProduction)
	bank.CreateDB()
}

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bank.GetVersion(), bank.GetIsProduction())

	//account.StructAndJSON()
	bank.StartAPI("dev")

}
