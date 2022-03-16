package m74bankapi

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetVersion(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
	}{
		{
			give:      "Test if get version OK",
			wantValue: "2022",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			version := GetVersion()
			assert.Equal(t, version[0:4], tt.wantValue)
			assert.Equal(t, version[0:4], tt.wantValue)

		})

	}

}

func TestSetAnGetIsProduction(t *testing.T) {

	tests := []struct {
		give      string
		wantValue bool
		inData    bool
	}{
		{
			give:      "test function with true",
			wantValue: true,
			inData:    true,
		},
		{
			give:      "test function with false",
			wantValue: false,
			inData:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			SetIsProduction(tt.inData)
			assert.Equal(t, GetIsProduction(), tt.wantValue)
		})

	}

}
