package client

import (
	"github.com/google/uuid"
	"testing"
)

var (
	GB                     = "GB"
	classificationPersonal = "Personal"
)

func TestPostAccount(t *testing.T) {

	//GIVEN
	id := uuid.New()

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
	}

	//WHEN
	accountResponse, err := post(accountDummy)

	//THEN
	if err != nil {
		t.Errorf("Test request failed with error: '%v'", err.Error())
	}

	if accountResponse.Data.ID != id.String() {
		t.Errorf("Expected account ID: '%v' but got account ID: '%v'", id.String(), accountResponse.Data.ID)
	}

}

func TestGetAccount(t *testing.T) {

	//GIVEN
	id := uuid.New()

	//WHEN
	accountResponse, err := Get("accountDummy")

	//THEN
	if err != nil {
		t.Errorf("Test request failed with error: '%v'", err.Error())
	}

	if len(accountResponse.Data) == 0 {
		t.Errorf("Expected account ID: '%v' but got account ID: '%v'", id.String(), accountResponse.Data)
	}

}
