package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CartRequest struct {
	Items []Product `json:"items"`
}

type SavedCart struct {
	ID        int       `json:"id"`
	Items     []Product `json:"items"`
	CreatedAt string    `json:"createdAt"`
}

type PaymentRequest struct {
	FullName string  `json:"fullName"`
	Email    string  `json:"email"`
	Amount   float64 `json:"amount"`
}

type SavedPayment struct {
	ID        int     `json:"id"`
	FullName  string  `json:"fullName"`
	Email     string  `json:"email"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"createdAt"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 3999.99},
	{ID: 2, Name: "Mysz", Price: 99.99},
	{ID: 3, Name: "Klawiatura", Price: 199.99},
	{ID: 4, Name: "Monitor", Price: 899.99},
	{ID: 5, Name: "Kabel HDMI", Price: 19.99},
}

var savedCarts []SavedCart
var savedPayments []SavedPayment

var cartID = 1
var paymentID = 1

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var req CartRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		newCart := SavedCart{
			ID:        cartID,
			Items:     req.Items,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		savedCarts = append(savedCarts, newCart)
		cartID++

		log.Println("Zapisano koszyk:", newCart)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Koszyk zapisany",
			"cart":    newCart,
		})
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(savedCarts)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var req PaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		newPayment := SavedPayment{
			ID:        paymentID,
			FullName:  req.FullName,
			Email:     req.Email,
			Amount:    req.Amount,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		savedPayments = append(savedPayments, newPayment)
		paymentID++

		log.Println("Zapisano płatność:", newPayment)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Płatność przyjęta",
			"payment": newPayment,
		})
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(savedPayments)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/cart", cartHandler)
	http.HandleFunc("/payments", paymentsHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}