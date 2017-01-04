package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库相关
const (
	username   = "root"
	password   = "root"
	ip         = "127.0.0.1"
	port       = "8889"
	db_name    = "jeason_daily"
	table_name = "mail_ss"
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

// 拿到 userEmail
func getUsersEmail() string {
	arrTag := make(map[string]string)
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+ip+":"+port+")/"+db_name+"?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	// 获取USERS表中的前十行记录
	rows, err := db.Query("SELECT * FROM mail_ss")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var id int
		var is_receive bool
		var name, userid, user_email string
		rows.Scan(&id, &name, &userid, &user_email, is_receive)
		fmt.Println("id:", id, "name:", name, "userid:", userid, "user_email:", user_email, "is_receive:", is_receive)

		tagItem := strconv.Itoa(i)

		arrTag[tagItem] = user_email
		i++
	}
	defer db.Close()

	fmt.Println(arrTag)
	resStr := ""

	for j := 0; j < len(arrTag); j++ {
		if j == len(arrTag)-1 {
			ii := strconv.Itoa(j)
			resStr = resStr + arrTag[ii]
		} else {
			ii := strconv.Itoa(j)
			resStr = resStr + arrTag[ii] + ";"
		}
	}

	return resStr
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

// 准备发射
func sendReady(innerHTML string) {
	user := "mailbyjeason@jeasonstudio.cn"
	password := "Admin12345"
	host := "smtp.exmail.qq.com:25"
	to := getUsersEmail()
	// fmt.Println(getUsersEmail())
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
	<header>` + myRes["0-1"] + `</header>
	</body>
	</html>`

	fmt.Println(myRes["0-1"])
	fmt.Println(tagHTML)

	// sendReady(tagHTML)
	// getUsersEmail()

	// ticker := time.NewTicker(time.Hour * 6)
	// for _ = range ticker.C {
	// 	// eMailTime := time.Now().Format("15:04")
	// 	sendReady(tagHTML)

	// }
}
