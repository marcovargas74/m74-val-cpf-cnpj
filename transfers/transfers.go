package transfers

import "fmt"

func transfer() {
	fmt.Println("transfer")
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
