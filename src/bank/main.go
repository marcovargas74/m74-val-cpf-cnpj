package main

import (
	"fmt"

	"github.com/marcovargas74/m74-bank-api/src/account"
	bank "github.com/marcovargas74/m74-bank-api/src/api-bank"
)

var isProduction = false

func init() {
	bank.SetIsProduction(isProduction)
	account.CreateDB()
}

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bank.GetVersion(), bank.GetIsProduction())

	//account.StructAndJSON()
	bank.StartAPI("dev")

}
