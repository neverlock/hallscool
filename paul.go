package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	//"strconv"
)

func main() {

	// Edit Post URL by view source from google form
	postUrl := "https://pala42.com/ee/post.php"
	// Edit Max your post
	Max := 100000

	values := make(url.Values)

	for i := 0; i < Max; i++ {

		//add and Edit your entry.xxxx it is name of google form
		user := randomString(10)
		//domain := randomString(10)
		//email := fmt.Sprintf("%s@%s.com", user, domain)
		email := fmt.Sprintf("%s@gmail.com", user)
		pass := randomString(10)

		//input text use randomString
		/*
		values.Set("lsd", "AVqkYwEU")
		values.Set("charset_test", "%E2%82%AC%2C%C2%B4%2C%E2%82%AC%2C%C2%B4%2C%E6%B0%B4%2C%D0%94%2C%D0%84")
		values.Set("version", "1")
		values.Set("ajax", "0")
		values.Set("width", "0")
		values.Set("pxr", "0")
		values.Set("gps", "0")
		values.Set("dimensions", "0")
		values.Set("m_ts", "1444981372")
		values.Set("li", "fKogVnFYklxjk1zuKIBvk-jr")
		values.Set("login", "%C4%90%C4%83ng+nh%E1%BA%ADp")
		*/

		values.Set("email", email)
		values.Set("pass", pass)
		//input digit like Tel No. use strconv.Itoa(rand.Intn(10000000000))
		//values.Set("entry.1007776359", strconv.Itoa(rand.Intn(10000000000)))
		// Submit form
		resp, err := http.PostForm(postUrl, values)
		if err != nil {
			log.Fatal(err)
		}
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
