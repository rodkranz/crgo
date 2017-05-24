// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package crgo

import (
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Request(url string, chn chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	secs := time.Since(start).Seconds()

	if err != nil {
		chn <- fmt.Sprintf("[%.2fs] Error: %s", secs, err.Error())
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	chn <- fmt.Sprintf("[%.2fs] elapsed time for request [%s] with [%d] ", secs, url, len(body))
}

func Run(urls []string) error {
	if len(urls) == 0 {
		return fmt.Errorf("You need to expesify the url one or more.")
	}

	start := time.Now()
	chn := make(chan string)

	for _, url := range urls {
		go Request(url, chn)
	}

	for range urls {
		fmt.Println(<-chn)
	}

	fmt.Printf("[%.2fs] elapsed time.\n", time.Since(start).Seconds())
	return nil
}
