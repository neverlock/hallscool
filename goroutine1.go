package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	)
var index int
func main() {
index = 1
for {
    nbConcurrentGet := 100
    urls :=  make(chan string, nbConcurrentGet)
    var wg sync.WaitGroup
    for i := 0; i < nbConcurrentGet; i++ {
        go func (){
            for url1 := range urls {
                postThem(url1)
                wg.Done()
            }
        }()
    }
    for i:=0; i<1040; i++ {
        wg.Add(1)
        urls <- fmt.Sprintf("http://morning-news.bectero.com.ppc.ac.th/save.php")
    }
    wg.Wait()
    fmt.Println("Finished")
}
}

func postThem(url1 string) {

		values := make(url.Values)

		values.Set("user", strconv.Itoa(rand.Intn(100)))
		values.Set("pass", randomString(20))
		resp, err := http.PostForm(url1, values)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[%d] %s\n",index,  resp.Status)
		index++
		defer resp.Body.Close()


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
