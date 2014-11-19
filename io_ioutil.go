package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//reading from file using ReadAll
	file, err := os.Open("io_ioutil.go")
	if err != nil {
		log.Fatal(err)
	}
	retorno, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(retorno))

	//reading from buffer
	bla := bufio.NewReader(strings.NewReader("readingfrombuffer"))
	retorno2, _ := ioutil.ReadAll(bla)
	fmt.Println(string(retorno2))

	//reading from file using ReadFile
	data, _ := ioutil.ReadFile("main4.go")
	fmt.Println(string(data))

	err = ioutil.WriteFile("temp.go", []byte("golang rules"), 0666)
	if err != nil {
		log.Fatal(err)
	}
	retorno3, _ := ioutil.ReadFile("temp.go")
	fmt.Println(string(retorno3))
}
