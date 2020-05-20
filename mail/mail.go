package mail

import (
	"bytes"
	"crypto/tls"
	"gin-monitor/model"
	"github.com/apex/log"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"html/template"
	"net/smtp"
	"time"
)

type Mail struct {
	Time string `json:"time" form:"time"`
	Results []model.ResultMgo `json:"results" form:"results"`
}

func GetInfo() []model.ResultMgo {
	var res []model.ResultMgo
	s, col := model.GetCol("results")
	defer s.Close()
	if  err := col.Find(nil).All(&res); err != nil {
		log.Errorf(err.Error())
	}
	return res
}

func DeleteAll()  {
	s, col := model.GetCol("results")
	defer s.Close()
	res := GetInfo()
	for i := range res {
		task := res[i]
		err := col.Remove(&task)
		if err != nil {
			log.Errorf(err.Error())
		}
		log.Info("deleted")
	}
}

func SendEmail(){
	var serverAddr = viper.GetString("mail.server")
	var port = viper.GetString("mail.stmp.port")
	var emailAddr = viper.GetString("mail.stmp.user")
	var password = viper.GetString("mail.stmp.password")
	e := email.NewEmail()
	e.From = emailAddr
	e.To = viper.GetStringSlice("mail.receiver")
	e.Subject = viper.GetString("mail.subject")

	t,err := template.ParseFiles("mail/mail-template.html")
	if err != nil {
		log.Errorf(err.Error())
	}
	body := new(bytes.Buffer)
	res := GetInfo()
	var mail = Mail{
		Time: time.Now().Format("2006/01/02 15:04:05"),
		Results: res,
	}
	err = t.Execute(body, mail)
	if err != nil{
		log.Errorf(err.Error())
	}
	// html形式的消息
	e.HTML = body.Bytes()
	err = e.SendWithTLS(serverAddr+":"+port, smtp.PlainAuth("", emailAddr, password, serverAddr), &tls.Config{ServerName: serverAddr})
	if err != nil {
		log.Errorf(err.Error())
	}
	DeleteAll()
}