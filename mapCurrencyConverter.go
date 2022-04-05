package main

import "fmt"

//Declare Struct here
type currency struct {
	currency     string
	currencyName string
	rate         float64
}

//Declare Map here
var currencyMap map[string]currency

func main() {
	//Insert Main Code here
	//Activity #1
	currencyMap = make(map[string]currency)
	currencyMap["USD"] = currency{"USD", "US dollar", 1.1318}
	currencyMap["JPY"] = currency{"JPY", "Japanese yen", 121.05}
	currencyMap["GBP"] = currency{"GBP", "Pound sterling", 0.90630}
	currencyMap["CNY"] = currency{"CNY", "Chinese yuan renminbi", 7.9944}
	currencyMap["SGD"] = currency{"SGD", "Singapore dollar", 1.5743}
	currencyMap["MYR"] = currency{"MYR", "Malaysian ringgit", 4.8390}

	for code := range currencyMap {

		fmt.Printf("Currency Code: %v\nCurrency Name: %v\nExchange rate = 1 EUR: %v\n==================\n", currencyMap[code].currency, currencyMap[code].currencyName, currencyMap[code].rate)

	}

	//Activity #2 Currency conversion application

	var inputCurrency string
	var outputCurrency string
	var amountOfInputCurrency float64
	var euro float64
	var outputAmount float64

	fmt.Println("Please key in the currency to be exchange.")
	_, _ = fmt.Scanln(&inputCurrency)
	fmt.Println("How much do you want to exchange?")
	_, _ = fmt.Scanln(&amountOfInputCurrency)
	fmt.Println("Please key in which currency do you want.")
	_, _ = fmt.Scanln(&outputCurrency)

	//Test using range method
	for key := range currencyMap {

		if key == inputCurrency {
			//convert inputCurrency into euro
			euro = currencyMap[key].rate * amountOfInputCurrency
		}
	}
	for key := range currencyMap {
		if key == outputCurrency {
			outputAmount = currencyMap[key].rate / euro
		}
	}
	fmt.Printf("You get %.2f of %v\n", outputAmount, outputCurrency)

	//test school formula - (Amount of currency / Currency to convert from) * Currency to convert to
	result := (amountOfInputCurrency / currencyMap[inputCurrency].rate) * currencyMap[outputCurrency].rate
	fmt.Printf("You get %.2f of %v\n", result, outputCurrency)
}
