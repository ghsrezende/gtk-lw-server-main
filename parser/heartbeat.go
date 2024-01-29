package parser

import (
	"fmt"

	"github.com/howeyc/crc16"
)

func responseStatus(buffer []byte) ([]byte, error) {
	if len(buffer) != 15 {
		return nil, fmt.Errorf("invalid heartbeat package length: %d", len(buffer))
	}

	// Initialize the response with the expected header and footer
	response := []byte{0x78, 0x78, 0x05, 0x13}

	// Extract the data from the input buffer and append it to the response
	response = append(response, buffer[9:10]...)

	// Calculate the CRC16 checksum and append it to the response
	checksum := crc16.Checksum(response[2:6], crc16.CCITTTable)
	response = append(response, byte(checksum>>8), byte(checksum&0xff))
	response = append(response, 0x0D, 0x0A)

	return response, nil
}
