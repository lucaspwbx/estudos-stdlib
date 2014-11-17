package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io/ioutil"
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

func main() {
	base64Example()
}
