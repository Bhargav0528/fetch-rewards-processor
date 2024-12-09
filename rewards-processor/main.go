package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)

var store = NewStorage()

/*
Process Receipt method accepts receipts and calculates rewards based on various parameters

	After calculating rewards, it creates a new ID for the receipt and stores the receipt with its generated id
*/
func processReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	var receipt Receipt
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in reading the request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &receipt); err != nil {
		http.Error(w, "Unable to process receipt; Please check the field names and data types.", http.StatusBadRequest)
		return
	}

	fmt.Printf("Receipt processed :: %+v \n", receipt)

	receipt.ID = uuid.New().String()
	receipt.Points = calculatePoints(receipt)

	fmt.Printf("ID assigned to receipt :: %s \n", receipt.ID)

	store.AddNewReceipt(receipt)

	w.Header().Set("Content-Type", "application/json")
	response := ProcessReceiptResponse{ID: receipt.ID}
	json.NewEncoder(w).Encode(response)
}

/*
GetPoints method retrieves the points for the given receipt ID stored in the system
*/
func getPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	fmt.Printf("Received request with ID :: %s \n", id)

	receipt, err := store.RetrieveById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := GetPointsResponse{Points: receipt.Points}
	json.NewEncoder(w).Encode(response)
}

/*
getPointsForAllReceipts generates a list of receiptID and corresponding points for all the receipts stored in the system
*/
func getPointsForAllReceipts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Received request for fetching points for all receipts")

	var result []GetAllPointsResponse
	receipts := store.RetrieveAllReceipts()

	for _, receipt := range receipts {
		result = append(result, GetAllPointsResponse{
			ReceiptId: receipt.ID,
			Points:    receipt.Points,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

/*
getReceiptsByRetailerName retrieves the receipts with all its contents for the given retailer Name
*/
func getReceiptsByRetailerName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	retailer := r.PathValue("retailerName")
	receipts, err := store.RetrieveByRetailerName(retailer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipts)
}

/*
getTotalPointsByRetailerName retrieves the aggregate sum of points for each of the receipt having the given retailerName
*/
func getTotalPointsByRetailerName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	retailer := r.PathValue("retailerName")
	receipts, err := store.RetrieveByRetailerName(retailer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	totalPoints := 0

	for _, receipt := range receipts {
		totalPoints += receipt.Points
	}

	w.Header().Set("Content-Type", "application/json")
	response := ReceiptRetailerNamePointsResponse{RetailerName: retailer, Points: totalPoints}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/receipts/process", processReceipt)
	http.HandleFunc("/receipts/{id}/points", getPoints)
	http.HandleFunc("/receipts/points", getPointsForAllReceipts)
	http.HandleFunc("/receipts/retailer/{retailerName}/all", getReceiptsByRetailerName)
	http.HandleFunc("/receipts/retailer/{retailerName}/points", getTotalPointsByRetailerName)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
