package main

import (
	"encoding/json"
	"github.com/go-openapi/strfmt"
	"github.com/gopdf/models"
	"log"
	"testing"
)


const value = `{
      "code" : "00001014",
      "guid" : "ab8ceedf-ec74-4c31-bca0-49a39c3c8fcd",
      "placed" : "2018-12-04T14:01:13+0000",
      "status" : "READY",
      "statusDisplay" : "processing",
      "total" : {
         "currencyIso" : "USD",
         "formattedValue" : "$260.78",
         "priceType" : "BUY",
         "value" : 260.78
      }
   }`

func TestUnmarshalOrder(t *testing.T) {
	order := models.OrderWsDTO{}
	err := json.Unmarshal([]byte(value), &order)
	if err != nil {
		t.Error(err)
		t.Fail()
	}


}
func TestParseDatetime(t *testing.T) {

	reg := strfmt.NewFormats()
	reg.Add("datetime", &models.CustomDateTime{}, nil)
	result, err := reg.Parse("datetime", "2018-12-04T14:01:13+0000")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	log.Printf("Result %v", result)
}