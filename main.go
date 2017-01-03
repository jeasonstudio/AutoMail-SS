package main

import (
	"io/ioutil"
	"net/http"
)

// 获取带 ss 账号的 html 字符串
func getSS() string {
	tagLoginURL := "http://www.ishadowsocks.me/"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", tagLoginURL, nil)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	// fmt.Println(string(data), err)
	return string(data)
}

// 解析 html 返回 map
func main() {
	getSS()
}
