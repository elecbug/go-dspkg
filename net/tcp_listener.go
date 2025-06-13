// Package net provides utilities for working with TCP connections and listeners.
package net

import (
	"net"
	"time"
)

// TCPListener represents a TCP listener with an address.
type TCPListener struct {
	net.Listener
	addr string
}

// NewListener creates a new TCP listener for the specified address.
// The address should be in the format "host:port", e.g., "localhost:8080".
func NewListener(addr string) *TCPListener {
	return &TCPListener{
		addr: addr,
	}
}

// ListenTCP listens on the TCP address with a timeout.
func (li *TCPListener) ListenTCP(timeout time.Duration) error {
	listener, err := net.Listen("tcp", li.addr)

	if err != nil {
		return err
	}

	li.Listener = listener

	// Set the read and write deadlines to the same timeout.
	if err := listener.(*net.TCPListener).SetDeadline(time.Now().Add(timeout)); err != nil {
		listener.Close()

		return err
	}

	return nil
}

// Accept accepts a connection on the TCP listener with a timeout.
func (li *TCPListener) Accept(timeout time.Duration) (*TCPClient, error) {
	conn, err := li.Listener.Accept()

	if err != nil {
		return nil, err
	}

	// Set the read and write deadlines to the same timeout.
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		conn.Close()

		return nil, err
	}

	return &TCPClient{conn}, nil
}

// Close closes the TCP listener.
func (li *TCPListener) Close() error {
	err := li.Listener.Close()

	if err != nil {
		return err
	}

	return nil
}

// Addr returns the address of the TCP listener.
func (li *TCPListener) Addr() string {
	if li.Listener == nil {
		return li.addr
	}

	addr := li.Listener.Addr()

	if addr == nil {
		return li.addr
	}

	return addr.String()
}
