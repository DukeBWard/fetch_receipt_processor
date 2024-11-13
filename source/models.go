package main

type Receipt struct {
	ID           string `json:"id"`
	userId       string `json:"userId"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
	Points       int    `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type User struct {
	ID           string `json:"id"`
	ReceiptCount int    `json:"receiptCount"`
}
