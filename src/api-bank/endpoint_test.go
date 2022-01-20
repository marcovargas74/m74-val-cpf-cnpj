package m74bankapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

func newReqGetPoints(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", name), nil)
	return request
}

func newReqEndpoints(urlPrefix, urlName string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", urlPrefix, urlName), nil)

	fmt.Printf("endpoint: %v\n", request.URL)
	return request
}

/*
func checkResult(t *testing.T, result, wait string) {
	t.Helper()
	if result != wait {
		t.Errorf(erroMsg, result, wait)
	}
}

*/
func TestServidorJogador(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Retornar resultado de Nobody",
			wantValue: "",
			inData:    "Nobody",
		},
		{
			give:      "Retornar resultado de Maria",
			wantValue: "20",
			inData:    "Maria",
		},

		{
			give:      "Retornar resultado de Pedro",
			wantValue: "10",
			inData:    "Pedro",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpoints("/jogadores", tt.inData)
			resposta := httptest.NewRecorder()

			ServidorJogador(resposta, requisicao)

			recebido := resposta.Body.String()
			//esperado := "10"

			//checkResult(t, recebido, esperado)
			assert.Equal(t, recebido, tt.wantValue)
			//fmt.Printf("valor: %v\n", valorRetornado[0:4])
		})

	}

}

/*
	//*TODO endpoint usado no banc
	func callbackAccount(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("callbackAccount data in %v\n", r.URL)
		fmt.Fprint(w, message)
	}

	func callbackLogin(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("callbackLogin data in %v\n", r.URL)
		fmt.Fprint(w, message)
	}

	func callbackTransfer(w http.ResponseWriter, r *http.Request) {*/

func TestCallbackAccount(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: "callbackAccount data in /accounts/Nobody\n",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: "callbackAccount data in /accounts/\n",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: "callbackAccount data in /accounts/123\n",
			inData:    "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpoints("/accounts", tt.inData)
			resposta := httptest.NewRecorder()

			callbackAccount(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}

func TestCallbackLogin(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint login com NOBODY",
			wantValue: "callbackLogin data in /login/Nobody\n",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Login com caracter vazio",
			wantValue: "callbackLogin data in /login/\n",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Login com ID 123",
			wantValue: "callbackLogin data in /login/123\n",
			inData:    "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpoints("/login", tt.inData)
			resposta := httptest.NewRecorder()

			callbackLogin(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}

func TestCallbackTransfer(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint Transfer com NOBODY",
			wantValue: "callbackTransfer data in /transfers/Nobody\n",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Transfer com caracter vazio",
			wantValue: "callbackTransfer data in /transfers/\n",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Transfer com ID 123",
			wantValue: "callbackTransfer data in /transfers/123\n",
			inData:    "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpoints("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			callbackTransfer(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}
