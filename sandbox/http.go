package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

func main() {
	org := "feedforce"
	res, err := http.Get("https://api.github.com/orgs/" + org + "/members")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	fmt.Println(res.Status)
}
