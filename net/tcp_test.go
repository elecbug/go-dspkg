package net_test

import (
	"testing"
	"time"

	"github.com/elecbug/go-dspkg/net"
)

func TestTcp(t *testing.T) {
	li := net.NewListener("localhost:0")
	cli := net.NewClient()

	err := li.ListenTCP(5 * time.Second)

	if err != nil {
		t.Error("Failed to listen on TCP address")
		return
	}

	go func() {
		err := cli.DialTCP(li.Addr(), 5*time.Second)

		if err != nil {
			t.Errorf("Failed to dial TCP address: %v", err)
			return
		}
		defer cli.Close()

		cli.Send([]byte("Hello, TCP!"))

		select {}
	}()

	acceptedConn, err := li.Accept(5 * time.Second)

	if err != nil {
		t.Errorf("Failed to accept TCP connection: %v", err)
		return
	}
	defer acceptedConn.Close()

	buffer := make([]byte, 1024)

	data, err := acceptedConn.Receive(buffer)

	if err != nil {
		t.Errorf("Failed to receive data: %v", err)
		return
	}

	t.Logf("Received data: %s", string(buffer)[:data])
}
