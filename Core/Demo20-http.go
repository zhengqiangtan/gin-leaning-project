package main

import (
	"fmt"
	"net/http"
	"time"
)

func call(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("user-agent", "golang application")
	req.Header.Add("foo", "bar1")
	req.Header.Add("foo", "bar2")
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}

func main() {
	call("https://google.com", "GET")
}
