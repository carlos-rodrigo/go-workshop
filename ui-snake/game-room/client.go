package gameroom

import (
	"encoding/json"
	"net"
)

type Client interface {
	Connect(s string) error
	ReadStatus() (SnakeStatus, error)
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

	var rawStatus []byte
	_, errRead := g.conn.Read(rawStatus)

	if errRead != nil {
		return SnakeStatus{}, ErrCantReadStatus
	}

	var status SnakeStatus
	errUnmarshal := json.Unmarshal(rawStatus, &status)

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
