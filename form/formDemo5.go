package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"time"
	"crypto/md5"
	"strconv"
	"html/template"
)

func main() {
	fmt.Println("处理文件的上传")
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9090",nil)
}
// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("form/upload.gtpl")
		t.Execute(w, token)
	} else {
		//调用r.ParseMultipartForm，里面的参数表示maxMemory，调用ParseMultipartForm之后，上传的文件存储在maxMemory大小的内存里面，如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中

		/**

		32 的二进制为    100000   左移一位相当于 乘以 2
		32<<20  左移20位，相当于乘以20个2
2^10  ===   1 024
       预留额外的10 MB的非文件部分。
		// Reserve an additional 10 MB for non-file parts.
	maxValueBytes := maxMemory + int64(10<<20)
		 */
        // 所以  todo 下面的代表最多是 32M的文件 ，还预留了  10的非文件的部分，源码中 调用了 r.ParseForm()   所以不要感到奇怪
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")// 读取表单中的 uploadfile的key
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		/*
// A FileHeader describes a file part of a multipart request.
type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64

	content []byte
	tmpfile string
}
		 */

		//map[Content-Disposition:[form-data; name="uploadfile"; filename="oMCb64uc0JAG2HxbEmjfAV5WLibw.jpg"] Content-Type:[image/jpeg]]
		fmt.Println("handler.Header=",handler.Header)
		fmt.Fprintf(w, "%v", handler.Header)
		// 这个 test文件夹是我自己创建的 这样就可以通过 上传上来了
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)  // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		// defer  Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题
		//如果有很多调用defer，那么defer是采用后进先出模式   用于类似析构函数
		defer f.Close()
		io.Copy(f, file)

		//表单中增加enctype="multipart/form-data"
		//服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
		//使用r.FormFile获取文件句柄，然后对文件进行存储等处理。
	}
}