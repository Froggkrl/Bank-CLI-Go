package models

type Account struct {
	ID      string
	Name    string
	Balance float64
}

var Accounts = make(map[string]Account)
