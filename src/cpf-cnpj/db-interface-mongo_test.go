package cpfcnpj

import (
	"testing"
)

func TestCreateDBMongo(t *testing.T) {

	tests := []struct {
		give      string
		wantValue bool
	}{
		{
			give:      "Test Open MONGO DB",
			wantValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			result := CreateDBMongo(false)
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}
