package main

import (
	"fmt"

	gameroom "github.com/go-workshop/ui-snake/game-room"
)

func main() {

	client := gameroom.NewClient()
	err := client.Connect(":8082")
	if err != nil {
		fmt.Printf("%q", err)
	}
	fmt.Println("Client Connected")

	defer client.Disconnect()

	//Read status of the game
	for {
		status, err1 := client.ReadStatus()
		if err1 != nil {
			fmt.Printf("%q", err)
			//Exist
			break
		}
		//Stop the game
		if !status.GameOver {
			fmt.Printf("%#v\n", status)
		}
	}
	//Read actions from console
	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')

	//err2 := client.SendAction(text)

	//if err2 != nil {
	//	fmt.Printf("%q", err)
	//}
}
