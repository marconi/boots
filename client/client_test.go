package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectFailed(t *testing.T) {
	c := NewClient(
		"localhost:61613",
		"stomp",
		"user",
		"pass",
	)
	err := c.Connect()
	assert.NotNil(t, err)
	assert.Equal(t, c.Session, "", "session should be empty")
	assert.Equal(t, c.Heartbeat, "", "heart-beat should be empty")
}

func TestConnectSuccess(t *testing.T) {
	c := NewClient(
		"localhost:61613",
		"stomp",
		"admin",
		"admin",
	)
	err := c.Connect()
	assert.Nil(t, err)
	assert.NotEqual(t, c.Session, "", "session should not be empty")
	assert.NotEqual(t, c.Heartbeat, "", "heart-beat should not be empty")
}

func TestSendSuccess(t *testing.T) {
	c := NewClient(
		"localhost:61613",
		"stomp",
		"admin",
		"admin",
	)
	c.Connect()

	param := SendParam{
		Dest:  "/queue/hello-queue",
		Body:  "Hello!",
		Ctype: "text/plain",
	}
	err := c.Send(param)
	assert.Nil(t, err)
}
