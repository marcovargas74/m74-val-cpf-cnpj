package main

import (
	"fmt"

	account "github.com/marcovargas74/m74-bank-api/src/account"
	bank "github.com/marcovargas74/m74-bank-api/src/api-bank"
)

var isProduction = false

func init() {
	bank.SetIsProduction(isProduction)
}

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bank.GetVersion(), bank.GetIsProduction())

	account.StructAndJSON()
	bank.StartAPI("dev")

	/*tratador := http.HandlerFunc(bankAPI.ServidorJogador)
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}*/

}
