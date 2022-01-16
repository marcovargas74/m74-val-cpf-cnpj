package main

import (
	"fmt"
	"log"
	"net/http"

	bankAPI "github.com/marcovargas74/m74-bank-api"
)

var IsProduction = false

func init() {
	//IsProduction = true
	bankAPI.SetIsProduction(IsProduction)
}

/*
func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<valor><font color='#2e802e' size='4'>INFO Teste de %s</font></valor>", bankAPI.GetVersion())
}*/

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bankAPI.GetVersion(), bankAPI.GetIsProduction())

	tratador := http.HandlerFunc(bankAPI.ServidorJogador)
	if err := http.ListenAndServe(":5000", tratador); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
	//if IsProduction {
	//bankAPI.StartAPI("")
	//go appl.StartAppliance(appl.GetMode())
	//}
}
