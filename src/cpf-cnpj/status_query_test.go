package cpfcnpj

import (
	"testing"
	"time"
)

func TestCreateStatus(t *testing.T) {

	tests := []struct {
		give                string
		wantTotalQueryValue uint64
	}{
		{
			give:                "Testa Se Criou Status Como Zero",
			wantTotalQueryValue: 0,
		},
	}

	CreateStatus()
	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			CheckIfEqualInt(t, GetNumQuery(), tt.wantTotalQueryValue)
		})

	}

}

func TestUpdateStatus(t *testing.T) {

	tests := []struct {
		give                string
		wantTotalQueryValue uint64
	}{
		{
			give:                "Testa Se Incrementou o numero de consultas_1",
			wantTotalQueryValue: 1,
		},
		{
			give:                "Testa Se Incrementou o numero de consultas_2",
			wantTotalQueryValue: 2,
		},
		{
			give:                "Testa Se Incrementou o numero de consultas_3",
			wantTotalQueryValue: 3,
		},
		{
			give:                "Testa Se Incrementou o numero de consultas_4",
			wantTotalQueryValue: 4,
		},
	}

	CreateStatus()
	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			UpdateStatus()
			CheckIfEqualInt(t, GetNumQuery(), tt.wantTotalQueryValue)
		})

	}

}

func TestUpTimeStatus(t *testing.T) {

	CreateStatus()
	time.Sleep(3 * time.Second)
	lastTime := GetUptimeQuery()
	CheckIfUptimeIsOK(t, lastTime, 3)
}
