package dlink

import (
	// "fmt"
	// "os"
	"time"

	"github.com/fatih/color"
	"github.com/ziutek/telnet"
	"log"
)

const timeout = 30 * time.Second

type Telnet struct {
	connection  *telnet.Conn
	destination string
}

func NewTelnet(destination string) (*Telnet, error) {
	t, err := telnet.Dial("tcp", destination)
	if err != nil {
		return nil, err
	}

	t.SetUnixWriteMode(true)

	return &Telnet{
		connection:  t,
		destination: destination,
	}, nil
}

// func (t *Telnet) CheckErr(err error) {
// 	if err != nil {
// 		log.Println("[Error]:", err)
// 	}
// }

func (t *Telnet) Reconnect() error {
	newT, err := telnet.Dial("tcp", t.destination)
	if err != nil {
		return err
	}

	newT.SetUnixWriteMode(true)

	t.connection = newT
	return nil
}

func (t *Telnet) ExpectWithError(d ...string) error {
	// fmt.Println("Expecting: ", d)
	err := t.connection.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}

	err = t.connection.SkipUntil(d...)
	if err != nil {
		return err
	}
	return nil
}

func (t *Telnet) SendlnWithError(s string) error {
	// fmt.Println("Sending: ", s)
	err := t.connection.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}

	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err = t.connection.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

func (t *Telnet) Expect(s string) {
	err := t.ExpectWithError(s)
	if err != nil {
		log.Printf(color.RedString("Error: %s\n"), err)
	}
}

func (t *Telnet) Sendln(s string) {
	err := t.SendlnWithError(s)
	if err != nil {
		log.Printf(color.RedString("Error: %s\n"), err)
	}
}

// func (t *Telnet) retryableExpect(count int, d ...string) {
// 	err := t.Expect(d...)
//
// 	if err != nil {
// 		if err == io.EOF {
// 			log.Println("Trying to reconnect...")
// 			t.reconnect()
//
// 		}
// 		log.Println(err)
// 		log.Println("Sleeping for: ", count)
// 		time.Sleep(time.Second * time.Duration(count))
// 		t.retryableExpect(count+3, d...)
// 	}
//
// }
//
// func (t *Telnet) RetryableExpect(d ...string) {
// 	t.retryableExpect(0, d...)
// }
//
// func (t *Telnet) retryableSendln(count int, d ...string) {
// 	err := t.Expect(d...)
//
// 	if err != nil {
// 		log.Println(err)
// 		log.Println("Sleeping for: ", count)
// 		time.Sleep(time.Second * time.Duration(count))
// 		t.retryableSendln(count+3, d...)
// 	}
//
// }
//
// func (t *Telnet) RetryableSendln(d ...string) {
// 	t.retryableSendln(0, d...)
// }
