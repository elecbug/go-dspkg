// Package net provides utilities for working with TCP connections and listeners.
package net

import (
	"net"
	"time"
)

// TCPClient represents a TCP client connection.
type TCPClient struct {
	net.Conn
}

// NewClient creates a new TCP client connection.
// It does not connect to any address; use DialTCP to establish a connection.
func NewClient() *TCPClient {
	return &TCPClient{}
}

// DialTCP connects to the TCP address with a timeout.
func (cli *TCPClient) DialTCP(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)

	if err != nil {
		return err
	}

	cli.Conn = conn

	// Set the read and write deadlines to the same timeout.
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		conn.Close()

		return err
	}

	return nil
}

// Close closes the TCP connection.
func (cli *TCPClient) Close() error {
	err := cli.Conn.Close()

	if err != nil {
		return err
	}

	return nil
}

// RemoteAddr returns the address of the TCP client.
func (cli *TCPClient) RemoteAddr() string {
	if cli.Conn == nil {
		return ""
	}

	addr := cli.Conn.RemoteAddr()

	if addr == nil {
		return ""
	}

	return addr.String()
}

// LocalAddr returns the local address of the TCP client connection.
// It returns an empty string if the connection is nil or the address is nil.
func (cli *TCPClient) LocalAddr() string {
	if cli.Conn == nil {
		return ""
	}

	addr := cli.Conn.LocalAddr()

	if addr == nil {
		return ""
	}

	return addr.String()
}

// Send sends data over the TCP connection.
func (cli *TCPClient) Send(data []byte) (int, error) {
	if cli.Conn == nil {
		return 0, net.ErrClosed
	}

	return cli.Conn.Write(data)
}

// Receive reads data from the TCP connection into the provided buffer.
func (cli *TCPClient) Receive(buffer []byte) (int, error) {
	if cli.Conn == nil {
		return 0, net.ErrClosed
	}

	return cli.Conn.Read(buffer)
}
