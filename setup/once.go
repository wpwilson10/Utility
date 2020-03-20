package setup

import (
	"fmt"
	"net"
	"strings"
)

// Once binds to the given port on localhost and does nothing forever.
// Useful for when we want a single instances of an app at a time.
func Once(port int) error {
	// point to port on localhost
	address := fmt.Sprintf("127.0.0.1:%d", port)
	// connect to port
	listener, err := net.Listen("tcp", address)
	// failed, return error
	if err != nil {
		return err
	}
	// send listener off to wait forever
	go silent(listener)
	// done
	return nil
}

func silent(listener net.Listener) {
	// hang out and do nothing forever
	for {
		listener.Accept()
	}
}

// CheckOnce returns true if the given port is busy, false otherwise.
func CheckOnce(port int) bool {
	// point to port on localhost
	address := fmt.Sprintf("127.0.0.1:%d", port)
	// try to connect to port
	_, err := net.Listen("tcp", address)
	// failed, check error
	if err != nil {
		// if we got the expected single connection error, given port is busy
		if strings.Index(err.Error(), "Only one usage of each socket address") != -1 {
			return true
		}
		// got an unexpected error
		LogCommon(err).Fatal("Unexpected error")
	}
	// nothing using the port
	return false

}
