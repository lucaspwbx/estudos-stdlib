package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//ExampleLookPath()
	CreateDir("foo")
	//defer RemoveDir("foo")
	CreateFile("foo", "filename.txt")
	CreateFile("foo", "filename2.txt")

	CreateDir("bar")
	CreateFile("bar", "teste.txt")
	defer RemoveDir("bar")
	PrintFile("bar", "teste.txt")
}

func ExampleLookPath() {
	path, err := exec.LookPath("fortune")
	if err != nil {
		log.Fatal("no fortune")
	}
	fmt.Printf("fortune is availabe at %s\n", path)
}

func CreateDir(path string) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		log.Fatal("problem creating dir")
	}
	fmt.Println("Directory created with success")
}

func RemoveDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal("Problem removing dir")
	}
}

func CreateFile(folder, file string) {
	f, err := os.Create(folder + "/" + file)
	if err != nil {
		log.Fatal("Problem creating file")
	}
	defer f.Close()
}

func PrintFile(folder, file string) {
	f, err := os.Open(folder + "/" + file)
	if err != nil {
		log.Fatal("Problem opening file")
	}
	defer f.Close()

	fi, _ := f.Stat()
	fmt.Println("Printing file info")
	fmt.Println("Name: ", fi.Name())
	fmt.Println("Size: ", fi.Size())
	fmt.Println("Mode: ", fi.Mode())
	fmt.Println("Modification time: ", fi.ModTime())
	fmt.Println("IsDir: ", fi.Mode().IsDir())
	fmt.Println("IsRegular: ", fi.Mode().IsRegular())
	fmt.Println("SameFile: ", os.SameFile(fi, fi))
}
