package logs

import (
	"github.com/cihub/seelog"
	"fmt"
)
/*
通过上面对seelog系统及如何基于它进行自定义日志系统的学习，现在我们可以很轻松的随需构建一个合适的功能强大的日志处理系统了。日志处理系统为数据分析提供了可靠的数据源，比如通过对日志的分析，我们可以进一步优化系统，或者应用出现问题时方便查找定位问题，另外seelog也提供了日志分级功能，通过对minlevel的配置，我们可以很方便的设置测试或发布版本的输出消息级别。
 */
var Logger seelog.LoggerInterface

//    todo   loadAppConfig
//根据配置文件初始化seelog的配置信息，这里我们把配置文件通过字符串读取设置好了，当然也可以通过读取XML文件。里面的配置说明如下：
//seelog
//minlevel参数可选，如果被配置,高于或等于此级别的日志会被记录，同理maxlevel。
//outputs
//输出信息的目的地，这里分成了两份数据，一份记录到log rotate文件里面。另一份设置了filter，如果这个错误级别是critical，那么将发送报警邮件。
//formats
//定义了各种日志的格式
func loadAppConfig() {
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="/data/logs/roll.log" maxsize="100000" maxrolls="5"/>
		<filter levels="critical">
            <file path="/data/logs/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="astaxie@gmail.com" sendername="ShortUrl API" hostname="smtp.gmail.com" hostport="587" username="mailusername" password="mailpassword">
                <recipient address="xiemengjun@gmail.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
	    <format id="critical" format="%File %FullPath %Func %Msg%n" />
	    <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

//  初始化全局变量Logger为seelog的禁用状态，主要为了防止Logger被多次初始化
func DisableLog() {
	Logger = seelog.Disabled
}


//UseLogger
//设置当前的日志器为相应的日志处理
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}