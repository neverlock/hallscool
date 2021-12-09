package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
	//"io/ioutil"
	"io"
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
	//'type=Garena&taikhoan=0812345678&matkhau=1234
	//var userAgent, email, password, fullname, ccno, expmonth, expyear, cvv, citizen, climit, phone, dobday, dobmonth, dobyear string
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
	password = randomString(10)
	phone = fmt.Sprintf("08%s", randomNString(8))

}

func main() {

	MAX := 100000
	//MAX := 1
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
	urlStr := fmt.Sprintf("https://rov-membership.com/assets/ajax/tocchien.php")

	r := random.Int(0, len(listProxy))
	proxyStr := listProxy[r]
	//proxyURL, err := url.Parse(proxyStr)
	proxyURL, err := url.Parse("http://" + proxyStr)
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
	query := fmt.Sprintf("type=Garena&taikhoan=%s&matkhau=%s", phone, password)
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
			log.Printf("[❌] Can't client.Do(req)[proxy = %s]\n",proxyStr)
			/*
			log.Printf("[‼️ ] Remove =%s from slice\n", proxyStr)
			listProxy[r] = listProxy[len(listProxy)-1]
			listProxy[len(listProxy)-1] = ""
			listProxy = listProxy[:len(listProxy)-1]
			*/
		}()
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("[✅][response][%s] %s\n", proxyStr, resp.Status)
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	} else {
		fmt.Printf("[❗️][response][%s] %s\n", proxyStr, resp.Status)
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
