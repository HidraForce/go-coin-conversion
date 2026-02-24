package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type Data struct {
	Currency string
	Value    string
}

func readFile() map[string]Data {
	currency, err := os.Open("currency.txt")
	if err != nil {
		panic(err)
	}
	defer currency.Close()
	dataPopulate := make(map[string]Data)
	scanData := bufio.NewScanner(currency)
	for scanData.Scan() {
		line := scanData.Text()
		splitLine := strings.Split(line, ";")
		if len(splitLine) == 2 {
			dataPopulate[splitLine[0]] = Data{
				Currency: splitLine[0],
				Value:    splitLine[1],
			}
		}
	}

	return dataPopulate
}

func getAvailableCurrencies() []string {
	data := readFile()
	currencies := make([]string, 0, len(data))
	for key := range data {
		currencies = append(currencies, key)
	}
	return currencies
}

func getChosenValue(chosenCurrency string) (string, error) {
	data := readFile()
	for _, value := range data {
		if value.Currency == chosenCurrency {
			return value.Value, nil
		}
	}
	return "", errors.New("Currency not found")
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
	coinS, err := getChosenValue(to)
	if err != nil {
		log.Printf("Error getting chosen value: %v", err)
		return ""
	}

	coinF, err := strconv.ParseFloat(coinS, 64)
	if err != nil {
		log.Printf("Error parsing coin value: %v", err)
		return ""
	}
	valueConverted := value * coinF

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
		if c == currency {
			return true
		}
	}
	return false
}
