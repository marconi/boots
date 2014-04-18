package constants

const (
	ACCEPT_VERSION = "1.2"
	EOL            = "\n"
	NULL           = '\x00' // null byte
)

var COMMANDS = map[string]string{
	"SEND":        SEND,
	"SUBSCRIBE":   SUBSCRIBE,
	"UNSUBSCRIBE": UNSUBSCRIBE,
	"BEGIN":       BEGIN,
	"COMMIT":      COMMIT,
	"ABORT":       ABORT,
	"ACK":         ACK,
	"NACK":        NACK,
	"DISCONNECT":  DISCONNECT,
	"CONNECT":     CONNECT,
	"CONNECTED":   CONNECTED,
	"MESSAGE":     MESSAGE,
	"RECEIPT":     RECEIPT,
	"ERROR":       ERROR,
}
