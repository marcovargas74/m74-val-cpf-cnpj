package m74bankAPI

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	bankAPI "github.com/marcovargas74/m74-bank-api"
)

func newReqGetPoints(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", name), nil)
	return request
}

func checkResult(t *testing.T, result, wait string) {
	t.Helper()
	if result != wait {
		t.Errorf(erroMsg, result, wait)
	}
}

func TestObterJogadores(t *testing.T) {
	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := newReqGetPoints("Maria")
		resposta := httptest.NewRecorder()

		bankAPI.ServidorJogador(resposta, requisicao)
	})

	t.Run("retornar resultado de Pedro", func(t *testing.T) {
		requisicao := newReqGetPoints("Pedro")
		resposta := httptest.NewRecorder()

		bankAPI.ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "10"

		checkResult(t, recebido, esperado)
	})

}
