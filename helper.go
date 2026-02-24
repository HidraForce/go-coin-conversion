package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ExchangeData struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func readFile() ExchangeData {
	file, err := os.Open("currency.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data ExchangeData
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func getAvailableCurrencies() []string {
	data := readFile()
	currencies := make([]string, 0, len(data.Rates))
	for key := range data.Rates {
		currencies = append(currencies, key)
	}
	return currencies
}

func getChosenValue(chosenCurrency string) (float64, error) {
	data := readFile()
	rate, ok := data.Rates[chosenCurrency]
	if !ok {
		return 0, errors.New("Currency not found")
	}
	return rate, nil
}

func validateValue(value string) (float64, error) {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Println("Invalid number, please enter a valid number")
		return 0.0, err
	}
	return floatValue, nil
}

func convertCurrency(value float64, to string) string {
	rate, err := getChosenValue(to)
	if err != nil {
		log.Printf("Error getting chosen value: %v", err)
		return ""
	}

	valueConverted := value * rate
	return strconv.FormatFloat(valueConverted, 'f', 2, 64)
}

func readValue() string {
	var value string
	fmt.Scanln(&value)
	return value
}

func readCurrency() string {
	var currency string
	fmt.Scanln(&currency)
	return currency
}

func isValidCurrency(currency string, available []string) bool {
	for _, c := range available {
		if strings.EqualFold(c, currency) {
			return true
		}
	}
	return false
}
