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

const tipoDoConteudoJSON = "application/json"

func verificaTipoDoConteudo(t *testing.T, resposta *httptest.ResponseRecorder, esperado string) {
	t.Helper()
	if resposta.Result().Header.Get("content-type") != esperado {
		t.Errorf("resposta não obteve content-type de %s, obtido %v", esperado, resposta.Result().Header)
	}
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

	servidor := NewServerBank("dev")
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
		wantValue int
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: 404,
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: 404,
			inData:    "123",
		},
	}

	server := NewServerBank("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsGET("/accounts", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			assert.Equal(t, resposta.Code, tt.wantValue)
		})

	}

}

func TestCallbackAccountPost(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: 404,
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: 404,
			inData:    "123",
		},
	}

	server := NewServerBank("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/accounts", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			assert.Equal(t, resposta.Code, tt.wantValue)
		})

	}
}

func TestCallbackLoginGET(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "Testa o Endpoint LOgin com NOBODY",
			wantValue: 404,
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint LOgin com caracter vazio",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Login com ID 123",
			wantValue: 404,
			inData:    "123",
		},
	}

	server := NewServerBank("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsGET("/login", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			assert.Equal(t, resposta.Code, tt.wantValue)
		})

	}

}

func TestCallbackLoginPOST(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: 404,
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: 401,
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: 404,
			inData:    "123",
		},
	}

	server := NewServerBank("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/login", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			//verificaTipoDoConteudo(t, resposta, tipoDoConteudoJSON)

			//recebido := resposta.Body.String()
			//assert.Equal(t, recebido, tt.wantValue)
			assert.Equal(t, resposta.Code, tt.wantValue)
		})

	}

}

func TestCallbackTransferGET(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "Testa o Endpoint TRANFER com NOBODY",
			wantValue: 500,
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint TRANFER com caracter vazio",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "Testa o Endpoint TRANFER com ID 123",
			wantValue: 500,
			inData:    "123",
		},
	}

	server := NewServerBank("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsGET("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			//verificaTipoDoConteudo(t, resposta, tipoDoConteudoJSON)

			//recebido := resposta.Body.String()
			assert.Equal(t, resposta.Code, tt.wantValue)
			//assert.Equal(t, recebido, tt.wantValue)
		})

	}

}

/*
func TestCallbackTransferPOST(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "Testa o Endpoint Account com NOBODY",
			wantValue: 500,
			inData:    "Nobody",
		},
		{
			give:      "Testa o Endpoint Account com caracter vazio",
			wantValue: 401,
			inData:    "",
		},
		{
			give:      "Testa o Endpoint Account com ID 123",
			wantValue: 500,
			inData:    "123",
		},
	}

	server := NewServerBank("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			//verificaTipoDoConteudo(t, resposta, tipoDoConteudoJSON)

			//recebido := resposta.Body.String()
			//assert.Equal(t, recebido, tt.wantValue)
			assert.Equal(t, resposta.Code, tt.wantValue)
		})

	}

}
*/
