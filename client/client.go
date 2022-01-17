package client

import (
	"fmt"

	"github.com/marcovargas74/m74-bank-api/login"
)

type Client struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	login.LoginBank
}

func (c *Client) setId(aID, name string) {
	c.Name = "Maria"
	c.ID = aID
}

func (c Client) printCPF(name string) {
	fmt.Println(c.CPF)
}

func useDataLogin() {
	client := Client{}
	client.CPF = "111.111.111-11"
	client.Secret = "111"
	fmt.Println(client.CPF)
}
