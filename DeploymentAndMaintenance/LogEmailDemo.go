package main

func main() {
	Email()
}

func Email(){
	//logger:= logrus.New()

	////parameter"APPLICATION_NAME", "HOST", PORT, "FROM", "TO"
	////首先开启smtp服务，最后两个参数是smtp的用户名和密码
	//hook, err := logrus_mail.NewMailAuthHook("testapp", "smtp.163.com",25,"username@163.com","username@163.com","smtp_name","smtp_password")
	//if err == nil {
	//	logger.Hooks.Add(hook)
	//}
	////生成*Entry
	//var filename="123.txt"
	//contextLogger :=logger.WithFields(logrus.Fields{
	//	"file":filename,
	//	"content":  "GG",
	//})
	////设置时间戳和message
	//contextLogger.Time=time.Now()
	//contextLogger.Message="这是一个hook发来的邮件"
	////只能发送Error,Fatal,Panic级别的log
	//contextLogger.Level=logrus.FatalLevel
	//
	////使用Fire发送,包含时间戳，message
	//hook.Fire(contextLogger)
}
