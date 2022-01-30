package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/validator.v2"
)

/*
func transfer() {
	fmt.Println("transfer")
}*/

/*
type TransferBank struct {
	ID                   string    `json:"id" validate:"required,uuid4"`
	AccountOriginID      string    `json:"acount_origin_id" validate:"required,uuid4"`
	AccountDestinationID string    `json:"Account_destination_id" validate:"required,uuid4"`
	Amount               float64   `json:"Amount" validate:"gt=0,required"`
	CreatedAt            time.Time `json:"created_at"`
}
*/

//TransferBank is A struct to used to make a transfer
type TransferBank struct {
	ID                   string    `json:"id"`
	AccountOriginID      string    `json:"account_origin_id" `
	AccountDestinationID string    `json:"account_destination_id" `
	Amount               float64   `json:"amount" `
	CreatedAt            time.Time `json:"created_at"`
}

/*
func structAndJSONTransfer() {
	transfer1 := TransferBank{"xyz", "abc", "def", 12.00, "17-01-2022"}
	transfJSON, _ := json.Marshal(transfer1)
	fmt.Println(string(transfJSON))
	//Convert Json To struct
	var aTransfFromJSON TransferBank
	json.Unmarshal(transfJSON, &aTransfFromJSON)
	fmt.Println(aTransfFromJSON.ID)
}*/

