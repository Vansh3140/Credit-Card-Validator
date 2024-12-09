package routes

import (
	"encoding/json"

	"github.com/Vansh3140/credit-card-validator/validator" // Importing custom validator package for credit card validation
	"github.com/gofiber/fiber/v2"                          // Importing Fiber framework for HTTP request handling
)

// UserRequest represents the structure for the incoming request payload
type UserRequest struct {
	Card string `json:"card"` // Credit card number sent in the request
}

// UserResponse represents the structure for the outgoing response payload
type UserResponse struct {
	IsValid bool   `json:"isvalid"` // Indicates if the credit card number is valid
	Company string `json:"company"` // The company of the credit card (e.g., Visa, MasterCard)
}

// CheckMail handles POST requests to validate a credit card number
// @param c *fiber.Ctx - The Fiber context for handling the HTTP request and response
// @return error - Returns an error if processing fails, otherwise nil
func CheckMail(c *fiber.Ctx) error {
	// Parse the incoming JSON request body into the UserRequest struct
	cardNumber := new(UserRequest)
	err := json.Unmarshal(c.Body(), cardNumber)
	if err != nil {
		// Return a 500 status code if there's an error parsing the request
		return c.Status(500).SendString(err.Error())
	}

	// Create a new UserResponse instance to hold the validation results
	response := new(UserResponse)

	// Validate the credit card number and determine the company
	response.IsValid = validator.IsValidCreditCard(cardNumber.Card)
	response.Company = validator.TypeOfCard(cardNumber.Card)

	// Respond with a 200 status code and the validation results in JSON format
	return c.Status(200).JSON(response)
}
