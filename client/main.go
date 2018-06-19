package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
