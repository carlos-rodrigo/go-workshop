package gameroom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewGameRoomClient() Client {
	return &GameRoom{}
}

func TestClientConnect(t *testing.T) {
	t.Run("Connect to Gameroom that not exist should fail", func(t *testing.T) {
		client := NewGameRoomClient()
		url := "127.0.0.1:8082"

		err := client.Connect(url)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrCantConnectWithServer)

		defer client.Disconnect()
	})

	t.Run("Connect to Gameroom that exist should fail", func(t *testing.T) {
		client := NewGameRoomClient()
		url := "127.0.0.1:8083"

		err := client.Connect(url)

		assert.Nil(t, err)
		defer client.Disconnect()
	})
}

func TestClientReadStatus(t *testing.T) {
	t.Run("Read status from connected client should return a status", func(t *testing.T) {
		client := NewGameRoomClient()
		url := "127.0.0.1:8083"
		client.Connect(url)

		status, err := client.ReadStatus()

		assert.NotNil(t, err)
		assert.NotNil(t, status)
		defer client.Disconnect()
	})

	t.Run("Read status from unconnected client should return an error", func(t *testing.T) {
		client := NewGameRoomClient()

		_, err := client.ReadStatus()

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrConnectionNotCreated)

		defer client.Disconnect()
	})
}

func TestClientDisconnect(t *testing.T) {
	t.Run("Close an opened connection should not return an error", func(t *testing.T) {
		client := NewGameRoomClient()
		url := "127.0.0.1:8083"
		client.Connect(url)

		err := client.Disconnect()

		assert.Nil(t, err)

		defer client.Disconnect()
	})

	t.Run("Close a not opened connection should fail", func(t *testing.T) {
		client := NewGameRoomClient()

		err := client.Disconnect()

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrConnectionNotCreated)

		defer client.Disconnect()
	})
}
