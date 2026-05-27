package main

import (
	"net"

	"github.com/DarthPestilane/easytcp"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var packer easytcp.Packer

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	packer = easytcp.NewDefaultPacker()
}

func main() {
	senderClient()
	for i := 0; i < 10; i++ {
		readerClient(i)
	}

	select {}
}

func establish() (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

func senderClient() { _ = "STUB: not implemented"; return }

// send

// read

func readerClient(id int) { _ = "STUB: not implemented"; return }
