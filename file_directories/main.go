package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//ExampleLookPath()
	CreateDir("foo")
	defer RemoveDir("foo")
	CreateFile("foo", "filename.txt")
	CreateFile("foo", "filename2.txt")

	CreateDir("bar")
	CreateFile("bar", "teste.txt")
	defer RemoveDir("bar")
	PrintFile("bar", "teste.txt")
	GetWorkDirectory()
	WriteStringToFile()
	WriteBytesToFile()
	ReadContentFromFile()
	ReadContentFromFileAt()
	WriteAfterOffset()
	CreateDirDois()
}

func CreateDirDois() {
	err := os.Mkdir("testando", 0777)
	if err != nil {
		log.Fatalln("Problems creating dir: ", err)
	}
	fmt.Println("Dir created")

	//changing working directory
	err = os.Chdir("testando")
	if err != nil {
		log.Fatalln("Problem changing dir: ", err)
	}
	fmt.Println(os.Getwd()) //current dir is now testando

	_, err = os.Create("blabla.txt")
	if err != nil {
		log.Fatalln("Problem creatingg file: ", err)
	}
	err = os.Rename("blabla.txt", "ok.txt")
	if err != nil {
		log.Fatalln("bla ")
	}
	err = os.Chdir("..")
	fmt.Println(os.Getwd())
	os.RemoveAll("testando")
}

func WriteAfterOffset() {
	f, err := os.OpenFile("teste.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//seeking end of file
	offset, err := f.Seek(0, 2)
	if err != nil {
		fmt.Println("problem getting offset")
		return
	}
	_, err = f.WriteAt([]byte("helloworld"), offset)
	if err != nil {
		log.Fatal(err)
	}

	//reading after write
	var buf []byte
	buf = make([]byte, 50)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
}

func WriteBytesToFile() {
	f, err := os.Open("teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("modifyingfile"))

	//checking if file was really modified
	var buf []byte
	buf = make([]byte, 60)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if result := strings.Contains(string(buf[:n]), "modifyingfile"); !result {
		fmt.Println("Problems")
		return
	}
	fmt.Println("File was modified")
}

func ReadContentFromFileAt() {
	f, err := os.Open("teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var buf []byte
	buf = make([]byte, 60)
	n, err := f.ReadAt(buf, 3) //Read after 3 byte offset
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("Content read: %s\n", string(buf[:n]))
}

func ReadContentFromFile() {
	f, err := os.Open("teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var buf []byte
	buf = make([]byte, 60)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes\n", n)
	fmt.Println("Contents: ", string(buf))
}

func WriteStringToFile() {
	f, err := os.Create("teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n, err := f.WriteString("golangfuckinrules")
	if err != nil {
		log.Fatalln("Problem writing to file: ", err)
	}
	fmt.Printf("Wrote %d bytes to file", n)
}

func GetWorkDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("problem getting current directory")
	}
	fmt.Println("Current directory is: ", dir)
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
