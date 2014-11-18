package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	unixSock()
	udp()
	readFromUdp()
	writeToUdp()
}

func unixSock() {
	//for unix, unixgram and unixpacket
	addr, err := net.ResolveUnixAddr("unix", "127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr)           // 127.0.0.1
	fmt.Println(addr.Network()) // unix

	//forcing error
	addr, err = net.ResolveUnixAddr("tcp", "127.0.0.1")
	if err != nil {
		fmt.Println(err) //unknown network tcp
	}
}

func udp() {
	//may be udp, udp4 or udp6 or ""(tcp)
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:80")
	if err != nil {
		fmt.Println("Error resolving UDP address: ", err)
		return
	}
	fmt.Println(addr)           //127.0.0.1:80
	fmt.Println(addr.Network()) //udp
}

func readFromUdp() {
	ra, err := net.ResolveUDPAddr("udp", "127.0.0.1:7")
	if err != nil {
		log.Fatalln(err)
	}
	la, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	c, err := net.ListenUDP("udp", la)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	_, err = c.WriteToUDP([]byte("a"), ra)
	if err != nil {
		log.Fatal(err)
	}

	err = c.SetDeadline(time.Now().Add(100 * time.Millisecond))
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 1)
	_, _, err = c.ReadFromUDP(b)
	if err == nil {
		log.Fatal(err)
	}
}

func writeToUdp() {
	l, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	c, err := net.Dial("udp", l.LocalAddr().String())
	if err != nil {
		log.Fatalln("Dial failed ", err)
	}
	defer c.Close()

	ra, err := net.ResolveUDPAddr("udp", l.LocalAddr().String())
	if err != nil {
		log.Fatalln("ResolveUDPAddr failed")
	}
	_, err = c.(*net.UDPConn).WriteToUDP([]byte("connection oriented mode socket"), ra)
	if err == nil {
		log.Fatalln("WriteToUDP should fail")
	}
	if err != nil && err.(*net.OpError).Err != net.ErrWriteToConnected {
		log.Fatalln()
	}

	_, err = c.(*net.UDPConn).WriteTo([]byte("connection oriented mode socket"), ra)
	if err == nil {
		log.Fatalln("WriteTo should fail")
	}
	if err != nil && err.(*net.OpError).Err != net.ErrWriteToConnected {
		log.Fatalln()
	}

	_, err = c.Write([]byte("Connection oriented mode socket"))
	if err != nil {
		fmt.Println("Write failed", err)
	}

	//no previous connection

	c2, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatalln("Listen packet failed")
	}
	defer c2.Close()

	ra, err = net.ResolveUDPAddr("udp", c2.LocalAddr().String())
	if err != nil {
		log.Fatalln("Resolve udp addr failed: ", err)
	}
	_, err = c2.(*net.UDPConn).WriteToUDP([]byte("Connection-less mode socket"), ra)
	if err != nil {
		log.Fatalln("WriteToUDP failed")
	}

	_, err = c2.WriteTo([]byte("connection less mode socket"), ra)
	if err != nil {
		log.Fatalln("writeTo failed")
	}

	_, err = c2.(*net.UDPConn).Write([]byte("connectionless mode socket"))
	if err == nil {
		log.Fatalln("Write should fail")
	}
}
