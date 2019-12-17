package main

import (
	"encoding/json"
	"fmt"
	"net"

	gameroom "github.com/go-workshop/ui-snake/game-room"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8082")
	fmt.Printf(ln.Addr().String())

	defer ln.Close()
	for {
		// accept connection on port

		conn, _ := ln.Accept()
		// will listen for message to process ending in newline (\n)
		gamestatus := gameroom.SnakeStatus{
			Board:     gameroom.Board{20, 20},
			PlayerOne: []gameroom.Position{gameroom.Position{1, 1}},
			PlayerTwo: []gameroom.Position{gameroom.Position{1, 1}},
			Fruit:     gameroom.Position{2, 2},
			GameOver:  false,
			Winner:    "El negro de whatsapp",
		}

		bytesStatus, _ := json.Marshal(gamestatus)
		// send new string back to client
		stringStatus := string(bytesStatus)
		// adding delimiter
		stringStatus = stringStatus + "\n"

		conn.Write([]byte(stringStatus))
	}
}
