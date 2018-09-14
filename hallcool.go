package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	var MAX, i int
	MAX = 1000000

	red := color.New(color.FgRed).SprintFunc()

	//fmt.Println("************************************************")
	fmt.Printf("%s\n", red("************************************************"))
	fmt.Printf("%s I will give you free %s register user %s\n", red("*"), red(strconv.Itoa(MAX)), red("*"))
	fmt.Printf("%s\n", red("************************************************"))
	//fmt.Println("************************************************")
	for i = 1; i <= MAX; i++ {
		/**This is example var**/
		/* please edit */
		red := color.New(color.FgRed).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		user := randomString(10)
		domain := randomString(10)
		email := fmt.Sprintf("%s@%s.com", user, domain)
		line := randomString(10)
		tel := randInt(10000000, 99999999)

		/** This is example Url**/
		/* please edit */
		postUrl := "http://www.sawasdeethailand.net/gopage14.html"
		/** This is example query post value**/
		/* please edit */
		//query := fmt.Sprintf("lsd=AVrmP5gx&legacy_return=1&trynum=1&lgnrnd=082906_Ew7t&lgnjs=n&email=%s&pass=%s&default_persistent=0&login=เข้าสู่ระบบ", email, pass)
		query := fmt.Sprintf("email=%s&line_id=%s&message=.&var_fullname=%s&รู้จักสวัสดีไทยแลนด์จาก=Google&อื่นๆ_โปรดระบุ=.&เบอร์โทรติดต่อกลับ=08%d", email, line, user, tel)

		fmt.Printf("[%s / %s] Req to %s", yellow(strconv.Itoa(i)), red(strconv.Itoa(MAX)), postUrl)

		req, err := http.NewRequest("POST", postUrl, strings.NewReader(query))
		/** This is example url referer**/
		/* please edit */
		req.Header.Set("origin", "http://www.sawasdeethailand.net")
		req.Header.Set("upgrade-insecure-requests", "1")
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
		req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		req.Header.Set("referer", "http://www.sawasdeethailand.net/pro=")
		req.Header.Set("accept-encoding", "gzip, deflate")
		req.Header.Set("accept-language", "en-US,en;q=0.9,th;q=0.8")
		req.Header.Set("cache-control", "no-cache")
		start := time.Now()
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			defer func() {
				recover()
				fmt.Println(err)
			}()
			panic(err)
		} else {
			defer resp.Body.Close()
		}

		fmt.Printf("  [response Status: %s] [%s]\n", green(resp.Status), time.Since(start))
		//	fmt.Println("response Headers:", resp.Header)
		//	fmt.Println("response Bodys:", resp.Body)
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
