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
	gameStatus := make(chan gameroom.GameStatus)
	gameErrors := make(chan error)
	//Read status of the game
	for {
		go readGameStatus(client, gameStatus, gameErrors)

		go printGameStatus(gameStatus)

		go errorHandling(gameErrors)

		go recieveAction(client, gameErrors)
	}
}

func readGameStatus(client gameroom.Client, gameStatus chan gameroom.GameStatus, errors chan error) {
	status, err := client.ReadStatus()
	if err != nil {
		fmt.Printf("%q", err)
		errors <- err
	}
	gameStatus <- status
}

func printGameStatus(channel chan gameroom.GameStatus) {
	status := <-channel

	fmt.Printf("%#v\n", status)
}

func errorHandling(errors chan error) {
	err := <-errors

	fmt.Printf("An error was founded, the game session is terminated")
	panic(err)
}

func recieveAction(client gameroom.Client, errors chan error) {
	var action string
	fmt.Scanf("%s", &action)

	err := client.SendAction(action)

	if err != nil {
		errors <- err
	}
}
