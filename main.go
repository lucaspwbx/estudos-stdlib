package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	bufStrReader := bufio.NewReader(strings.NewReader("golang"))

	retorno, err := bufStrReader.Peek(6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(retorno)) //peek 6 first bytes

	retorno, err = bufStrReader.Peek(10)
	if err != nil {
		fmt.Println(err) //EOF
	}

	retorno, err = bufStrReader.Peek(50)
	if err != nil {
		fmt.Println(err)
	}

	var teste []byte
	teste = make([]byte, 10)
	n, err := bufStrReader.Read(teste)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)                 //6
	fmt.Println(string(teste[:n])) //prints golang

	bufByteReader := bufio.NewReader(bytes.NewReader([]byte("lucas")))
	c, err := bufByteReader.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(c)) //prints l

	err = bufByteReader.UnreadByte()
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err = bufByteReader.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(c))

	r, _, err := bufByteReader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r)) //prints `u`

	value := bufByteReader.Buffered()
	fmt.Println(value) //3 bytes are on the buffer

	retorno2, err := bufByteReader.ReadSlice(byte('\n'))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(retorno2)) //print cas

	bufByteReader.Reset(strings.NewReader("helloween"))

	retorno3, err := bufByteReader.ReadSlice(byte('z'))
	if err != nil && err == io.EOF {
		fmt.Println("OK, encontrou EOF")
	}
	fmt.Println(string(retorno3))

	//TODO - Reset

	bufByteReader.Reset(strings.NewReader("novastring"))
	retorno4, _, err := bufByteReader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(retorno4)) //prints novastring

	bufByteReader.Reset(strings.NewReader("testeblablaokrulesheahaehaeheaheaaeheaheaheah"))
	retorno5, prefix, err := bufByteReader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	if prefix {
		fmt.Println(string(retorno5)) //should be true
	}

	bufByteReader.Reset(strings.NewReader("heaheahbea"))
	teste, err = bufByteReader.ReadBytes(byte('b'))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(teste))

	bufByteReader.Reset(strings.NewReader("testando"))
	bla, err := bufByteReader.ReadString(byte('a'))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bla)

	bufByteReader.WriteTo(os.Stdout)

	fmt.Println()

	//buffered output
	bufWriter := bufio.NewWriter(os.Stdout)

	//TODO - flush and reset, available, buffered
	_, err = bufWriter.Write([]byte("inter"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Buffered: ", bufWriter.Buffered())
	fmt.Println("Available bytes in the buffer: ", bufWriter.Available())

	//writing one more byte
	err = bufWriter.WriteByte(byte('c'))

	//writing one more rune
	_, err = bufWriter.WriteRune(rune('g'))
	if err != nil {
		fmt.Println(err)
	}

	//writing string
	_, err = bufWriter.WriteString("teste")
	if err != nil {
		fmt.Println(err)
	}

	//Reading from io.Reader
	_, err = bufWriter.ReadFrom(strings.NewReader("gremiosucks"))
	if err != nil {
		fmt.Println(err)
	}
	err = bufWriter.Flush()

	fmt.Println()

	//buffered input and output
	readWriter := bufio.NewReadWriter(bufio.NewReader(strings.NewReader("okteste")), bufio.NewWriter(os.Stdout))
	readWriter.Write([]byte("blalala"))
	err = readWriter.Flush()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()

	//buffer
	buffer := bytes.NewBuffer([]byte("dalessandro"))

	fmt.Println("Unread bytes: ", string(buffer.Bytes()))

	fmt.Println("String: ", buffer.String())

	fmt.Println("Length: ", buffer.Len())

	//discards all but the first n bytes
	buffer.Truncate(3)
	fmt.Println("String: ", buffer.String())

	//write to buffer
	t, err := buffer.Write([]byte("novastring"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote %d bytes", t)

	fmt.Println()

	t2, err := buffer.WriteString("fdps")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote %d bytes", t2)

	fmt.Println()
	t3, err := buffer.ReadFrom(strings.NewReader("recover"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote %d bytes", t3)
	fmt.Println()

	buffer.WriteByte(byte('g'))
	buffer.WriteRune(rune('t'))
	//buffer.WriteTo(os.Stdout)

	//var readSlice []byte
	readSlice := make([]byte, 10)
	bla2, err := buffer.Read(readSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read %d bytes into slice\n", bla2)
	fmt.Println("Word: ", string(readSlice[:n]))

	//TODO - netx, readbyte, readrune, unreadrune, unreadbyte, readbytes, readstring
	fmt.Println()
	//testTcp()
	//testUnixSocket()
}

func testUnixSocket() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		unixServer()
	}()
	go func() {
		defer wg.Done()
		unixClient()
	}()
	wg.Wait()
}

func unixClient() {
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

func unixServer() {
	os.Remove("/tmp/unixdomain1")
	l, err := net.ListenUnix("unix", &net.UnixAddr{"/tmp/unixdomain1", "unix"})
	if err != nil {
		panic(err)
	}
	fmt.Println("listening socket")
	defer os.Remove("/tmp/unixdomain1")

	for {
		conn, err := l.AcceptUnix()
		if err != nil {
			panic(err)
		}
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(buf[:n]))
		conn.Close()
	}
}

func testTcp() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		server()
	}()
	go func() {
		defer wg.Done()
		client()
	}()
	wg.Wait()
}

func server() {
	//tcp server
	host := "localhost"
	port := "3333"

	addr := fmt.Sprintf("%s:%s", host, port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	fmt.Println("listening")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			return
		}
		go handleRequest(conn)
	}
}

func client() {
	conn2, err := net.Dial("tcp", "localhost"+":"+"3333")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := conn2.Write([]byte("teste"))
	buf := make([]byte, 128)
	n, err = conn2.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf[:n]))
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 128)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte("messagr received"))
	if err != nil {
		fmt.Println(err)
	}
}
