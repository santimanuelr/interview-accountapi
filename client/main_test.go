package client

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	GB                     = "GB"
	classificationPersonal = "Personal"
	version                = int64(0)
)

var client = NewAccountsClient("http://localhost:8080")

var id = uuid.New()

func TestPostAccount(t *testing.T) {

	//GIVEN
	var accountDummy = &AccountData{
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		ID:             id.String(),
		Type:           "accounts",
		Attributes: &AccountAttributes{
			BankID:                "400300",
			Name:                  []string{"santiago", "test"},
			BankIDCode:            "GBDSC",
			Bic:                   "NWBKGB22",
			Country:               &GB,
			BaseCurrency:          "GBP",
			Iban:                  "GB11NWBK40030041426818",
			AccountNumber:         "41426818",
			AccountClassification: &classificationPersonal,
		},
		Version: &version,
	}

	//WHEN
	accountResponse, err := client.Create(accountDummy)

	//THEN
	if err != nil {
		t.Errorf("Test request failed with error: '%v'", err.Error())
	}

	if accountResponse.Data.ID != id.String() {
		t.Errorf("Expected account ID: '%v' but got account ID: '%v'", id.String(), accountResponse.Data.ID)
	}

	assert.Equal(t, accountDummy, &accountResponse.Data)

}

func TestGetAccount(t *testing.T) {

	//GIVEN

	//WHEN
	accountResponse, err := client.Fetch(id.String())

	//THEN
	if err != nil {
		t.Errorf("Test request failed with error: '%v'", err.Error())
	}

	assert.Equal(t, id.String(), accountResponse.ID)

}
