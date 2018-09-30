package dlink

import (
	"fmt"
	"log"
	// "os"
	"time"

	"github.com/ziutek/telnet"
)

const timeout = 30 * time.Second

type Telnet struct {
	connection *telnet.Conn
}

func NewTelnet(destination string) (*Telnet, error) {
	t, err := telnet.Dial("tcp", destination)
	if err != nil {
		return nil, err
	}

	t.SetUnixWriteMode(true)

	return &Telnet{
		connection: t,
	}, nil
}

func (t *Telnet) CheckErr(err error) {
	if err != nil {
		log.Println("[Error]:", err)
	}
}

func (t *Telnet) Expect(d ...string) {
	fmt.Println("Expecting: ", d)
	t.CheckErr(t.connection.SetReadDeadline(time.Now().Add(timeout)))
	t.CheckErr(t.connection.SkipUntil(d...))
}

func (t *Telnet) Sendln(s string) {
	fmt.Println("Sending: ", s)
	t.CheckErr(t.connection.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.connection.Write(buf)
	t.CheckErr(err)
}
