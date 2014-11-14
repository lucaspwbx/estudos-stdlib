package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
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

	//looking ip for given hostname - returns slice of strings
	addrs, err := net.LookupHost("www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range addrs {
		fmt.Println(v)
	}

	fmt.Println()

	//looking IP for given hostname - returns slice of IPs
	addrs2, err := net.LookupIP("www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range addrs2 {
		fmt.Println(v)
	}

	//looking up port of service
	port, err := net.LookupPort("tcp", "ftp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Port: ", port)

	//package Strings

	//Fields - splits on whitespace
	fmt.Printf("Fiels are %q", strings.Fields(" foo bar baz "))

	fmt.Println()

	//FieldsFunc - splits based on a func - not a letter neither a number
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f)) //["foo1", "bar2", "baz3"]

	fmt.Println()

	//Contains
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false

	//ContainsAny
	fmt.Println(strings.ContainsAny("team", "i")) //false

	//Count
	fmt.Println(strings.Count("cheese", "e")) //3

	//Replace
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo

	//Reader
	reader := strings.NewReader("lucas")

	//Len
	fmt.Println(reader.Len()) //5

	//Read
	var buf []byte
	buf = make([]byte, reader.Len())
	number, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:number])) //lucas

	//ReadAt - read after third index, including it
	reader2 := strings.NewReader("lucas")
	var buf2 []byte
	buf2 = make([]byte, reader2.Len())
	number2, err := reader2.ReadAt(buf2, 3)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf2[:number2])) //as

	reader3 := strings.NewReader("lucas")
	var buf3 []byte
	buf3 = make([]byte, reader3.Len())
	_, err = reader3.ReadAt(buf3, -1)
	if err != nil && err != io.EOF {
		fmt.Println(err) // negative offset
	}

	_, err = reader3.ReadAt(buf3, 10)
	if err != nil {
		fmt.Println(err) // EOF - offset >= length of string
	}

	b, err := reader3.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b)) // l

	err = reader3.UnreadByte() //go back to "l"
	if err != nil {
		fmt.Println("Error unreading byte")
	}

	r, _, err = reader3.ReadRune()
	if err != nil {
		fmt.Println("Error reading rune")
	}
	fmt.Println(string(r)) //l

	err = reader3.UnreadRune() //go back to "l"
	if err != nil {
		fmt.Println("Error unreading rune")
		return
	}
	fmt.Println("ok, unread rune")

	//Seek
	_, err = reader3.Seek(1, 0) //0 means relative to beginning of string
	if err != nil {
		fmt.Println(err)
	}
	omik, _ := reader3.ReadByte()
	fmt.Println(string(omik))

	_, err = reader3.Seek(2, 1) //1 means relative to current offset
	if err != nil {
		fmt.Println(err)
	}
	omik2, _ := reader3.ReadByte()
	fmt.Println(string(omik2))

	reader4 := strings.NewReader("omik")
	_, err = reader4.Seek(-2, 2) //2 means relative to end of string
	if err != nil {
		fmt.Println(err)
	}
	omik3, err := reader4.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(omik3))

	//package sync

	//waitgroup
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <= 5; i++ {
			fmt.Println("Goroutine 1: ", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <= 5; i++ {
			fmt.Println("Goroutine 2: ", i)
		}
	}()
	wg.Wait()
	fmt.Println("goroutines processed")

	test := &TestRWMutex{critical: "blalal"}
	var wg2 sync.WaitGroup
	wg2.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go func(i int) {
			defer wg2.Done()
			bla := fmt.Sprintf("goroutine %d modificando secao critica\n", i)
			test.Write(bla)
		}(i)
	}
	wg2.Wait()
	fmt.Println(test.critical)

	fmt.Println()

	var mutex sync.Mutex
	var criticalSection string
	go func() {
		mutex.Lock()
		time.Sleep(5000)
		criticalSection = "lucas"
		mutex.Unlock()
	}()
	go func() {
		mutex.Lock()
		criticalSection = "weiblen"
		mutex.Unlock()
	}()
	time.Sleep(7000)
	fmt.Println(criticalSection)

	//package strconv

	//parse string to bool
	value2, err := strconv.ParseBool("TRUE")
	if err != nil {
		fmt.Println("Error converting")
	}
	fmt.Println(value2) // true - bool

	//receives a bool and returns a string
	backToString := strconv.FormatBool(value2)
	fmt.Println(backToString) //true - string

	fmt.Println("-----")

	//append bool converted to string to slice
	var sliceBool []byte
	sliceBool = strconv.AppendBool(sliceBool, false)
	fmt.Println(string(sliceBool)) //false

	float, err := strconv.ParseFloat("1.39", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(float) // 1.39

	integer, _ := strconv.ParseInt("190", 10, 64) //base 10, int64
	fmt.Println(integer)                          // 190

	//Atoi - shorthand for ParseInt(s, 10, 0)
	integer2, _ := strconv.Atoi("192")
	fmt.Println(integer2) // 192

	//package http
	//client

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Last-Modified", "sometime")
		fmt.Fprintf(w, "User-agent: go\nDisallow: /something/")
	})

	//testing client
	ts := httptest.NewServer(handler)

	req, err := http.Get(ts.URL)
	var readNr []byte
	if err != nil {
		fmt.Println(err)
		return
	}
	readNr, err = ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(readNr)) //User-agent: go\nDisallow: /something/
	ts.Close()

	//testing head
	ts = httptest.NewServer(handler)
	req, err = http.Head(ts.URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(req.Header["Last-Modified"]) //["sometime"]

	transport := &recordingTransport{}
	client := &http.Client{Transport: transport}
	url2 := "http://dummy.faketld/"
	client.Get(url2)                        //does not hit network
	fmt.Println(transport.req.Method)       //GET
	fmt.Println(transport.req.URL.String()) //dummy.fakeltd
	fmt.Println(transport.req.Header)       //map[]

	json := `{"key":"value"}`
	body := strings.NewReader(json)
	client.Post(url2, "application/json", body)
	fmt.Println(transport.req.Method)        //POST
	fmt.Println(transport.req.URL.String())  //dummy.fakeltd
	fmt.Println(transport.req.Header)        //map[]
	fmt.Println(transport.req.Close)         //true
	fmt.Println(transport.req.ContentLength) //14

	fmt.Println()

	form := url.Values{}
	form.Set("foo", "bar")
	form.Add("foo", "bar2")
	form.Set("bar", "baz")
	client.PostForm(url2, form)
	fmt.Println(transport.req.Method) //POST
	fmt.Println(transport.req.URL.String())
	fmt.Println(transport.req.Header.Get("Content-Type")) //application/x-www-form-urlencoded
	fmt.Println(transport.req.Close)                      // false
	fmt.Println(transport.req.ContentLength)              //24
	bodyR, _ := ioutil.ReadAll(transport.req.Body)
	fmt.Println(string(bodyR)) //bar=baz&foo=bar&foo=bar2

	//testing redirects
	var ts2 *httptest.Server
	ts2 = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.Atoi(r.FormValue("n"))
		if n == 7 {
			fmt.Println(r.Referer())
		}
		if n < 15 {
			http.Redirect(w, r, fmt.Sprintf("/?n=%d", n+1), http.StatusFound)
			return
		}
		fmt.Fprintf(w, "n=%d", n)
	}))
	defer ts2.Close()

	c2 := &http.Client{}
	_, err = c2.Get(ts2.URL)
	fmt.Println(err)

	_, err = c2.Head(ts2.URL)
	fmt.Println(err)

	greq, _ := http.NewRequest("GET", ts2.URL, nil)
	_, err = c2.Do(greq)
	fmt.Println(err)

	var checkErr error
	var lastVia []*http.Request
	c2 = &http.Client{CheckRedirect: func(_ *http.Request, via []*http.Request) error {
		lastVia = via
		return checkErr
	}}
	response, _ := c2.Get(ts2.URL)
	response.Body.Close()
	fmt.Println(response.Request.URL.String())
	fmt.Println(len(lastVia))

	//strange behavior
	response, err = c2.Get(ts2.URL)
	fmt.Println(err)
	fmt.Println(response)
	fmt.Println(response.Header.Get("Location"))

	//redirects via POST
	PostRedirects()

	ClientSendsCookieFromJar()

	RedirectCookiesJar()
}

