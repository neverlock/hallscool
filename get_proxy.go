package main

import (
	"fmt"
	"net"
	"time"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/neverlock/utility/random"
)

func main() {
	MAX := 10
	for key := 0; key <= MAX; key++ {
		time.Sleep(1 * time.Second)
		go req()
	}
}

func req() {
	urlStr := fmt.Sprintf("http://www.google.com")

	r := random.Int(0, len(proxy))
	proxyStr := proxy[r]
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	req, err := http.NewRequest("GET", urlStr, nil) // <-- URL-encoded payload
	if err != nil {
		defer func() {
			recover()
			log.Printf("Can't http.NewRequest\n")
		}()
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		defer func() {
			recover()
			log.Printf("Can't client.Do(req)\n")
		}()
		panic(err)
	}
	fmt.Printf("[%s] %s\n", proxyStr, resp.Status)
	defer resp.Body.Close()
	/*
	   body, err := ioutil.ReadAll(resp.Body)
	   if err != nil {
	           log.Fatal(err)
	   }
	   fmt.Println(string(body))
	*/

}
