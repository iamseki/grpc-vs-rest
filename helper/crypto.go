package helper

import (
	"log"
	"path/filepath"
)

func GetPricingFromJSONFile(target interface{}) {
	file := filepath.Join("dataset", "eth-prices.json")

	err := JSONFileToStruct(file, target)
	if err != nil {
		log.Fatalln(err)
	}
}
