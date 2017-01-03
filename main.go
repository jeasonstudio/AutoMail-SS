package main

import (
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
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
func analysisHTML() map[string]string {
	result := make(map[string]string)

	analys := getSS()
	str := strings.NewReader(analys)
	doc, _ := goquery.NewDocumentFromReader(str)

	res := doc.Find("#free .container .col-sm-4")

	for i := range res.Nodes {
		// fmt.Println(res.Eq(i))
		// res.Eq(i).Find("h4").Nodes[0].Text()
		item := res.Eq(i).Find("h4")

		for j := range item.Nodes {
			// fmt.Println(item.Eq(j).Text())
			thisItem := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			result[thisItem] = item.Eq(j).Text()
		}
	}
	return result
}

// 发送邮件
func SendToMAIL(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func sendReady(innerHTML string) {
	user := "mailbyjeason@jeasonstudio.cn"
	password := "Admin12345"
	host := "smtp.exmail.qq.com:25"
	to := "me@jeasonstudio.cn"

	subject := "Jeason Studio"

	body := innerHTML
	fmt.Println("send email")
	err := SendToMAIL(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}

func main() {
	myRes := analysisHTML()
	tagHTML := `<html>
    <head>
    <title>test</title>
    </head>
    <body>
    <header>aaa</header>
    </body>
    </html>`

	fmt.Println(myRes["0-1"])

	sendReady(tagHTML)

	// ticker := time.NewTicker(time.Hour * 6)
	// for _ = range ticker.C {
	// 	// eMailTime := time.Now().Format("15:04")
	// 	sendReady(tagHTML)

	// }
}
