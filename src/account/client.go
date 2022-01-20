package account

import (
	"fmt"

	bankAPI "github.com/marcovargas74/m74-bank-api/src/api-bank"
)

//Client Struct a Client Type
type Client struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	bankAPI.Login
}

//SetID used to set ID at Client
func (c *Client) SetID(aID, name string) {
	c.Name = name
	c.ID = aID
}

//PrintCPF used to print a CPF by name
func (c Client) PrintCPF(name string) {
	fmt.Println(c.CPF)
}

//GetCPFbyName used to Get CPF by Name
func (c Client) GetCPFbyName(name string) string {
	return c.CPF
}

//GetCPFbyID used to Get CPF by ID
func (c Client) GetCPFbyID(ID string) string {
	return c.CPF
}

func useDataLogin() {
	client := Client{}
	client.CPF = "111.111.111-11"
	client.Secret = "111"
	fmt.Println(client.CPF)
}
