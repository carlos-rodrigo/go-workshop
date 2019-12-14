package client

import "fmt"

func main() {

	client := NewClient()
	err := client.Connect("127.0.0.1:8083")
	if err != nil {
		fmt.Printf("#v%q", err)
	}

	for {
		status, err := client.ReadStatus()
		if err != nil {
			fmt.Printf("#v%q", err)
		}
		//Here we should print the snakes in console
		fmt.Printf(status)

		if status.GameOver {
			break
		}
	}

	defer client.Disconnect()
}
