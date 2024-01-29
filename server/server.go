package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"gtk-lw-server-main/command"
	"gtk-lw-server-main/log"
	"gtk-lw-server-main/parser"
	"gtk-lw-server-main/util"
)

const maxReconnectAttempts = 10

type Client struct {
	conn              net.Conn
	mutex             sync.RWMutex
	reconnectAttempts int
	reconnectSignal   chan struct{}
}

func HandleClient(conn net.Conn) {
	client := &Client{
		conn:            conn,
		reconnectSignal: make(chan struct{}, 1),
	}
	defer func() {
		client.CloseConnection()
	}()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ReceiveToDevice(client)
	}()

	go func() {
		defer wg.Done()
		SendCommand(client)
	}()

	for range client.reconnectSignal {
		err := HandleReconnection(client)
		if err != nil {
			break
		}

		go SendCommand(client)
	}

	wg.Wait()
}

func ReceiveToDevice(client *Client) {
	var deviceID string

	buffer := make([]byte, 8192)

	defer func() {
		client.CloseConnection()
	}()

	for {
		n, err := client.conn.Read(buffer)
		if err != nil {
			log.Error("Error reading from connection: ", err)
			fmt.Printf("Error reading from connection: %v\n", err)
			break
		}

		if n == 0 {
			log.Info("Client closed the connection")
			break
		}

		client.mutex.Lock()

		data := buffer[:n]

		imei, response, err := parser.Parse(data)
		if err != nil {
			log.Error(err)
			client.mutex.Unlock()
			break
		}

		if response != nil {
			//err := sendToDevice(client, response)

			if client.conn == nil {
				log.Info("Client closed")
			}

			_, err := client.conn.Write(response)
			if err != nil {
				select {
				case client.reconnectSignal <- struct{}{}:
				default:
				}
			}

			if err != nil {
				log.Error("Error writing to device: ", err)
				fmt.Printf("Error writing to device: %v\n", err)
				client.mutex.Unlock()
				break
			}
		}

		if imei != "" {
			deviceID = imei
		}

		log.PacketReceived(deviceID, "%02X", data)
		fmt.Printf("Package received: %02X imei: %v\n", data, deviceID)

		if response != nil {
			log.PacketSent(deviceID, "%02X", response)
			fmt.Printf("Package sent: %02X imei: %v\n", response, deviceID)
		}

		client.mutex.Unlock()
	}
}

func SendCommand(client *Client) {
	for {
		client.mutex.Lock()
		if client.reconnectAttempts > 0 {
			client.mutex.Unlock()
			time.Sleep(1 * time.Second)
			continue
		}
		client.mutex.Unlock()

		command, err := readCommand()
		if err != nil {
			log.Error("Error reading from terminal: ", err)
			fmt.Printf("Error reading from terminal: %v\n", err)
			continue
		}

		if err := sendToDevice(client, command); err != nil {
			log.Error("Error writing to device: ", err)
			fmt.Printf("Error writing to device: %v\n", err)
			break
		}

		log.PacketSent("", "%02X", command)
		fmt.Printf("Package sent: %02X\n", command)
	}
}

func readCommand() ([]byte, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	log.Info("Typed command: ", scanner.Text())
	fmt.Println("Typed command: ", scanner.Text())

	hexData, err := util.ConvertToHexArray(scanner.Text())
	if err != nil {
		log.Error("Error decoding hex string: ", err)
		return nil, err
	}

	commandToSend, err := command.SendCommand(hexData)
	if err != nil {
		log.Error("Error processing command string: ", err)
		return nil, err
	}

	return commandToSend, nil
}

func sendToDevice(client *Client, message []byte) error {

	client.mutex.Lock()
	defer client.mutex.Unlock()

	if client.conn == nil {
		log.Info("Client closed")
		return fmt.Errorf("client closed")
	}

	_, err := client.conn.Write(message)
	if err != nil {
		select {
		case client.reconnectSignal <- struct{}{}:
		default:
		}

		return err
	}

	return nil
}

func HandleReconnection(client *Client) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if client.reconnectAttempts >= maxReconnectAttempts {
		log.Error("Maximum reconnection attempts reached for client")
		fmt.Printf("Maximum reconnection attempts reached for client\n")
		return fmt.Errorf("maximum reconnection attempts reached")
	}

	client.reconnectAttempts++
	log.Info("Attempting to reconnect to client...")
	fmt.Printf("Attempting to reconnect to client...\n")

	time.Sleep(2 * time.Second)

	newConn, err := net.Dial("tcp", client.conn.RemoteAddr().String())
	if err != nil {
		log.Error("Reconnection failed:", err)
		fmt.Printf("Reconnection failed: %v\n", err)
		return err
	}

	client.conn = newConn
	client.reconnectAttempts = 0
	log.Info("Device reconnected successfully")
	fmt.Printf("Device reconnected successfully\n")

	return nil
}

func (c *Client) CloseConnection() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.conn != nil {
		c.conn.Close()
	}
}
