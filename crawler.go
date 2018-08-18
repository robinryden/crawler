package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (j Job) Crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	// Read content of response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)
}
