package main

import ("net/http"
        "fmt"
        "strings"
	"math/rand"
        )

func main()  {

	/**This is example var**/
	/* please edit */
	user := randomString(10)
	domain := randomString(10)
	email := fmt.Sprintf("%s@%s.com",user,domain)
	pass := randomString(10)

	/** This is example Url**/
	/* please edit */
	postUrl := "http://www.your_domain.com/post.php"
	/** This is example query post value**/
	/* please edit */
	query := fmt.Sprintf("lsd=AVrmP5gx&legacy_return=1&trynum=1&lgnrnd=082906_Ew7t&lgnjs=n&email=%s&pass=%s&default_persistent=0&login=เข้าสู่ระบบ",email,pass)

	req, err := http.NewRequest("POST", postUrl, strings.NewReader(query))
	/** This is example url referer**/
	/* please edit */
	req.Header.Set("Referer", "http://www.your_domain.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    fmt.Println("response Bodys:",resp.Body)

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
