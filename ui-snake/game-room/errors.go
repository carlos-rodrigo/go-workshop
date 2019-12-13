package gameroom

import "errors"

var (
	ErrConnectionNotCreated  = errors.New("Connection not created")
	ErrCantReadStatus        = errors.New("Can't read status from gameroom server")
	ErrCantDecodeStatus      = errors.New("Can't decode game status")
	ErrCantConnectWithServer = errors.New("Can't Connect with Server")
)
