package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
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


var db *sql.DB

func initDB() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./shop.db"
	}

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Błąd otwarcia bazy:", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS products (
		id    INTEGER PRIMARY KEY AUTOINCREMENT,
		name  TEXT    NOT NULL,
		price REAL    NOT NULL
	);

	CREATE TABLE IF NOT EXISTS carts (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS cart_items (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		cart_id    INTEGER NOT NULL REFERENCES carts(id),
		product_id INTEGER NOT NULL,
		name       TEXT    NOT NULL,
		price      REAL    NOT NULL
	);

	CREATE TABLE IF NOT EXISTS payments (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		full_name  TEXT    NOT NULL,
		email      TEXT    NOT NULL,
		amount     REAL    NOT NULL,
		created_at TEXT    NOT NULL
	);
	`

	if _, err := db.Exec(schema); err != nil {
		log.Fatal("Błąd tworzenia schematu:", err)
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if count == 0 {
		seed := []Product{
			{Name: "Laptop", Price: 3999.99},
			{Name: "Mysz", Price: 99.99},
			{Name: "Klawiatura", Price: 199.99},
			{Name: "Monitor", Price: 899.99},
			{Name: "Kabel HDMI", Price: 19.99},
			{Name: "Słuchawki", Price: 349.99},
			{Name: "Kamera internetowa", Price: 249.99},
			{Name: "Dysk SSD", Price: 459.99},
		}
		for _, p := range seed {
			db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", p.Name, p.Price)
		}
		log.Println("Produkty dodane do bazy")
	}

	log.Println("Baza danych zainicjalizowana:", dbPath)
}


func allowedOrigin() string {
	if o := os.Getenv("ALLOWED_ORIGIN"); o != "" {
		return o
	}
	return "http://localhost:5173"
}

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin())
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

// ── Handlers ──────────────────────────────────────────────────────────────────

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, name, price FROM products ORDER BY id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price)
		products = append(products, p)
	}
	writeJSON(w, http.StatusOK, products)
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
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

		now := time.Now().Format("2006-01-02 15:04:05")
		res, err := db.Exec("INSERT INTO carts (created_at) VALUES (?)", now)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cartID, _ := res.LastInsertId()

		for _, item := range req.Items {
			db.Exec(
				"INSERT INTO cart_items (cart_id, product_id, name, price) VALUES (?, ?, ?, ?)",
				cartID, item.ID, item.Name, item.Price,
			)
		}

		log.Printf("Zapisano koszyk #%d z %d produktami", cartID, len(req.Items))
		writeJSON(w, http.StatusOK, map[string]any{
			"message": "Koszyk zapisany",
			"cart":    map[string]any{"id": cartID, "items": req.Items, "createdAt": now},
		})
		return
	}

	if r.Method == http.MethodGet {
		rows, err := db.Query("SELECT id, created_at FROM carts ORDER BY id DESC")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var carts []SavedCart
		for rows.Next() {
			var c SavedCart
			rows.Scan(&c.ID, &c.CreatedAt)

			itemRows, _ := db.Query(
				"SELECT product_id, name, price FROM cart_items WHERE cart_id = ?", c.ID,
			)
			for itemRows.Next() {
				var p Product
				itemRows.Scan(&p.ID, &p.Name, &p.Price)
				c.Items = append(c.Items, p)
			}
			itemRows.Close()
			carts = append(carts, c)
		}
		writeJSON(w, http.StatusOK, carts)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
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

		now := time.Now().Format("2006-01-02 15:04:05")
		res, err := db.Exec(
			"INSERT INTO payments (full_name, email, amount, created_at) VALUES (?, ?, ?, ?)",
			req.FullName, req.Email, req.Amount, now,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, _ := res.LastInsertId()

		log.Printf("Zapisano płatność #%d: %s %.2f PLN", id, req.Email, req.Amount)
		writeJSON(w, http.StatusOK, map[string]any{
			"message": "Płatność przyjęta",
			"payment": SavedPayment{
				ID: int(id), FullName: req.FullName,
				Email: req.Email, Amount: req.Amount, CreatedAt: now,
			},
		})
		return
	}

	if r.Method == http.MethodGet {
		rows, err := db.Query(
			"SELECT id, full_name, email, amount, created_at FROM payments ORDER BY id DESC",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var payments []SavedPayment
		for rows.Next() {
			var p SavedPayment
			rows.Scan(&p.ID, &p.FullName, &p.Email, &p.Amount, &p.CreatedAt)
			payments = append(payments, p)
		}
		writeJSON(w, http.StatusOK, payments)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}


func main() {
	initDB()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/cart", cartHandler)
	http.HandleFunc("/payments", paymentsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
