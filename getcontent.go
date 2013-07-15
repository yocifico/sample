package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getContent(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func main() {
	flag.Parse()
	url := flag.Arg(0)
	content, err := getContent(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
