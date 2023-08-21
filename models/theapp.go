package models

import "time"

type APIResponseMessage struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Loan struct {
	ID         int         `json:"id"`
	Amount     float64     `json:"amount"`
	Term       int         `json:"term"`
	Status     string      `json:"status"`
	Repayments []Repayment `json:"repayments"`
}

type Repayment struct {
	ID      int       `json:"id"`
	LoanID  int       `json:"loan_id"`
	Amount  float64   `json:"amount"`
	DueDate time.Time `json:"due_date"`
	Status  string    `json:"status"`
}
