# Zestaw10 – Shop App

Go backend z SQLite3 + React/Vite frontend.

## Uruchomienie lokalne

```bash
# Backend
cd server
go mod tidy
go run main.go

# Frontend (nowy terminal)
cd shop-client
npm install
npm run dev
```


## Endpointy API

| Metoda | URL          | Opis                    |
|--------|--------------|-------------------------|
| GET    | /health      | Status serwera          |
| GET    | /products    | Lista produktów         |
| GET    | /cart        | Zapisane koszyki        |
| POST   | /cart        | Zapisz koszyk           |
| GET    | /payments    | Historia płatności      |
| POST   | /payments    | Złóż płatność           |
