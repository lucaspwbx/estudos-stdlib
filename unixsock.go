package main

import (
	"net"
	"os"
)

func main() {
	os.Remove("/tmp/unixdomaincli")
	laddr := net.UnixAddr{"/tmp/unixdomaincli", "unix"}
	conn, err := net.DialUnix("unix", &laddr, &net.UnixAddr{"/tmp/unixdomain1", "unix"})
	if err != nil {
		panic(err)
	}
	defer os.Remove("/tmp/unixdomaincli")
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}
	conn.Close()
}
