package m74bankapi

import (
	"encoding/json"
	"fmt"
)

// LoginBank
type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

//StructAndJson teste do struct
func StructAndJson() {
	myLogin := Login{CPF: "111.111.111-11", Secret: "111"}
	loginJson, _ := json.Marshal(myLogin)
	fmt.Println(string(loginJson))
	//Convert Json To struct
	var myNewLogin Login
	json.Unmarshal(loginJson, &myNewLogin)
	fmt.Println(myNewLogin.Secret)

}

/*
func Secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
	}
	return ""
}*/
/*
func (l *Login) Prepare() {
	authenticator := auth.NewBasicAuthenticator("example.com", Secret)

	//a := auth.NewBasicAuthenticator("example.com", Secret)
	if username := authenticator.CheckAuth(this.Ctx.Request); username == "" {
		authenticator.RequireAuth(this.Ctx.ResponseWriter, this.Ctx.Request)
	}
}*/

/*
type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
	a := auth.NewBasicAuthenticator("example.com", Secret)
	if username := a.CheckAuth(this.Ctx.Request); username == "" {
		a.RequireAuth(this.Ctx.ResponseWriter, this.Ctx.Request)
	}
}

func (this *MainController) Get() {
	this.Data["Username"] = "astaxie"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
*/
/*
/login
A entidade Login possui os seguintes atributos:

cpf
secret
Espera-se as seguintes ações:

POST /login - autentica a usuaria
Regras para esta rota

Deve retornar token para ser usado nas rotas autenticadas
*/

/*

		http.StatusNetworkAuthenticationRequired

package controllers

import (
    "github.com/abbot/go-http-auth"
    "github.com/astaxie/beego"
)

func Secret(user, realm string) string {
    if user == "john" {
        // password is "hello"
        return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
    }
    return ""
}

type MainController struct {
    beego.Controller
}

func (this *MainController) Prepare() {
    a := auth.NewBasicAuthenticator("example.com", Secret)
    if username := a.CheckAuth(this.Ctx.Request); username == "" {
        a.RequireAuth(this.Ctx.ResponseWriter, this.Ctx.Request)
    }
}

func (this *MainController) Get() {
    this.Data["Username"] = "astaxie"
    this.Data["Email"] = "astaxie@gmail.com"
    this.TplNames = "index.tpl"
}



*/
