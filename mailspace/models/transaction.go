package models

import "time"

// Gateway a gateway for payment
type Gateway struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TransactionRequest make a transaction for an invoice
type TransactionRequest struct {
	ID      string `json:"id"`
	Invoice string `json:"invoice"`
	Amount  int64  `json:"amount"`
	Method  string `json:"method"` //Transaction method, paypal, check ...
}

// Transaction an invoice transaction
type Transaction struct {
	ID         string    `json:"id"`
	Amount     int64     `json:"amount"`
	Invoice    Invoice   `json:"invoice"`
	Gateway    Gateway   `json:"gateway"`
	Applied    bool      `json:"applied"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Status     string    `json:"status"`
}
