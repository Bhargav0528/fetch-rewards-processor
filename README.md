# Receipt Processor

Receipt Processor is a Go-based web service that processes receipts to calculate points based on specific rules. This service provides API endpoints to process receipts and retrieve points.

## Features

- Process receipts and calculate points based on retailer name, total amount, item descriptions, purchase date, and time.
- Retrieve points for a specific receipt.
- Fetch all receipts or total points by retailer name.

## Getting Started

### Prerequisites

- Go 1.20 or later
- Docker

### Installation

1. **Clone the Repository**

   ```bash
   git clone <repository-url>
   cd receipt-processor
   ```

2. **Install Dependencies**

   ```bash
   go mod download
   ```

### Running the Application

#### Locally

Run the application using the Go client:

```bash
go run ./rewards-processor
```

The server will start on port 8080.

#### Using Docker

1. **Build the Docker Image**

   ```bash
   docker build -t receipt-processor .
   ```

2. **Run the Docker Container**

   ```bash
   docker run -p 8080:8080 receipt-processor
   ```

### API Endpoints

- **Process a Receipt**

    - **Endpoint**: `POST /receipts/process`
    - **Request Body**: JSON object containing retailer, total, items, purchaseDate, and purchaseTime.
    - **Response**: JSON object containing the receipt ID.

- **Get Points for a Receipt**

    - **Endpoint**: `GET /receipts/{id}/points`
    - **Response**: JSON object containing the points for the specified receipt ID.

- **Fetch Receipts by Retailer Name**

    - **Endpoint**: `GET /receipts/retailer/{retailerName}/all`
    - **Response**: JSON array of receipts for the specified retailer name.

- **Get Total Points by Retailer Name**

    - **Endpoint**: `GET /receipts/retailer/{retailerName}/points`
    - **Response**: JSON object containing the total points for the specified retailer name.

### Example Usage

- **Process a Receipt**

  ```bash
  curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
      "retailer": "Walmart",
      "total": "25.50",
      "items": [
          {"shortDescription": "Banana", "price": "0.50"},
          {"shortDescription": "Apple", "price": "0.75"}
      ],
      "purchaseDate": "2023-10-15",
      "purchaseTime": "15:30"
  }'
  ```

- **Get Points for a Receipt**

  ```bash
  curl http://localhost:8080/receipts/{id}/points
  ```

- **Fetch Receipts by Retailer Name**

  ```bash
  curl http://localhost:8080/receipts/retailer/Walmart/all
  ```

- **Get Total Points by Retailer Name**

  ```bash
  curl http://localhost:8080/receipts/retailer/Walmart/points
  ```

### Additional Information

- This project uses the `uuid` package by Google for generating unique IDs for the receipts.