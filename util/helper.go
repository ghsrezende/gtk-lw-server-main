package util

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func convertHexToBinary(hex string) []byte {
	// Create a new byte array to store the binary representation of the hexadecimal string.
	binary := make([]byte, len(hex)/2)

	// Iterate over the hexadecimal string and convert each character to its corresponding binary value.
	for i := 0; i < len(hex); i += 2 {
		binary[i/2] = hexToBin(hex[i : i+2])
	}

	// Return the binary array.
	return binary

}

func ConvertToHexArray(word string) ([]byte, error) {
	// Convert the word to HEX
	hexWord := hex.EncodeToString([]byte(word))

	// Convert the HEX string to a byte array
	hexArray, err := hex.DecodeString(hexWord)
	if err != nil {
		return nil, err
	}

	return hexArray, nil
}

func hexToBin(hex string) byte {
	// Convert the hexadecimal string to an integer.
	value, err := strconv.ParseInt(hex, 16, 8)
	if err != nil {
		return 0
	}

	// Return the binary representation of the integer value.
	return byte(value)

}

func arrayToBson(array []interface{}) bson.M {
	bsonMap := make(bson.M)
	for i, value := range array {
		bsonMap[fmt.Sprintf("array_%d", i)] = value
	}
	return bsonMap
}
