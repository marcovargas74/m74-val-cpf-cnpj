package m74validatorapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

const (
	UserAgentTest = "self_test"
)

func newReqEndpointsGET(urlPrefix, urlName string) *http.Request {
	request, error := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", urlPrefix, urlName), nil)
	if error != nil {
		panic(error)
	}

	request.Header.Set("User-Agent", UserAgentTest)
	fmt.Printf("endpoint: %v\n", request.URL)
	return request
}

func newReqEndpointsPOST(urlPrefix, urlName string) *http.Request {
	request, error := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", urlPrefix, urlName), nil)
	if error != nil {
		panic(error)
	}

	request.Header.Set("User-Agent", UserAgentTest)
	return request
}

/*
func newReqEndpointsBodyPOST(urlPrefix, bodyData string) *http.Request {

	jsonData := []byte(bodyData)
	request, error := http.NewRequest(http.MethodPost, urlPrefix, bytes.NewBuffer(jsonData))
	if error != nil {
		panic(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	fmt.Printf("endpoint: %v body: %v\n", request.URL, bodyData)
	return request
}*/

func TestServerAPI(t *testing.T) {
	assert.Equal(t, 1, 1)

}

func TestServerApi_default(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "Nobody return result",
			wantValue: "Endpoint not found",
			inData:    "Nobody",
		},
	}

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			answer := httptest.NewRecorder()

			server.ServeHTTP(answer, request)

			received := answer.Body.String()
			assert.Equal(t, received, tt.wantValue)
			assert.Equal(t, answer.Code, http.StatusOK)

		})
	}

}

func TestServerApi_status(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
	}{
		{
			give:      "status Endpoint test",
			wantValue: "{\"num_total_query\":0,\"start_time\":\"0001-01-01T00:00:00Z\",\"up_time\":9223372036.854776}",
		},
	}

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			request, _ := http.NewRequest(http.MethodGet, "/status", nil)
			request.Header.Set("User-Agent", UserAgentTest)
			answer := httptest.NewRecorder()

			server.ServeHTTP(answer, request)

			//received := answer.Body.String()
			//assert.Equal(t, received, tt.wantValue)
			assert.Equal(t, answer.Code, http.StatusOK)

		})

	}

}

func TestCallbackCpfsGET(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "cpfs Endpoint test with empty char",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "cpfs Endpoint test with NOBODY",
			wantValue: 406,
			inData:    "Nobody",
		},
		{
			give:      "cpfs Endpoint test with cnpj",
			wantValue: 404,
			inData:    "36.562.098/0001-18",
		},
		/*{
			give:      "cpfs Endpoint test with cpf 111.111.111-11",
			wantValue: 200,
			inData:    "111.111.111-11",
		},*/
	}

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			request := newReqEndpointsGET("/cpfs", tt.inData)
			answer := httptest.NewRecorder()

			server.ServeHTTP(answer, request)
			assert.Equal(t, answer.Code, tt.wantValue)
		})

	}

}

func TestCallbackCpfsPost(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "cpfs Endpoint test with NOBODY",
			wantValue: 400,
			inData:    "Nobody",
		},
		{
			give:      "cpfs Endpoint test with empty char",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "cpfs Endpoint test with CNPJ",
			wantValue: 404,
			inData:    "36.562.098/0001-18",
		},
		/*{
			give:      "cpfs Endpoint test with cpf 111.111.111-11",
			wantValue: 200,
			inData:    "111.111.111-11",
		},*/
	}

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			request := newReqEndpointsPOST("/cpfs", tt.inData)
			answer := httptest.NewRecorder()

			server.ServeHTTP(answer, request)
			assert.Equal(t, answer.Code, tt.wantValue)
		})

	}
}

/*
func TestCallbackCnpjGET(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "cnpjs Endpoint test with empty char",
			wantValue: 404,
			inData:    "",
		},
		/*{
			give:      "cnpjs Endpoint test with NOBODY",
			wantValue: 404,
			inData:    "Nobody",
		},
		{
			give:      "cnpjs Endpoint test with empty char",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "cnpjs Endpoint test with complete CPF",
			wantValue: 404,
			inData:    "111.111.111-49",
		},
		{
			give:      "cnpjs Endpoint test with incomplete cnpj",
			wantValue: 406,
			inData:    "36.562.098",
		},
		{
			give:      "cnpjs Endpoint test with not formated cnpj",
			wantValue: 406,
			inData:    "36562098000118",
		},*/
/*{
			give:      "cnpjs Endpoint test with complete cnpj",
			wantValue: 200,
			inData:    "36.562.098/0001-18",
		},* /

	}

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			request := newReqEndpointsGET("/cnpjs", tt.inData)
			answer := httptest.NewRecorder()

			server.ServeHTTP(answer, request)
			assert.Equal(t, answer.Code, tt.wantValue)
		})

	}

}


func TestCallbackCnpjPost(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "cnpjs Endpoint test with NOBODY",
			wantValue: 404,
			inData:    "Nobody",
		},
		{
			give:      "cnpjs Endpoint test with empty char",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "cnpjs Endpoint test with cpf",
			wantValue: 400,
			inData:    "111.111.111-49",
		},
		/*{
			give:      "cnpjs Endpoint test with cpf 36.562.098/0001-18",
			wantValue: 400,
			inData:    "36.562.098/0001-18",
		},* /
	}

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			request := newReqEndpointsPOST("/cnpjs", tt.inData)
			answer := httptest.NewRecorder()

			server.ServeHTTP(answer, request)
			assert.Equal(t, answer.Code, tt.wantValue)
		})

	}
}

/*
func TestCallbackCpfsPostBody(t *testing.T) {

	tests := []struct {
		give      string
		wantValue int
		inData    string
	}{
		{
			give:      "cpfs Endpoint test with NOBODY",
			wantValue: 404,
			inData:    "Nobody",
		},
		{
			give:      "cpfs Endpoint test with empty char",
			wantValue: 404,
			inData:    "",
		},
		{
			give:      "cpfs Endpoint test with cpf 111.111.111-11",
			wantValue: 404,
			inData:    " \"cpf\": \"111.111.111-11\" ",
		},
	}

	server := NewServerValidator("dev")
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {

			request := newReqEndpointsBodyPOST("/cpfs", tt.inData)
			answer := httptest.NewRecorder()

			fmt.Printf("endpoint TO SERVER: %v body: %v\n", request.URL, request.Body)
			server.ServeHTTP(answer, request)
			assert.Equal(t, answer.Code, tt.wantValue)
		})

	}
}*/
