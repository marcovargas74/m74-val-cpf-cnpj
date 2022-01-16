package main

import (
	"fmt"

	bankAPI "github.com/marcovargas74/m74-bank-api"
)

var IsProduction = false

func init() {
	bankAPI.SetIsProduction(IsProduction)
}

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bankAPI.GetVersion(), bankAPI.GetIsProduction())

	bankAPI.StartAPI("dev")
	/*tratador := http.HandlerFunc(bankAPI.ServidorJogador)
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}*/

}
