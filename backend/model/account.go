package model

import "time"

type Account struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	UserName  string
}

type AccountList []Account

func NewAccount() *Account {
	return &Account{
		ID: NewULIDString(),
	}
}
