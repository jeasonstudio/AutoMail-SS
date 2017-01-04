# AutoMail-SS
Spy SS from some website And mail it to me
自动爬取Shadowsocks账号密码并发送邮件给指定名单。

## 功能

 - 自动从免费ShadowSocks账号提供站点：[http://ishadowsocks.me/](http://ishadowsocks.me/) 爬取ss账号，每六个小时更新一次。
 - 自动从数据库读取Email地址，并在每天相应节点发送账号密码。

## TODO

 - 添加自动退订接口（点击url自动退订）
 - 添加自动增加Email接口，并写前端页面、或接收邮件并实现从邮件自动读取。

## 使用

 - 安装GoLang编译器
 - 安装依赖：
```go
go get github.com/PuerkitoBio/goquery
go get github.com/go-sql-driver/mysql
```

 - 修改相关数据库配置(前四项)
```go
// 数据库相关
const (
	username   = "root"
	password   = "root"
	ip         = "127.0.0.1"
	port       = "8889"
	db_name    = "jeason_daily"
	table_name = "mail_ss"
)
```

 - 注册企业邮箱，修改相关配置
```go
user := "admin@test.com"
password := "password"
host := "smtp.exmail.qq.com:25"
```

 - 编译
  - 本机编译： `go build main.go`
  - 编译服务端(CentOs 6)：`GOOS=linux GOARCH=amd64 go build main.go`

 - 通过ftp放到服务器，借助supervisor等工具做进程守护。

## 开源协议

[MIT](LICENSE)