# Credit Card Validator

This project is a simple **Credit Card Validator** implemented in Go. It validates credit card numbers using the Luhn algorithm and identifies the type of card based on its number prefix.

## Project Structure

```
root
├── main.go              # Entry point of the application
├── validator
│   └── validator.go     # Contains logic for validation and card type detection
└── routes
    └── routes.go        # HTTP request handlers
```

---

## Features

- **Luhn Algorithm**: Verifies if the credit card number is valid.
- **Card Type Identification**: Identifies the type of credit card (e.g., Visa, Mastercard, etc.) based on its prefix.
- **REST API**: Provides an API endpoint to validate credit card numbers.

---

## Requirements

- **Go**: 1.17 or higher
- Any terminal or tool to run Go applications

---

## Installation

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/your-username/credit-card-validator.git
   cd credit-card-validator
   ```

2. **Install Dependencies**  
   The project uses the [Fiber](https://github.com/gofiber/fiber) framework. Run the following command to install it:  
   ```bash
   go get github.com/gofiber/fiber/v2
   ```

---

## Usage

1. **Run the Application**  
   From the project root directory, execute:  
   ```bash
   go run main.go
   ```

2. **API Endpoint**  
   - **URL**: `http://localhost:8080/api/v1/check`
   - **Method**: `POST`
   - **Request Body**: JSON format
     ```json
     {
       "card": "4111111111111111"
     }
     ```
   - **Response**: JSON format
     ```json
     {
       "isvalid": true,
       "company": "Visa"
     }
     ```

---

## Code Overview

### `main.go`

- Initializes the Fiber application.
- Defines the `/api/v1/check` POST route.
- Gracefully handles server shutdown on termination signals.

### `routes/routes.go`

- Contains the `CheckMail` function, which:
  - Parses the incoming request.
  - Validates the card using `validator.IsValidCreditCard`.
  - Identifies the card type using `validator.TypeOfCard`.
  - Sends the validation results as a JSON response.

### `validator/validator.go`

- Implements the **Luhn Algorithm** to check if a card number is valid.
- Identifies the card type based on its prefix using a predefined map and range.

---

## Example Request and Response

### Valid Card

**Request**:  
```bash
curl -X POST http://localhost:8080/api/v1/check \
-H "Content-Type: application/json" \
-d '{"card": "4111111111111111"}'
```

**Response**:  
```json
{
  "isvalid": true,
  "company": "Visa"
}
```

### Invalid Card

**Request**:  
```bash
curl -X POST http://localhost:8080/api/v1/check \
-H "Content-Type: application/json" \
-d '{"card": "1234567890123456"}'
```

**Response**:  
```json
{
  "isvalid": false,
  "company": "Unknown"
}
```

---

## Future Enhancements

- Support for additional card types and ranges.
- More detailed error handling.
- Add unit tests for validation logic.

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or new features.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
