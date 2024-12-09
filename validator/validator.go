package validator

import (
	"log"
	"strconv"
	"unicode"
)

// luhnAlgo implements the Luhn Algorithm to validate or calculate the check digit for a credit card number.
// @param numbers []int - The slice of integers representing the digits of the card number.
// @return int - The calculated Luhn checksum value.
func luhnAlgo(numbers []int) int {
	finalSum := 0

	// Iterate over the card number digits starting from the second-to-last digit
	for index := len(numbers) - 2; index >= 0; index = index - 1 {
		// Double every second digit from the right
		if (len(numbers)-index)%2 == 0 {
			numbers[index] *= 2
			// If the doubled value is greater than 9, subtract 9
			if numbers[index] > 9 {
				numbers[index] -= 9
			}
		}
		finalSum += numbers[index]
	}

	// Calculate the final checksum digit
	return (10 - (finalSum % 10)) % 10
}

// IsValidCreditCard validates a credit card number using the Luhn Algorithm.
// @param cardNumber string - The credit card number as a string.
// @return bool - Returns true if the card number is valid, false otherwise.
func IsValidCreditCard(cardNumber string) bool {
	log.Println("Entered valid function")
	var numbers []int

	// Convert the card number string into a slice of integers
	for _, value := range cardNumber {
		if unicode.IsDigit(value) {
			num, err := strconv.Atoi(string(value))
			if err != nil {
				return false // Return false if conversion fails
			}
			numbers = append(numbers, num)
		} else {
			log.Printf("This is not a digit %v \n", string(value))
			return false // Return false if a non-digit character is found
		}
	}

	// Extract the check digit (last digit of the card number)
	checkValue := numbers[len(numbers)-1]
	log.Printf("This is a checkValue %v \n", checkValue)

	// Compute the Luhn checksum
	luhnValue := luhnAlgo(numbers)
	log.Printf("This is a luhnValue %v \n", luhnValue)

	// Compare the computed checksum with the provided check digit
	return luhnValue == checkValue
}

// TypeOfCard identifies the type of a credit card based on its number prefix.
// @param cardNumber string - The credit card number as a string.
// @return string - Returns the type of the credit card (e.g., Visa, MasterCard).
func TypeOfCard(cardNumber string) string {
	// Map of fixed prefixes to their corresponding card types
	cardPrefixes := map[string]string{
		"34": "American Express", "37": "American Express",
		"5610": "Bankcard", "560221": "Bankcard", "560222": "Bankcard", "560223": "Bankcard", "560224": "Bankcard", "560225": "Bankcard",
		"31": "China T-Union",
		"62": "China UnionPay",
		"30": "Diners Club International", "36": "Diners Club International", "38": "Diners Club International", "39": "Diners Club International",
		"55":   "Diners Club United States & Canada",
		"6011": "Discover Card", "644": "Discover Card", "645": "Discover Card", "646": "Discover Card", "647": "Discover Card", "648": "Discover Card", "649": "Discover Card", "65": "Discover Card",
		"2200": "Mir", "2201": "Mir", "2202": "Mir", "2203": "Mir", "2204": "Mir", "2205": "BORICA",
		"5019": "Dankort", "4571": "Dankort (Visa co-branded)",
		"9792": "Troy",
		"4":    "Visa",
		"4026": "Visa Electron", "417500": "Visa Electron", "4508": "Visa Electron", "4844": "Visa Electron", "4913": "Visa Electron", "4917": "Visa Electron",
		"1":      "UATP",
		"506099": "Verve", "506198": "Verve", "650002": "Verve", "650027": "Verve", "507865": "Verve", "507964": "Verve",
		"357111": "LankaPay",
		"8600":   "UzCard", "5614": "UzCard",
		"9860": "Humo",
		"1946": "GPN",
	}

	// Ranges for card prefixes to their corresponding card types
	rangedCardPrefixes := []struct {
		RangeStart string
		RangeEnd   string
		CardType   string
	}{
		{"622126", "622925", "Discover Card"},
		{"2221", "2720", "Mastercard"},
		{"3528", "3589", "JCB"},
		{"6334", "6767", "Solo"},
		{"633110", "6333", "Switch"},
	}

	// Check fixed prefixes by iterating from the longest to the shortest prefix
	for length := 6; length > 0; length-- {
		if len(cardNumber) >= length {
			prefix := cardNumber[:length]
			if cardType, exists := cardPrefixes[prefix]; exists {
				return cardType
			}
		}
	}

	// Check ranged prefixes for matching card types
	for _, r := range rangedCardPrefixes {
		if len(cardNumber) >= len(r.RangeStart) {
			numberPrefix := cardNumber[:len(r.RangeStart)]
			if numberPrefix >= r.RangeStart && numberPrefix <= r.RangeEnd {
				return r.CardType
			}
		}
	}

	// Default to "Unknown" if no match is found
	return "Unknown"
}
