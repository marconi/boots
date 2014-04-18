package client

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/marconi/boots"
	"github.com/marconi/boots/constants"
)

type Client struct {
	Host      string
	Vhost     string
	Login     string
	Passcode  string
	Session   string
	Heartbeat string
	Conn      net.Conn
}

func NewClient(host, vhost, login, passcode string) *Client {
	return &Client{
		Host:     host,
		Vhost:    vhost,
		Login:    login,
		Passcode: passcode,
	}
}

// connects to stomp server and initiates handshake
func (c *Client) Connect() error {
	headers := boots.Headers{
		"accept-version": constants.ACCEPT_VERSION,
		"host":           c.Vhost,
		"login":          c.Login,
		"passcode":       c.Passcode,
	}
	frame := boots.NewFrame(constants.CONNECT, headers, "")

	// stablish tcp connection to server
	log.Println("connecting to broker...")
	conn, err := net.Dial("tcp", c.Host)
	if err != nil {
		return err
	}
	c.Conn = conn // keep track of tcp connection
	log.Println("connected")

	// send CONNECT frame
	log.Println("initiating handshake...")
	count, err := fmt.Fprintf(c.Conn, frame.Build())
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("%d bytes written", count))

	// process response
	log.Println("waiting for reply...")
	response, err := bufio.NewReader(c.Conn).ReadString(constants.NULL)
	if err != nil {
		return err
	}

	responseFrame, err := boots.ParseFrame(response)
	if err != nil {
		return err
	}

	if responseFrame.Command != constants.CONNECTED {
		return errors.New(fmt.Sprintf("Unable to connect: %s", responseFrame.Body))
	}

	// keep track of session and heartbeat
	c.Session = responseFrame.Headers["session"]
	c.Heartbeat = responseFrame.Headers["heart-beat"]

	log.Println("client ready")
	return nil
}

func (c *Client) Send(param SendParam) error {
	headers := boots.Headers{
		"destination":    param.Dest,
		"content-type":   param.Ctype,
		"content-length": strconv.Itoa(len(param.Body)),
	}
	headers.Append(param.ExtraHead)
	frame := boots.NewFrame(constants.SEND, headers, param.Body)

	// send SEND frame
	log.Println("sending...")
	count, err := fmt.Fprintf(c.Conn, frame.Build())
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("%d bytes written", count))

	// process response
	log.Println("waiting for reply...")

	// if no error within 1 second, the request went fine
	c.Conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	response, err := bufio.NewReader(c.Conn).ReadString(constants.NULL)
	if err != nil {
		if strings.Contains(err.Error(), "timeout") {
			c.Conn.SetReadDeadline(time.Time{}) // reset deadline
			log.Println("sent")
			return nil
		} else {
			return err
		}
	}

	responseFrame, err := boots.ParseFrame(response)
	if err != nil {
		return err
	}

	return errors.New(fmt.Sprintf("Unable to send: %s", responseFrame.Body))
}

// parameter for SEND frame
type SendParam struct {
	Dest      string
	Body      string
	Ctype     string
	ExtraHead boots.Headers
}
