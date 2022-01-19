package account

import (
	"testing"
)

/*const (
	erroMsg = "Valor esperado %v, resultado encontrado %v"
)*/

func TestSetTransfers(t *testing.T) {

	checkResult := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf(erroMsg, resultado, esperado)
		}
	}

	t.Run("test function with true", func(t *testing.T) {
		valorEsperado := "xyz"

		transfer1 := TransferBank{"xyz", "abc", "def", 12.00, "17-01-2022"}
		valorRetornado := transfer1.ID
		//transfJson, _ := json.Marshal(transfer1)
		//fmt.Println(string(transfJson))
		//Convert Json To struct
		//var aTransfFromJson TransferBank
		//json.Unmarshal(transfJson, &aTransfFromJson)
		//fmt.Println(aTransfFromJson.ID)

		checkResult(t, valorRetornado, valorEsperado)
	})

}
