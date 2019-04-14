package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/neverlock/utility/random"
)

var userAgent, email, password, fullname, ccno, expmonth, expyear, cvv, citizen, climit, phone, dobday, dobmonth, dobyear string
var cLimit []string = []string{
	"50000",
	"52000",
	"55000",
	"60000",
	"62000",
	"65000",
	"70000",
	"72000",
	"75000",
	"80000",
	"82000",
	"85000",
	"90000",
	"92000",
	"95000",
	"100000",
}
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

var rXpYear []string = []string{
	"2020",
	"2021",
	"2022",
	"2023",
}

var domain []string = []string{
	"gmail.com",
	"hotmail.com",
	"yahoo.com",
	"windowslive.com",
	"live.com",
}

func initValue() {
	//var userAgent, email, password, fullname, ccno, expmonth, expyear, cvv, citizen, climit, phone, dobday, dobmonth, dobyear string
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
	user := randomString(10)
	r1 := random.Int(0, len(domain))
	email = fmt.Sprintf("%s@%s", user, domain[r1])
	password = randomString(10)
	fn := randomString(5)
	ln := randomString(5)
	fullname = fmt.Sprintf("%s %s", fn, ln)
	fullname = FULL
	ccno = randomNString(16)
	r2 := random.Int(0, len(rMonth))
	expmonth = fmt.Sprintf("%s", rMonth[r2])
	r3 := random.Int(0, len(rXpYear))
	expyear = fmt.Sprintf("%s", rXpYear[r3])
	cvv = randomNString(3)
	citizen = fmt.Sprintf("%s", randomNString(13))
	r4 := random.Int(0, len(cLimit))
	climit = fmt.Sprintf("%s", cLimit[r4])
	phone = fmt.Sprintf("08%s", randomNString(8))
	r5 := random.Int(0, len(rDay))
	dobday = fmt.Sprintf("%s", rDay[r5])
	r6 := random.Int(0, len(rMonth))
	dobmonth = fmt.Sprintf("%s", rMonth[r6])
	dobyear = fmt.Sprintf("%d", random.Int(1969, 2000))

}

func main() {

	MAX := 100000
	for key := 0; key <= MAX; key++ {
		time.Sleep(1 * time.Second)
		go req()
	}
}

func req() {
	initValue()
	/*
		fmt.Println(email)
		fmt.Println(password)
		fmt.Println(fullname)
		fmt.Println(ccno)
		fmt.Println(expmonth)
		fmt.Println(expyear)
		fmt.Println(cvv)
		fmt.Println(citizen)
		fmt.Println(climit)
		fmt.Println(phone)
		fmt.Println(dobday)
		fmt.Println(dobmonth)
		fmt.Println(dobyear)
	*/
	urlStr := fmt.Sprintf("https://ipay.bualuang.com.checkinhairh.com/locked.php")

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
	query := fmt.Sprintf("masterMerId=1&email=%s&password=%s&submit=เข้าสู่ระบบ", email, password)
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
	fmt.Printf("[login][%s] %s\n", proxyStr, resp.Status)
	defer resp.Body.Close()

	req1(proxyStr)
	/*
	   body, err := ioutil.ReadAll(resp.Body)
	   if err != nil {
	           log.Fatal(err)
	   }
	   fmt.Println(string(body))
	*/

}

func req1(proxyStr string) {
	urlStr := fmt.Sprintf("https://ipay.bualuang.com.checkinhairh.com/verify.php")
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
	query := fmt.Sprintf("fullname=%s&masterMerId=1&ccno=%s&expmonth=%s&expyear=%s&cvv=%s&submit=ต่อ", fullname, ccno, expmonth, expyear, cvv)
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
	fmt.Printf("[P1][%s] %s\n", proxyStr, resp.Status)
	defer resp.Body.Close()

	req2(proxyStr)
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
