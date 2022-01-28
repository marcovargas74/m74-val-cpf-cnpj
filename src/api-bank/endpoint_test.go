package m74bankapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

func newReqEndpointsGET(urlPrefix, urlName string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", urlPrefix, urlName), nil)

	fmt.Printf("endpoint: %v\n", request.URL)
	return request
}

func newReqEndpointsPOST(urlPrefix, urlName string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", urlPrefix, urlName), nil)

	fmt.Printf("endpoint: %v\n", request.URL)
	return request
}
func TestServerBank(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Retornar resultado de Nobody",
			wantValue: "Endpoint not found",
			inData:    "Nobody",
		},
	}

	servidor := NewServerBank()
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao, _ := http.NewRequest(http.MethodGet, "/", nil)
			resposta := httptest.NewRecorder()

			servidor.ServeHTTP(resposta, requisicao)

			recebido := resposta.Body.String()
			//fmt.Printf("\ndado recebido %s\n", recebido)

			assert.Equal(t, recebido, tt.wantValue)
			assert.Equal(t, resposta.Code, http.StatusOK)

		})

	}

}

func TestCallbackAccountGET(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: "GET /accounts/Nobody",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: "GET /accounts/",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: "GET /accounts/123",
			inData:    "123",
		},
	}

	server := new(ServerBank)

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsGET("/accounts", tt.inData)
			resposta := httptest.NewRecorder()

			server.CallbackAccounts(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}

func TestCallbackAccountPost(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: "POST /accounts/Nobody",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: "POST /accounts/",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: "POST /accounts/123",
			inData:    "123",
		},
	}

	server := new(ServerBank)

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/accounts", tt.inData)
			resposta := httptest.NewRecorder()

			server.CallbackAccounts(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}
}

func TestCallbackLoginGET(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: "/login/Nobody",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: "/login/",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: "/login/123",
			inData:    "123",
		},
	}

	server := new(ServerBank)

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsGET("/login", tt.inData)
			resposta := httptest.NewRecorder()

			server.CallbackLogin(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}

func TestCallbackLoginPOST(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: "POST /login/Nobody",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: "POST /login/",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: "POST /login/123",
			inData:    "123",
		},
	}

	server := new(ServerBank)

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/login", tt.inData)
			resposta := httptest.NewRecorder()

			server.CallbackLogin(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}

func TestCallbackTransferGET(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint TRANFER com NOBODY",
			wantValue: "GET /transfers/Nobody",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint TRANFER com caracter vazio",
			wantValue: "GET /transfers/",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint TRANFER com ID 123",
			wantValue: "GET /transfers/123",
			inData:    "123",
		},
	}

	server := new(ServerBank)

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsGET("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			server.CallbackTransfer(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}

func TestCallbackTransferPOST(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: "POST /transfers/Nobody",
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: "POST /transfers/",
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: "POST /transfers/123",
			inData:    "123",
		},
	}

	server := new(ServerBank)

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			server.CallbackLogin(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}
