package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const apiURL = "https://api.exchangerate-api.com/v4/latest/"

func main() {
	fmt.Println("==== Welcome to GO Currency Converter ====")
	fmt.Println("Tip: Type 'list' to see available currencies.")
	fmt.Println("Tip: Type 'exit' anytime to quit.\n")

	for {
		baseCurrency := inputCurrency("Enter BASE currency (e.g., USD, EUR): ")
		if baseCurrency == "" {
			continue
		}

		targetCurrency := inputCurrency("Enter TARGET currency (e.g., NPR, INR): ")
		if targetCurrency == "" {
			continue
		}

		amount := inputAmount("Enter AMOUNT to convert: ")
		if amount < 0 {
			continue
		}

		rate, _, err := getExchangeRate(baseCurrency, targetCurrency)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		converted := amount * rate
		fmt.Printf("\n💱 %.2f %s = %.2f %s\n\n", amount, baseCurrency, converted, targetCurrency)

		fmt.Print("Do you want to convert again? (yes/no): ")
		var again string
		fmt.Scanln(&again)
		if strings.ToLower(strings.TrimSpace(again)) != "yes" {
			fmt.Println("Thanks for using GO Currency Converter!!. Dveveloped by Sarjak Khanal")
			break
		}
	}
}

func inputCurrency(prompt string) string {
	for {
		fmt.Print(prompt)
		var currency string
		fmt.Scanln(&currency)
		currency = strings.ToUpper(strings.TrimSpace(currency))

		if currency == "EXIT" {
			fmt.Println("Thanks for using GO Currency Converter!!. Dveveloped by Sarjak Khanal")
			os.Exit(0)
		} else if currency == "LIST" {
			showAvailableCurrencies()
			continue
		} else if len(currency) != 3 {
			fmt.Println(" Please enter a valid 3-letter currency code (e.g., USD, INR).")
			continue
		}
		return currency
	}
}

func inputAmount(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		amount, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil || amount < 0 {
			fmt.Println("Invalid amount. Please enter a positive number.")
			continue
		}
		return amount
	}
}

func getExchangeRate(baseCurrency, targetCurrency string) (float64, map[string]float64, error) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL + baseCurrency)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Rates map[string]float64 `json:"rates"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, nil, fmt.Errorf("failed to parse data: %v", err)
	}

	rate, ok := data.Rates[targetCurrency]
	if !ok {
		return 0, nil, fmt.Errorf("target currency %s not found", targetCurrency)
	}

	return rate, data.Rates, nil
}

func showAvailableCurrencies() {
	fmt.Println("\n Some countries currencies are listed below here :")
	fmt.Println("USD (US Dollar)")
	fmt.Println("EUR (Euro)")
	fmt.Println("INR (Indian Rupee)")
	fmt.Println("NPR (Nepalese Rupee)")
	fmt.Println("GBP (British Pound)")
	fmt.Println("AUD (Australian Dollar)")
	fmt.Println("CAD (Canadian Dollar)")
	fmt.Println("JPY (Japanese Yen)")
	fmt.Println("CNY (Chinese Yuan)")

}