var expectedCookies = []*http.Cookie{
	{Name: "ChocolateChip", Value: "tasty"},
	{Name: "First", Value: "Hit"},
	{Name: "Second", Value: "Hit"},
}

var echoCookiesRedirectHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		http.SetCookie(w, cookie)
	}
	if r.URL.Path == "/" {
		http.SetCookie(w, expectedCookies[1])
		http.Redirect(w, r, "/second", http.StatusMovedPermanently)
	} else {
		http.SetCookie(w, expectedCookies[2])
		w.Write([]byte("hello"))
	}
})

type TestJar struct {
	m      sync.Mutex
	perURL map[string][]*http.Cookie
}

func (j *TestJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.m.Lock()
	defer j.m.Unlock()
	if j.perURL == nil {
		j.perURL = make(map[string][]*http.Cookie)
	}
	j.perURL[u.Host] = cookies
}

func (j *TestJar) Cookies(u *url.URL) []*http.Cookie {
	j.m.Lock()
	defer j.m.Unlock()
	return j.perURL[u.Host]
}

func ClientSendsCookieFromJar() {
	tr := &recordingTransport{}
	client := &http.Client{Transport: tr}
	client.Jar = &TestJar{perURL: make(map[string][]*http.Cookie)}
	us := "http://dummy.fakelt/"
	u, _ := url.Parse(us)
	client.Jar.SetCookies(u, expectedCookies)
	client.Get(us)
	fmt.Println(tr.req.Cookies())

	client.Head(us)
	fmt.Println(tr.req.Cookies())

	client.Post(us, "text/plain", strings.NewReader("body"))
	fmt.Println(tr.req.Cookies())

	client.PostForm(us, url.Values{})
	fmt.Println(tr.req.Cookies())

	req, _ := http.NewRequest("GET", us, nil)
	client.Do(req)
	fmt.Println(tr.req.Cookies())

	req, _ = http.NewRequest("POST", us, nil)
	client.Do(req)
	fmt.Println(tr.req.Cookies())
}

