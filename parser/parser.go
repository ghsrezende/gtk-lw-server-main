package parser

func Parse(buffer []byte) (string, []byte, error) {

	var response []byte
	var imei string
	var err error

	typePacket := buffer[3]

	switch typePacket {
	case 0x01: // Login Information
		imei, response, err = responseLogin(buffer)

		if err != nil {
			return "", response, err
		}

		return imei, response, nil

	case 0x13: // Status Information (The heartbeat packets)
		response, err = responseStatus(buffer)
		if err != nil {
			return "", response, err
		}

		return "", response, nil

	case 0x15: // Alarm data
		return "", nil, err

	case 0x16: // Alarm data
		response, err = responseAlarm(buffer)
		return "", response, err
		//return "", nil, err

	case 0x17: // Location data
		response, err = responseLocation(buffer)
		return "", response, err
		//return "", nil, err

	default:
		//return "", nil, fmt.Errorf("invalid package type: %X", typePacket)
		//return "", buffer, err
		return "", nil, err
	}
}
