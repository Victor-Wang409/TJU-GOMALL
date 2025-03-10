package mq

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("nats connect failed")
		panic(err)
	}
}
