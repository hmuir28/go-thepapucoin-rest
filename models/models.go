package models

type Transaction struct {
	Sender			string		`json:"sender"			validate:"required"`
	Recipient		string		`json:"recipient" 		validate:"required"`
	Amount			int			`json:"amount"			validate:"required"`
}
