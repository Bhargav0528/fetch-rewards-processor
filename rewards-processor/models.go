package main

type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Points       int    `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

type GetPointsResponse struct {
	Points int `json:"points"`
}

type GetAllPointsResponse struct {
	ReceiptId string `json:"receiptId"`
	Points    int    `json:"points"`
}

type ReceiptRetailerNamePointsResponse struct {
	RetailerName string `json:"retailerName"`
	Points       int    `json:"points"`
}
