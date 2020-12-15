package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/neverlock/utility/random"
)

type Js1 struct {
	Account   string `json:"account"`
	Birth     string `json:"birth"`
	ID_card   string `json:id_card"`
	Phone     string `json:phone"`
	User_name string `json:user_name"`
}

type Js2 struct {
	ID_card  string `json:id_card"`
	Password string `json:password"`
	SMS      string `json:sms"`
}

var rjs1 Js1

var rjs2 Js2

var userAgent, email, password, sms, fullname, ccno, expmonth, expyear, cvv, citizen, climit, phone, dobday, dobmonth, dobyear string

var rMonth []string = []string{
	"01",
	"02",
	"03",
	"04",
	"05",
	"06",
	"07",
	"08",
	"09",
	"10",
	"11",
	"12",
}

var rDay []string = []string{
	"01",
	"02",
	"03",
	"04",
	"05",
	"06",
	"07",
	"08",
	"09",
	"10",
	"11",
	"12",
	"13",
	"14",
	"15",
	"16",
	"17",
	"18",
	"19",
	"20",
	"21",
	"22",
	"23",
	"24",
	"25",
	"26",
	"27",
	"28",
}

var domain []string = []string{
	"gmail.com",
	"hotmail.com",
	"yahoo.com",
	"windowslive.com",
	"live.com",
}

var proxyType string

func initValue() {
	//var userAgent, email, password, fullname, ccno, expmonth, expyear, cvv, citizen, climit, phone, dobday, dobmonth, dobyear string
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"

	fn := randomString(10)
	ln := randomString(10)
	fullname = fmt.Sprintf("%s %s", fn, ln)

	citizen = fmt.Sprintf("%s", randomNString(13))

	atm := fmt.Sprintf("%s", randomNString(16))

	phone = fmt.Sprintf("08%s", randomNString(8))

	r5 := random.Int(0, len(rDay))
	dobday = fmt.Sprintf("%s", rDay[r5])
	r6 := random.Int(0, len(rMonth))
	dobmonth = fmt.Sprintf("%s", rMonth[r6])
	dobyear = fmt.Sprintf("%d", random.Int(2500, 2560))
	birth := fmt.Sprintf("%s-%s-%s", dobyear, dobmonth, dobday)

	password = fmt.Sprintf("%s", randomNString(6))

	sms = fmt.Sprintf("%s", randomNString(6))

	rjs1.Account = atm
	rjs1.Birth = birth
	rjs1.ID_card = citizen
	rjs1.Phone = phone
	rjs1.User_name = fullname

	rjs2.ID_card = citizen
	rjs2.Password = password
	rjs2.SMS = sms

}

func main() {

	proxyType = "socks" //http or socks

	MAX := 100000
	for key := 0; key <= MAX; key++ {
		//		time.Sleep(1 * time.Second)
		req()
	}
}

func req() {
	initValue()

	urlStr := fmt.Sprintf("https://www.scbstar.com/page/i.php")

	if len(proxy) == 0 {
		fmt.Println("Exit program no proxy left")
		os.Exit(1)
	}
	r := random.Int(0, len(proxy))
	proxyStr := proxy[r]
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}

	if proxyType == "socks" {
		dialSocksProxy, err := proxy.SOCKS5("tcp", "proxy_ip", nil, proxy.Direct)
		if err != nil {
			fmt.Println("Error connecting to proxy:", err)
		}
		transport := &http.Transport{Dial: dialSocksProxy.Dial}

	} else {
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		}
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(rjs1)

	req, err := http.NewRequest("POST", urlStr, payloadBuf) // <-- URL-encoded payload
	if err != nil {
		defer func() {
			recover()
			log.Printf("[❌] Can't http.NewRequest[proxy = %s]\n", proxyStr)
			log.Printf("[‼️ ] Remove =%s from slice\n", proxyStr)
			//remove element from slice
			proxy[r] = proxy[len(proxy)-1]
			proxy[len(proxy)-1] = ""
			proxy = proxy[:len(proxy)-1]
			//end remove element from slice
		}()
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		defer func() {
			recover()
			log.Printf("[❌] Can't client.Do(req)[proxy = %s]\n", proxyStr)
			log.Printf("[‼️ ] Remove =%s from slice\n", proxyStr)
			//remove element from slice
			proxy[r] = proxy[len(proxy)-1]
			proxy[len(proxy)-1] = ""
			proxy = proxy[:len(proxy)-1]
			//end remove element from slice
		}()
		panic(err)
	}
	fmt.Printf("[first page post][proxy = %s] %s\n", proxyStr, resp.Status)
	defer resp.Body.Close()

	req1(proxyStr, r)
	/*
	   body, err := ioutil.ReadAll(resp.Body)
	   if err != nil {
	           log.Fatal(err)
	   }
	   fmt.Println(string(body))
	*/

}

func req1(proxyStr string, proxyIndex int) {
	urlStr := fmt.Sprintf("https://www.scbstar.com/page/i3.php")
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
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(rjs2)

	req, err := http.NewRequest("POST", urlStr, payloadBuf) // <-- URL-encoded payload
	if err != nil {
		defer func() {
			recover()
			log.Printf("[❌] Can't http.NewRequest\n")
			log.Printf("[‼️ ] Remove =%s from slice\n", proxyStr)
			//remove element from slice
			proxy[proxyIndex] = proxy[len(proxy)-1]
			proxy[len(proxy)-1] = ""
			proxy = proxy[:len(proxy)-1]
			//end remove element from slice
		}()
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		defer func() {
			recover()
			//			log.Printf("[❌] Can't client.Do(req)\n")
			log.Printf("[❌] Can't client.Do(req)[proxy = %s]\n", proxyStr)
			log.Printf("[‼️ ] Remove =%s from slice\n", proxyStr)
			//remove element from slice
			proxy[proxyIndex] = proxy[len(proxy)-1]
			proxy[len(proxy)-1] = ""
			proxy = proxy[:len(proxy)-1]
			//end remove element from slice
		}()
		panic(err)
	}
	fmt.Printf("[Page2][proxy = %s] %s\n", proxyStr, resp.Status)
	defer resp.Body.Close()

	//req2(proxyStr)
}

func req2(proxyStr string) {
	urlStr := fmt.Sprintf("https://ipay.bualuang.com.checkinhairh.com/assets/includes/get_step1.php")
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
	//var userAgent, email, password, fullname, ccno, expmonth, expyear, cvv, citizen, climit, phone, dobday, dobmonth, dobyear string
	query := fmt.Sprintf("fullname=%s&masterMerId=1&ccno=%s&expmonth=%s&expyear=%s&cvv=%s&submit=ต่อ&address2=&city=&state=&post=&citizen=%s&climit=%s&phone=%s&dobday=%s&dobmonth=%s&dobyear=%s", fullname, ccno, expmonth, expyear, cvv, citizen, climit, phone, dobday, dobmonth, dobyear)
	//fmt.Println(query)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(query)) // <-- URL-encoded payload
	if err != nil {
		defer func() {
			recover()
			log.Printf("[❌] Can't http.NewRequest\n")
		}()
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		defer func() {
			recover()
			log.Printf("[❌] Can't client.Do(req)\n")
		}()
		panic(err)
	}
	fmt.Printf("[P2][%s] %s\n", proxyStr, resp.Status)
	defer resp.Body.Close()

	req2(proxyStr)
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

func randomNString(l int) string {
	var char string
	buf := make([]byte, l)
	char = "0123456789"
	for i := 0; i < l; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
