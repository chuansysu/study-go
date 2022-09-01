package main

import (
	"fmt"
	"net/url"
)

func main() {
	/*
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Println("get failed:", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read from resp.Body failed:", err)
			return
		}
		fmt.Println(string(body))
	*/
	//apiUrl := "http://127.0.0.1:9090/get"

	data := url.Values{}
	data.Set("name", "xiaomi")
	data.Set("age", "18")
	data.Add("age", "20")
	fmt.Printf("%T,%v\n", data, data)
	fmt.Println(data)
	fmt.Println(data["name"])
}
