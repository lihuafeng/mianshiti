package main

/**
实现发送邮件
1、发送模板邮件
 */

import (
	"fmt"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"net/smtp"
)
func main() {
	e := email.NewEmail()
	e.From = "Admin <lhf2008@yeah.net>"
	e.To = []string{"771831851@qq.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "测试邮件"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = emailBlade()
	fmt.Print(string(e.HTML))
	e.Send("smtp.yeah.net:25", smtp.PlainAuth("", "lhf2008@yeah.net", "NZWMEVSHIFXLRZTO", "smtp.yeah.net"))
}
//读取模板内容 使用go语言中的内置包，buffio和ioutil来读取
func emailBlade() []byte  {
	var txt []byte
	txt,err := ioutil.ReadFile("./email_blade.html") //读取文件
	if err != nil{
		return txt
	}
	return txt
}
