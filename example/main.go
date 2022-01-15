package main

import (
	"fmt"

	bankAPI "github.com/marcovargas74/m74-bank-api"
)

func init() {
	bankAPI.SetIsProduction(true)
}

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bankAPI.GetVersion(), bankAPI.GetIsProduction())
}
