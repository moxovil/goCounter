package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var urls []string
	/*
		Заполняем массив urls[] пока не придет пустая строка
		(два переноса строки подряд)
	*/
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		url := sc.Text()
		if len(url) > 0 {
			urls = append(urls, url)
		} else {
			break
		}
	}

	ch1 := make(chan int) //count для текущего url
	for _, url := range urls {
		go func(uri string) {
			r, err := http.Get(uri)
			if err != nil {
				log.Println(err)
				ch1 <- 0
				fmt.Printf("Count for %s: %d\n", uri, 0)
				return
			}

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				ch1 <- 0
				fmt.Printf("Count for %s: %d\n", uri, 0)
				return
			}

			cnt := bytes.Count(body, []byte("Go"))
			// if cnt == 0 {
			// 	fmt.Printf("NotFound\t uri: %s\nbody: %s\n", uri, string(body))
			// }
			ch1 <- cnt
			fmt.Printf("Count for %s: %d\n", uri, cnt)
		}(url)
		time.Sleep(50 * time.Millisecond) //Без задержки сервер выдаcт Too Many Requests
	}

	ch2 := make(chan int) //count total
	go func() {
		var total int = 0
		for i := 0; i < len(urls); i++ {
			dt := <-ch1
			total += dt
		}
		ch2 <- total
	}()
	fmt.Printf("Total: %d\n", <-ch2)
}
