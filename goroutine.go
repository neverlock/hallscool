package main

import (
	"fmt"
//	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
//	"os"
	"strconv"
)

var postUrl string = "http://morning-news.bectero.com.ppc.ac.th/save.php"

func main() {

	// Edit Post URL by view source from google form
	//	postUrl := "https://docs.google.com/forms/d/1QmlmgXokVvkH94szSJhU_Gs3KkRuyVYgi6nMjBsNgv0/formResponse"

	go postThem()
	var input string
	fmt.Scanln(&input)

}

func postThem() {
//	Max := 10

	values := make(url.Values)

//	for i := 0; i < Max; i++ {
	for i := 1 ;; i++ {

		//add and Edit your entry.xxxx it is name of google form

		values.Set("user", strconv.Itoa(rand.Intn(100)))
		values.Set("pass", randomString(20))
		// Submit form
		resp, err := http.PostForm(postUrl, values)
		if err != nil {
			log.Fatal(err)
		}

/*
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("[%d] %s\n", i, contents)
*/

		fmt.Printf("[%d] %s\n", i, resp.Status)
		defer resp.Body.Close()
	}
}

func randomString(l int) string {
	var char string
	buf := make([]byte, l)
	char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < l; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
