package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"container/heap"
	"container/list"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

type File struct {
	Name, Body string
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func intHeapExample() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func priorityQueueExample() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)

	}
}

func listExample() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func zipExample() {
	//writing
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)
	file := File{"readme.txt", "This archive contains some text files"}
	f, err := w.Create(file.Name)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(file.Body))
	if err != nil {
		log.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	//reading
}

func tarExample() {
	//writing tar archive
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	file := File{"readme.txt", "This archive contains some text files"}

	hdr := &tar.Header{
		Name: file.Name,
		Size: int64(len(file.Body)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		log.Fatalln(err)
	}
	if _, err := tw.Write([]byte(file.Body)); err != nil {
		log.Fatalln(err)
	}
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	//open the tar archive
	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}
}

func main() {
	intHeapExample()
	fmt.Println()
	priorityQueueExample()
	fmt.Println()
	listExample()

	//missing ring example

	cryptoExample()
	encoding()
}

func cryptoExample() {
	md5Example()
	sha1Example()
	randExample()
}

func encoding() {
	//base64()
	binary()
}

func base64() {
	//input := []byte("foobar")
	//encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	//encoder.Write(input)
}

func md5Example() {
	//example using new
	h := md5.New()
	io.WriteString(h, "The fog is getting ticker!")
	io.WriteString(h, "And leon blalalal")
	fmt.Printf("%x", h.Sum(nil))

	fmt.Println()

	//example using sum
	data := []byte("These pretzels")
	fmt.Printf("%x", md5.Sum(data))
}

func sha1Example() {
	h := sha1.New()
	io.WriteString(h, "Teste")
	io.WriteString(h, "ok")
	fmt.Printf("%x", h.Sum(nil))

	fmt.Println()

	data := []byte("ok blalalala")
	fmt.Printf("%x", sha1.Sum(data))

	fmt.Println()
}

func randExample() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(b)
}
