package account

import (
	"encoding/json"
	"fmt"
)

func transfer() {
	fmt.Println("transfer")
}

//TransferBank is A struct to used to make a transfer
type TransferBank struct {
	ID                   string  `json:"id"`
	AccountOriginID      string  `json:"acount_origin_id"`
	AccountDestinationID string  `json:"Account_destination_id"`
	Amount               float64 `json:"Amount"`
	CreatedAt            string  `json:"created_at"` //TODO change to date
}

func structAndJSONTransfer() {
	transfer1 := TransferBank{"xyz", "abc", "def", 12.00, "17-01-2022"}
	transfJSON, _ := json.Marshal(transfer1)
	fmt.Println(string(transfJSON))
	//Convert Json To struct
	var aTransfFromJSON TransferBank
	json.Unmarshal(transfJSON, &aTransfFromJSON)
	fmt.Println(aTransfFromJSON.ID)

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
