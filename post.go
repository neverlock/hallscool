package main

import ("net/http"
        "log"
        "net/url"
	"fmt"
        "math/rand"
	"strconv"
        )

func main()  {

	// Edit Post URL by view source from google form
        postUrl := "https://docs.google.com/forms/?/??????????????????????????????????/formResponse"
	// Edit Max your post
	Max := 10

        values := make(url.Values)

	for i:=0; i< Max ; i++{

		//add and Edit your entry.xxxx it is name of google form

	        values.Set("entry.1352624538", randomString(10))
	        values.Set("entry.668341357", randomString(10))
		values.Set("entry.1007776359",strconv.Itoa(rand.Intn(10000000000)))
	        // Submit form
	        resp, err := http.PostForm(postUrl, values)
	        if err != nil {
	                log.Fatal(err)
	        }
		fmt.Printf("[%d] %s\n",i,resp.Status)
	        defer resp.Body.Close()
	}

}

func randomString(l int) string{
	var char string
	buf := make([]byte, l)
	char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
        for i:=0 ; i<l ; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}
