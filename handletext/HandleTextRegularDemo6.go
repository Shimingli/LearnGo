package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("如果想截取字符串的一部分，过滤字符串。或者是提取符合字符条件的一批字符串")
	//  需要使用正则表达式的复杂的模式

	// 爬虫来过滤和截取抓到的数据
	resp,error:=http.Get("http://www.baidu.com")

	if error!=nil {
		fmt.Println(error)
	}

	fmt.Println("resp===",resp)
	//{200 OK 200 HTTP/1.1 1 1 map[Bdpagetype:[1] Cache-Control:[private] Content-Type:[text/html] Expires:[Wed, 11 Jul 2018 10:55:46 GMT] Server:[BWS/1.1] Cxy_all:[baidu+b5331167871983fb80121ae54f1cdff5] Date:[Wed, 11 Jul 2018 10:56:32 GMT] P3p:[CP=" OTI DSP COR IVA OUR IND COM "] Set-Cookie:[BAIDUID=0E00678867192DA9EC07D8D1F743DBC4:FG=1; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com BIDUPSID=0E00678867192DA9EC07D8D1F743DBC4; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com PSTM=1531306592; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com BDSVRTM=0; path=/ BD_HOME=0; path=/ H_PS_PSSID=1466_26431_21127_18560_22073; path=/; domain=.baidu.com] Vary:[Accept-Encoding] Bdqid:[0xcde2d51000091a46] Connection:[Keep-Alive] X-Ua-Compatible:[IE=Edge,chrome=1]] 0xc0420fe220 -1 [chunked] false true map[] 0xc0420dc000 <nil>}
	defer resp.Body.Close()

	//转化为了一个 自己数组
	boby,error:=ioutil.ReadAll(resp.Body)
	if error!=nil {
		fmt.Println(error)
	}
	fmt.Println("body=====",boby)
    src:= string(boby)
    fmt.Println("src===",src)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	fmt.Println("--------------------------     ------------------------------------")
	fmt.Println("src===",src)


	//使用复杂的正则首先是Compile，它会解析正则表达式是否合法，如果正确，那么就会返回一个Regexp，然后就可以利用返回的Regexp在任意的字符串上面执行需要的操作



      demoText()


}
func demoText() {
	a:= "I am learning Go language"
	//这个正则表达式的意思就是呢  找小写的字母，长度最短为2 最长为4
	re,_:=regexp.Compile("[a-z]{2,4}")

	fmt.Println("re:",re)//[a-z]{2,4}
	// 将字符串 a 转换为 []byte 类型
	//查找符合正则的第一个  查找返回一个保存正则表达式B中最左边匹配文本的片段。
	one:= re.Find([]byte(a))
	fmt.Println("Find:one====",one)//[97 109]
	fmt.Println("Find:",string(one))//am

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
     all :=re.FindAll([]byte(a),-1)
     fmt.Println("all==",all)//[[97 109] [108 101 97 114] [110 105 110 103] [108 97 110 103] [117 97 103 101]]
     fmt.Println("all==",string(all[0]))
     fmt.Println("all==",string(all[1]))



	//查找符合条件的index位置,开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)


	//I am learning Go language

	//查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
	//下面的输出第一个元素是"am learning Go language"
	//第二个元素是" learning Go "，注意包含空格的输出
	//第三个元素是"uage"
	re2, _ := regexp.Compile("am(.*)lang(.*)")
	submatch := re2.FindSubmatch([]byte(a))
	// [[97 109 32 108 101 97 114 110 105 110 103 32 71 111 32 108 97 110 103 117 97 103 101] [32 108 101 97 114 110 105 110 103 32 71 111 32] [117 97 103 101]]
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println("shiming == ",string(v))
	}


	//定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println("submatchindex==",submatchindex)



	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)



	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
