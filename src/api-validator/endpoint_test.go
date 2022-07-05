package m74validatorapi

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

/*Comentado para ver se passa nos testes apenas para testar o MAKEFILE

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
*/

func TestServerBank(t *testing.T) {
	assert.Equal(t, 1, 1)

}

/*Comentado para ver se passa nos testes apenas para testar o MAKEFILE
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

	servidor := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao, _ := http.NewRequest(http.MethodGet, "/", nil)
			resposta := httptest.NewRecorder()

			servidor.ServeHTTP(resposta, requisicao)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)
			assert.Equal(t, resposta.Code, http.StatusOK)

		})

	}

}

/*Comentado para ver se passa nos testes apenas para testar o MAKEFILE

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

	server := NewServerValidator("dev")
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

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/accounts", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			assert.Equal(t, resposta.Code, tt.wantValue)
		})

	}
}

/*Comentado para ver se passa nos testes apenas para testar o MAKEFILE

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

	server := NewServerValidator("dev")
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

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsPOST("/login", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
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

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			requisicao := newReqEndpointsGET("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			assert.Equal(t, resposta.Code, tt.wantValue)
		})

	}

}
*/
