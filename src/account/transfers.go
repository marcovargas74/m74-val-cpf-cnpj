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

//SaveTransfer function to pre save transfer in system
func (t *TransferBank) SaveTransfer(w http.ResponseWriter, r *http.Request, tokeOrigin string) {

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.AccountOriginID = tokeOrigin
	fmt.Printf("\nSaveTransfer OriID:%s  DestID:%s value %.2f\n", t.AccountOriginID, t.AccountDestinationID, t.Amount)
	if errs := validator.Validate(t); errs != nil {
		log.Printf("INVALIDO %v\n", errs)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	t.ID = NewUUID()
	t.CreatedAt = time.Now()
	log.Printf("UUIDv4: %s\n", t.ID)

	if !IsValidUUID(t.ID) {
		log.Printf("invalid UUID:")
		w.WriteHeader(http.StatusBadRequest)
		return
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
		message := fmt.Sprintf("Can not save account from %v", t.ID)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := t.ExecTransation(w, r)
	if err != nil {
		return
	}

	json, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	w.WriteHeader(http.StatusOK)

}

//SaveTransferInDB function to save transfer in DB - Save only have balance in source account
func (t *TransferBank) SaveTransferInDB() bool {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Print(err)
		return false
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		return false
	}

	stmt, err := tx.Prepare("insert into transfers(id, ori, dest, amount, createAt) values(?,?,?,?,?)")
	if err != nil {
		log.Print(err)
		return false
	}

	_, err = stmt.Exec(t.ID, t.AccountOriginID, t.AccountDestinationID, t.Amount, t.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}
	log.Printf("    -->>SAVE ID Destino:[%s] <- ID ORIGIN:[%s]\n", t.AccountOriginID, t.AccountDestinationID)

	tx.Commit()
	return true
}

//GetTransfers Show All transfers
func (t *TransferBank) GetTransfers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, ori, dest, amount, createAt from transfers")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
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

//GetTransfersByID Show All transfers by ID
func (t *TransferBank) GetTransfersByID(w http.ResponseWriter, r *http.Request, UserID string) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, ori, dest, amount, createAt from transfers where ori = ?", UserID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}

	defer rows.Close()

	var transfers []TransferBank
	for rows.Next() {
		var aTransfer TransferBank
		rows.Scan(&aTransfer.ID, &aTransfer.AccountOriginID, &aTransfer.AccountDestinationID, &aTransfer.Amount, &aTransfer.CreatedAt)
		log.Printf(">>ID Destino:[%s] <- ID ORIGIN:[%s]\n", aTransfer.AccountOriginID, aTransfer.AccountDestinationID)
		transfers = append(transfers, aTransfer)
	}

	json, _ := json.Marshal(transfers)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

//GetTranferByID Get All transfers by ID - Return a Transfers List
func GetTranferByID(findID string) (TransferBank, error) {

	if !IsValidUUID(findID) {
		return TransferBank{}, fmt.Errorf("invalid ID: %s", findID)
	}

	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Print(err)
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

//ExecTransation FUNCAO PRINCIPAL USado para fazer a transação*/
func (t *TransferBank) ExecTransation(w http.ResponseWriter, r *http.Request) error {

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

}
