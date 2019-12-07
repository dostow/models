package models

import (
	"time"
)

type Settings struct {
	Signup       bool      `json:"signup"`
	ModifiedAt   time.Time `json:"modified_at"`
	Login        bool      `json:"login"`
	ID           string    `json:"id"`
	Enabled      bool      `json:"enabled"`
	CreatedAt    time.Time `json:"created_at"`
	BankAccounts []struct {
		Number  string `json:"number"`
		Name    string `json:"name"`
		Branch  string `json:"branch"`
		Bank    string `json:"bank"`
		Account string `json:"account"`
	} `json:"bankAccounts"`
	Info struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Address2 string `json:"address2"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
	} `json:"info"`
	Logo string `json:"logo"`
}

type SettingsResult struct {
	TotalCount int        `json:"total_count"`
	Count      int        `json:"count"`
	Data       []Settings `json:"data"`
}
