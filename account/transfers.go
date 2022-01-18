package account

import (
	"encoding/json"
	"fmt"
)

func transfer() {
	fmt.Println("transfer")
}

type TransferBank struct {
	ID                     string  `json:"id"`
	Account_origin_id      string  `json:"acount_origin_id"`
	Account_destination_id string  `json:"Account_destination_id"`
	Amount                 float64 `json:"Amount"`
	Created_at             string  `json:"created_at"` //TODO change to date
}

func structAndJsonTransfer() {
	transfer1 := TransferBank{"xyz", "abc", "def", 12.00, "17-01-2022"}
	transfJson, _ := json.Marshal(transfer1)
	fmt.Println(string(transfJson))
	//Convert Json To struct
	var aTransfFromJson TransferBank
	json.Unmarshal(transfJson, &aTransfFromJson)
	fmt.Println(aTransfFromJson.ID)

}

/*
/transfers
A entidade Transfer possui os seguintes atributos:

id
account_origin_id
account_destination_id
amount
created_at
Espera-se as seguintes ações:

GET /transfers - obtém a lista de transferencias da usuaria autenticada.
POST /transfers - faz transferencia de uma Account para outra.
Regras para esta rota

Quem fizer a transferência precisa estar autenticada.
O account_origin_id deve ser obtido no Token enviado.
Caso Account de origem não tenha saldo, retornar um código de erro apropriado
Atualizar o balance das contas
*/
