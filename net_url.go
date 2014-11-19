package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://www.google.com/search?q=teste") //parse raw url into a URL structure
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"  //changing scheme
	u.Host = "bing.com" //changing host
	q := u.Query()
	q.Set("q", "golang")    //change key to query string
	q.Set("t", "teste")     //add new key to query string
	u.RawQuery = q.Encode() //encode
	fmt.Println(u)          //https://www.bing.com/search?q=golang&t=teste

	v := url.Values{}
	v.Set("name", "Xiao")
	v.Add("friend", "Dek")
	fmt.Println(v.Get("name")) //Xiao
	//get the first value associate with the given key
	fmt.Println(v.Get("friend")) //Dek
	//acessing all values
	fmt.Println(v["friend"]) // [dek]
	v.Set("blabla", "crazy")
	fmt.Println(v.Get("blabla")) //crazy
	v.Del("blabla")
	fmt.Println(v.Get("blabla")) // ""

	fmt.Println(v.Encode())

	user := url.User("xiaorules") //returns a userinfo containing username xiaorules and no password
	fmt.Println(user)

	user2 := url.UserPassword("xiaorules2", "senha")
	fmt.Println(user2.Username())
	pass, _ := user2.Password()
	fmt.Println(pass)

	//TODO
	//check implementation of the Stringer interface for URL to look at a good use for bytes.Buffer

	//parse query string -> return url.Values
	m, _ := url.ParseQuery("q=teste&b=ok")
	fmt.Println(m)

	// encode values to string
	bla := m.Encode()
	fmt.Println(bla)

	//check if url is absolute
	fmt.Println(u.IsAbs()) //tru

	temp, _ := url.Parse("http://www.terra.com.br")
	fmt.Println(temp)

	fmt.Println(u.RequestURI()) //returns requesturi

	values := u.Query()
	fmt.Println(values)
}
