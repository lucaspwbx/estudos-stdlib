package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

func binaryExample() {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed", err)
	}
	fmt.Printf("% x", buf.Bytes())

	fmt.Println()

	buf2 := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf2, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%x", buf2.Bytes())

	fmt.Println()

	var pi2 float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf3 := bytes.NewReader(b)
	err = binary.Read(buf3, binary.LittleEndian, &pi2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(pi2)

	fmt.Println()
}

func base64Example() {
	data := []byte("old data")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)
	data2, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error decoding: ", err)
		return
	}
	fmt.Printf("%q\n", data2)

	//using new encoder
	input := []byte("metallica")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close()

	fmt.Println()

	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	var res2 []byte
	res2, _ = ioutil.ReadAll(decoder)
	fmt.Println(string(res2))
}

func csvExample() {
	//writing
	b := &bytes.Buffer{}
	f := csv.NewWriter(b)
	err := f.WriteAll([][]string{{"abc"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	out := b.String()
	fmt.Println(out)

	//writing
	b2 := &bytes.Buffer{}
	f2 := csv.NewWriter(b2)
	f2.Write([]string{"bcd"})
	f2.Flush()
	fmt.Println(f2.Error())

	//writing using errorWriter
	f2 = csv.NewWriter(errorWriter{})
	f2.Write([]string{"bcd"})
	f2.Flush()
	fmt.Println(f2.Error())
}

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func exampleGob() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Threehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode erro 1")
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2: ", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}

type errorWriter struct{}

func (e errorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Test")
}

func main() {
	base64Example()
	csvExample()
	exampleGob()
}
