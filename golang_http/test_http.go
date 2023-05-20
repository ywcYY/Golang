package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func testGet() {
	// https://www.juhe.cn/box/index/id/73
	url := "http://apis.juhe.cn/simpleWeather/query?key=087d7d10f700d20e27bb753cd806e40b&city=北京"
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	fmt.Printf("b: %v\n", string(b))
}
