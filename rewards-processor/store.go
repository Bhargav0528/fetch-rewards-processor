package main

import (
	"errors"
)

type Storage struct {
	idToReceipt   map[string]Receipt
	retailerToIDs map[string][]string
}

func NewStorage() *Storage {
	return &Storage{
		idToReceipt:   make(map[string]Receipt),
		retailerToIDs: make(map[string][]string),
	}
}

// AddNewReceipt adds a new receipt to the storage
func (s *Storage) AddNewReceipt(receipt Receipt) {
	s.idToReceipt[receipt.ID] = receipt
	s.retailerToIDs[receipt.Retailer] = append(s.retailerToIDs[receipt.Retailer], receipt.ID)
}

// RetrieveById retrieves a receipt by its ID
func (s *Storage) RetrieveById(id string) (Receipt, error) {
	receipt, exists := s.idToReceipt[id]
	if !exists {
		return Receipt{}, errors.New("receipt not found")
	}
	return receipt, nil
}

// RetrieveByRetailerName retrieves receipts by retailer name
func (s *Storage) RetrieveAllReceipts() []Receipt {
	var receipts []Receipt
	for _, receipt := range s.idToReceipt {
		receipts = append(receipts, receipt)
	}
	return receipts
}

// RetrieveByRetailerName retrieves receipts by retailer name
func (s *Storage) RetrieveByRetailerName(retailer string) ([]Receipt, error) {
	ids, exists := s.retailerToIDs[retailer]
	if !exists {
		return nil, errors.New("no receipts found for retailer")
	}

	var receipts []Receipt
	for _, id := range ids {
		if receipt, ok := s.idToReceipt[id]; ok {
			receipts = append(receipts, receipt)
		}
	}
	return receipts, nil
}
