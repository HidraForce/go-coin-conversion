package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//Optimized for usability
	// fmt.Println("Welcome to the currency converter!")
	// fmt.Print("Enter the value you want to convert: ")

	// var value float64
	// var err error
	// for {
	// 	rawValue := readValue()
	// 	value, err = validateValue(rawValue)
	// 	if err == nil {
	// 		break
	// 	}
	// 	fmt.Println("Invalid value, please try again:")
	// }

	// availableCurrencies := getAvailableCurrencies()
	// fmt.Println("Available currencies:", availableCurrencies)
	// fmt.Print("Enter the currency you want to convert to: ")

	// var to string
	// for {
	// 	to = readCurrency()
	// 	to = strings.ToUpper(to)
	// 	if isValidCurrency(to, availableCurrencies) {
	// 		break
	// 	}
	// 	fmt.Println("Invalid currency, please try again:")
	// }

	// convertedValue := convertCurrency(value, to)
	// fmt.Printf("The converted value is: %s %s\n", convertedValue, to)





	
	//Made for challenge completion
	var rawValue string
	fmt.Scanln(&rawValue)
	value,_ := strconv.ParseFloat(rawValue, 64)
	var rawTo string
	fmt.Scanln(&rawTo)
	var to = strings.ToUpper(rawTo)
	result := convertCurrency(value, to)
	fmt.Printf("%s", result)



}