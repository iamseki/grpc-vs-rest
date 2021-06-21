package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iamseki/grpc-vs-rest/helper"
)

type Pricing struct {
	Symbol   string `json:"symbol,omitempty"`
	TimeFrom int64  `json:"timeFrom,omitempty"`
	TimeTo   int64  `json:"timeTo,omitempty"`
	Data     []Data `json:"data,omitempty"`
}

type Data struct {
	Time             int64   `json:"time,omitempty"`
	High             float32 `json:"high,omitempty"`
	Low              float32 `json:"low,omitempty"`
	Open             float32 `json:"open,omitempty"`
	Close            float32 `json:"close,omitempty"`
	ConversionType   string  `json:"conversionType,omitempty"`
	ConversionSymbol string  `json:"conversionSymbol,omitempty"`
}

func CryptoPricingHandler(w http.ResponseWriter, r *http.Request) {
	pricing := &Pricing{}

	helper.GetPricingFromJSONFile(pricing)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pricing)
}