func (t *TransferBank) SaveTransfer(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("\nSaveTransfer OriID:%s  DestID:%s value %.2f\n", t.AccountOriginID, t.AccountDestinationID, t.Amount)
	if errs := validator.Validate(t); errs != nil {
		fmt.Printf("INVALIDO %v\n", errs) // do something
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	t.ID = NewUUID()
	t.CreatedAt = time.Now()
	fmt.Printf("UUIDv4: %s\n", t.ID)

	if !IsValidUUID(t.ID) {
		fmt.Printf("Something gone wrong:")
	}

	if !IsValidUUID(t.AccountOriginID) || !IsValidUUID(t.AccountDestinationID) {
		fmt.Printf("Something gone wrong:")
		fmt.Printf("INVALID ID \n") // do something
		message := fmt.Sprintf("invalid ID %s or %s", t.AccountOriginID, t.AccountDestinationID)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if !t.SaveTransferInDB() {
		message := fmt.Sprintf("Can´t save account from %v", t.ID)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := t.ExecTransation(w, r)
	if err != nil {
		//message := fmt.Sprintf("Fail to Exec Transation %s", t.ID)
		//fmt.Fprint(w, message)
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	w.WriteHeader(http.StatusOK)

}

/* TODO so executar se tiver saldo na conta Origim
 */
func (t *TransferBank) SaveTransferInDB() bool {
	db, err := sql.Open("mysql", DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into transfers(id, ori, dest, amount, createAt) values(?,?,?,?,?)")

	_, err = stmt.Exec(t.ID, t.AccountOriginID, t.AccountDestinationID, t.Amount, t.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return false
	}
	fmt.Printf("    -->>SAVE ID Destino:[%s] <- ID ORIGIN:[%s]\n", t.AccountOriginID, t.AccountDestinationID)

	tx.Commit()
	return true
}

/*
func (t *TransferBank) GetAccountByID(w http.ResponseWriter, r *http.Request, ID string) {

	fmt.Printf("   -->GetAccountByID [%s] \n", ID)
	if !IsValidUUID(ID) {
		fmt.Printf("Something gone wrong:")
		fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message := StructAndJSON()
	a.ShowAccountByID(w, r, ID)

	//fmt.Fprint(w, message)
	w.WriteHeader(http.StatusOK)
}
*/

//ShowAccountAll mostra todos as contas
func (t *TransferBank) GetTransfers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, ori, dest, amount, createAt from transfers")

	defer rows.Close()

	var transfers []TransferBank
	for rows.Next() {
		var aTransfer TransferBank
		rows.Scan(&aTransfer.ID, &aTransfer.AccountOriginID, &aTransfer.AccountDestinationID, &aTransfer.Amount, &aTransfer.CreatedAt)
		fmt.Printf(">>ID Destino:[%s] <- ID ORIGIN:[%s]\n", aTransfer.AccountOriginID, aTransfer.AccountDestinationID)
		transfers = append(transfers, aTransfer)
	}

	json, _ := json.Marshal(transfers)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func GetTranferByID(findID string) (TransferBank, error) {

	if !IsValidUUID(findID) {
		return TransferBank{}, fmt.Errorf("invalid ID: %s", findID)
	}

	db, err := sql.Open("mysql", DBSource)
	if err != nil {
		log.Fatal(err)
		return TransferBank{}, err
	}
	defer db.Close()

	var aTransfer TransferBank
	db.QueryRow("select id, ori, dest, amount from transfers where id = ?", findID).Scan(&aTransfer.ID, &aTransfer.AccountOriginID, &aTransfer.AccountDestinationID, &aTransfer.Amount)
	fmt.Printf("DADOS DO BANOC id[%s] data[%s %s]\n", findID, aTransfer.AccountOriginID, aTransfer.AccountDestinationID)

	if !IsValidUUID(aTransfer.ID) {
		return aTransfer, fmt.Errorf("transfer not found ID[%s]", findID)
	}

	return aTransfer, nil
}

/*
func (a *Account) GetTansferAccountByID( findID string) {
	//db, err := sql.Open("mysql", "root:Mysql#2510@/bankAPI")
	db, err := sql.Open("mysql", DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var account Account
	db.QueryRow("select id, nome, cpf, balance, secret, createAt from accounts where id = ?", findID).Scan(&account.ID, &account.Name, &account.CPF, &account.Balance, &account.Secret, &account.CreatedAt)
	account.Secret = "*****"
	json, _ := json.Marshal(account)

	w.Header().Set("Content-Type", "application/json")

	fmt.Printf("DADOS DO BANOC id[%s] data[%s]\n", findID, string(json))

	fmt.Fprint(w, string(json))
}
*/

// MakeTransfer withdraw from source account and credit to destination account
func (t *TransferBank) MakeTransfer(source, destination *Account) (bool, error) {

	fmt.Printf(" MakeTransfer check -->>Atual Values Destino:[%.2f] <- ORIGIN:[%.2f]\n", destination.Balance, source.Balance)
	if source.Balance < t.Amount {
		return false, fmt.Errorf("Account to be debited does not have sufficient balance[%.2f]", source.Balance)
	}

	source.Balance = source.Balance - t.Amount
	destination.Balance = destination.Balance + t.Amount
	fmt.Printf(" transferencia SUCCESS   -->>New Values Destino:[%.2f] <- ORIGIN:[%.2f]\n", destination.Balance, source.Balance)
	return true, nil
}

/*FUNCAO PRINCIPAL USado para fazer a transação*/
func (t *TransferBank) ExecTransation(w http.ResponseWriter, r *http.Request) error {

	/*aTransfer, err := GetTranferByID(t.ID)
	if err != nil {
		//fmt.Println(err)
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}*/

	accountToDeb, err := GetAccountByID(t.AccountOriginID)
	if err != nil {
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	accountToCred, err := GetAccountByID(t.AccountDestinationID)
	if err != nil {
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	_, err = t.MakeTransfer(&accountToDeb, &accountToCred)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return err
	}

	//TODO Refatorar criando um função para a conta atualizar seus dados
	if !UpdateBalanceByID(accountToDeb.ID, accountToDeb.Balance) {
		fmt.Printf("fail to update Debit account %s\n", err)
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return err
	}

	if !UpdateBalanceByID(accountToCred.ID, accountToCred.Balance) {
		fmt.Printf("fail to update Debit account %s\n", err)
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return err
	}

	return nil
	/*json, _ := json.Marshal(aTransfer)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	w.WriteHeader(http.StatusOK)*/

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
