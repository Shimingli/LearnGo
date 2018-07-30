package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作 ，客户端通过multipart.Write把文件的文本流写入一个缓存中，然后调用http的Post方法把缓存传到服务器。
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("状态---》",resp.Status)
	fmt.Println("结果-----》",string(resp_body))
	return nil

	//  todo  你还有其他普通字段例如username之类的需要同时写入，那么可以调用multipart的WriteField方法写很多其他类似的字段。  是不是这就是代码的测试啊 写个无线的循环 哈哈哈  直接跑起来  啊哈啊哈哈 啊哈哈哈哈哈哈哈啊哈哈哈哈哈哈哈
}

// sample usage   o支持模拟客户端表单功能支持文件上传
func main() {
	target_url := "http://localhost:9090/upload"
	/**
	我在 test的目录下，又重新创建了一个form的目录  才可以upload 成功

	如果没有手动的创建的话 ，就会报错
	open ./test/form/upload.gtpl: The system cannot find the path specified.

	 */
	filename := "form/upload.gtpl"  //
	//filename := "./astaxie.pdf"
	postFile(filename, target_url)
}