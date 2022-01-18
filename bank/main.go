package main

import (
	"fmt"

	account "github.com/marcovargas74/m74-bank-api/account"
	bankAPI "github.com/marcovargas74/m74-bank-api/api-bank"
)

var IsProduction = false

func init() {
	bankAPI.SetIsProduction(IsProduction)
}

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bankAPI.GetVersion(), bankAPI.GetIsProduction())

	account.StructAndJson()
	bankAPI.StartAPI("dev")

	/*tratador := http.HandlerFunc(bankAPI.ServidorJogador)
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}*/

}