func RedirectCookiesJar() {
	var ts *httptest.Server
	ts = httptest.NewServer(echoCookiesRedirectHandler)
	defer ts.Close()
	c := &http.Client{
		Jar: new(TestJar),
	}
	u, _ := url.Parse(ts.URL)
	c.Jar.SetCookies(u, []*http.Cookie{expectedCookies[0]})
	resp, _ := c.Get(ts.URL)
	resp.Body.Close()
	fmt.Println(resp.Cookies())
}

func PostRedirects() {
	var log struct {
		sync.Mutex
		bytes.Buffer
	}

	var ts *httptest.Server
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Lock()
		fmt.Fprintf(&log.Buffer, "%s %s ", r.Method, r.RequestURI)
		log.Unlock()
		if v := r.URL.Query().Get("code"); v != "" {
			code, _ := strconv.Atoi(v)
			if code/100 == 3 {
				w.Header().Set("Location", ts.URL)
			}
			w.WriteHeader(code)
		}
	}))
	defer ts.Close()
	res, _ := http.Post(ts.URL+"/", "text/plain", strings.NewReader("Some content"))
	fmt.Println(res.StatusCode)

	res, _ = http.Post(ts.URL+"/?code=301", "text/plain", strings.NewReader("Some content"))
	fmt.Println(res.StatusCode)

	log.Lock()
	got := log.String()
	log.Unlock()
	fmt.Println(got)
}

type recordingTransport struct {
	req *http.Request
}

func (t *recordingTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t.req = req
	return nil, errors.New("dummy impl")
}

type TestRWMutex struct {
	critical string
	rw       sync.RWMutex
}

func (t *TestRWMutex) Write(word string) error {
	defer t.rw.Unlock()
	t.rw.Lock()
	t.critical = word
	return nil
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
