package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ConversionRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type ConversionResponse struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          string `json:"amount"`
	ConvertedAmount string `json:"convertedAmount"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/convert", convertHandler).Methods("POST")

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	var request ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedAmount, err := convertCurrency(request.From, request.To, request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ConversionResponse{
		From:            request.From,
		To:              request.To,
		Amount:          request.Amount,
		ConvertedAmount: convertedAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func convertCurrency(from string, to string, amount string) (string, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", from, to)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching conversion rate:", err)
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error parsing conversion rate response:", err)
		return "", err
	}

	rate, ok := result[from][to]
	if !ok {
		return "", fmt.Errorf("conversion rate not found")
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", fmt.Errorf("invalid amount")
	}

	convertedAmount := amountFloat * rate

	return fmt.Sprintf("%f", convertedAmount), nil
}
