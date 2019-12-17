package gameroom

import (
	"bufio"
	"encoding/json"
	"net"
	"strings"
)

type Client interface {
	Connect(s string) error
	ReadStatus() (SnakeStatus, error)
	SendAction(s string) error
	Disconnect() error
}

type GameRoom struct {
	conn net.Conn
}

func NewClient() GameRoom {
	return GameRoom{}
}

func (g *GameRoom) Connect(s string) error {
	conn, err := net.Dial("tcp", s)

	if err != nil {
		return ErrCantConnectWithServer
	}

	g.conn = conn

	return nil
}

func (g *GameRoom) ReadStatus() (SnakeStatus, error) {
	if g.conn == nil {
		return SnakeStatus{}, ErrConnectionNotCreated
	}

	stringStatus, errRead := bufio.NewReader(g.conn).ReadString('\n')

	if errRead != nil {
		return SnakeStatus{}, ErrCantReadStatus
	}

	stringStatus = strings.ReplaceAll(stringStatus, "\n", "")
	var status SnakeStatus
	errUnmarshal := json.Unmarshal([]byte(stringStatus), &status)

	if errUnmarshal != nil {
		return SnakeStatus{}, ErrCantDecodeStatus
	}

	return status, nil
}

func (g *GameRoom) Disconnect() error {
	if g.conn == nil {
		return ErrConnectionNotCreated
	}

	err := g.conn.Close()

	if err != nil {
		return err
	}

	return nil
}

func (g *GameRoom) SendAction(s string) error {
	if g.conn == nil {
		return ErrConnectionNotCreated
	}

	_, err := g.conn.Write([]byte(s))

	if err != nil {
		return err
	}

	return nil
}
