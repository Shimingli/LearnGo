package main

import (
	"fmt"
	"regexp"
	"os"
)

func init() {
}
func main() {


	fmt.Println("正则表达式是一种进行模式匹配和文本操纵的复杂而又强大的工具。虽然正则表达式比纯粹的文本匹配效率低，但是它却更灵活。按照它的语法规则，随需构造出的匹配模式就能够从原始文本中筛选出几乎任何你想要得到的字符组合。如果你在Web开发中需要从一些文本数据源中获取数据,那么你只需要按照它的语法规则，随需构造出正确的模式字符串就能够从原数据源提取出有意义的文本信息。")

	// Go 语言通过 regexp 标准包为正则表达式提供了官方的支持

	b:=IsIP("451.455.154.12")
    fmt.Println("是否是IP地址",b)

	//输入Demo
	osDemo()
	//if len(os.Args)==1 {
	//	fmt.Println("Usage:regexp [string]")
	//	os.Exit(1)
	//}

}

//当用户输入一个字符串，我们想知道是不是一次合法的输入：
func osDemo() {
	if len(os.Args)==1 {
		fmt.Println("Usage:regexp [string]")
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
}
func IsIP(ip string) (b bool) {
	 if m,_:=regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip) ;!m{
	 	return false
	 }
	 return true
}
