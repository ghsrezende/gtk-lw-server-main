package command

import (
	"github.com/howeyc/crc16"
)

func SendCommand(command []byte) ([]byte, error) {

	packet_lenght := byte(10 + len(command))
	command_lenght := byte(4 + len(command))

	comando := []byte(command)
	i := 11 + len(comando)

	message := []byte{0x78, 0x78}
	message = append(message, packet_lenght, 0x80, command_lenght, 0x00, 0x00, 0x00, 0x00)

	message = append(message, comando...)
	message = append(message, 0x00, 0x01)

	checksum := crc16.Checksum(message[2:i], crc16.CCITTTable)
	message = append(message, byte(checksum>>8), byte(checksum&0xff))
	message = append(message, 0x0D, 0x0A)

	return message, nil
}
