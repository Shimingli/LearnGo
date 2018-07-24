package main

import (
	"fmt"
	"errors"
	"os"
	"net/http"
)
/*
在Go中它是通过错误处理来实现的，error虽然只是一个接口，但是其变化却可以有很多，我们可以根据自己的需求来实现不同的处理，最后介绍的错误处理方案，希望能给大家在如何设计更好Web错误处理方案上带来一点思路
 */

func main() {

	fmt.Println("错误处理 ")
     /*
     Go语言主要的设计准则是：简洁、明白，简洁是指语法和C类似，相当的简单，明白是指任何语句都是很明显的，不含有任何隐含的东西，在错误处理方案的设计中也贯彻了这一思想。我们知道在C语言里面是通过返回-1或者NULL之类的信息来表示错误，但是对于使用者来说，不查看相应的API说明文档，根本搞不清楚这个返回值究竟代表什么意思，比如:返回0是成功，还是失败,而Go定义了一个叫做error的类型，来显式表达错误。在使用时，通过把返回的error变量与nil的比较，来判定操作是否成功
      */
     fil ,err:=  os.Open("dd/shiming")
	if err != nil {
		fmt.Println(err)
		fmt.Println(fil)
	}

     //  todo  标准包中所有可能出错的API都会返回一个error变量，以方便错误处理
     errDemo1()



	//自定义Error
	errDemo2()


     //错误处理
     //当有错误发生时，调用了统一的处理函数http.Error，返回给客户端500错误码，并显示相应的错误数据。但是当越来越多的HandleFunc加入之后，这样的错误处理逻辑代码就会越来越多，其实我们可以通过自定义路由器来缩减代码
     handleError1()

     //通过路由来减少代码    通过自定义路由器来缩减代码
     handleError2()

     //上面的例子错误处理的时候所有的错误返回给用户的都是500错误码，然后打印出来相应的错误代码，其实我们可以把这个错误信息定义的更加友好，调试的时候也方便定位问题，我们可以自定义返回的错误类型
     handleError3()
     // todo  在我们访问view的时候可以根据不同的情况获取不同的错误码和错误信息，虽然这个和第一个版本的代码量差不多，但是这个显示的错误更加明显，提示的错误信息更加友好，扩展性也比第一个更好






}
type appError struct {
	Error   error
	Message string
	Code    int
}
type appHandlerT func(http.ResponseWriter, *http.Request) *appError

func (fn appHandlerT) ServeHTTPT(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		c := appengine.NewContext(r)
		c.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}
func handleError3() {
	http.Handle("/view", appHandlerT(viewRecordT))
}
//这样修改完自定义错误之后，我们的逻辑处理可以改成如下方式
func viewRecordT(w http.ResponseWriter, r *http.Request) *appError {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return &appError{err, "Record not found", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appError{err, "Can't display record", 500}
	}
	return nil
}






type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//当请求/view的时候我们的逻辑处理可以变成如下代码，和第一种实现方式相比较已经简单了很多
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func handleError2() {
	//我们定义了自定义的路由器，然后我们可以通过如下方式来注册函数
	http.Handle("/view", appHandler(viewRecordTwo))
}
func viewRecordTwo(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return err
	}
	return viewTemplate.Execute(w, record)
}


func handleError1() {
	http.HandleFunc("/view", viewRecord)
}
func viewRecord(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		/*
		当有错误发生时，调用了统一的处理函数http.Error，返回给客户端500错误码，并显示相应的错误数据。但是当越来越多的HandleFunc加入之后，这样的错误处理逻辑代码就会越来越多，其实我们可以通过自定义路由器来缩减代码
		 */
		http.Error(w, err.Error(), 500)
		return
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		http.Error(w, err.Error(), 500)
	}
}





func errDemo2() {
	// todo  在实现自己的包的时候，通过定义实现此接口的结构，我们就可以实现自己的错误定义
	//type SyntaxError struct {
	//	msg    string // 错误描述
	//	Offset int64  // 错误发生的位置
	//}
	//
	//func (e *SyntaxError) Error() string { return e.msg }

	//if err := dec.Decode(&val); err != nil {
	//	if serr, ok := err.(*json.SyntaxError); ok {
	//		line, col := findLine(f, serr.Offset)
	//		return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
	//	}
	//	return err
	//}

	//上面例子简单的演示了如何自定义Error类型。但是如果我们还需要更复杂的错误处理呢？此时，我们来参考一下net包采用的方法：
	//
	//package net
	//
	//type Error interface {
	//	error
	//	Timeout() bool   // Is the error a timeout?
	//	Temporary() bool // Is the error temporary?
	//}
	//在调用的地方，通过类型断言err是不是net.Error,来细化错误的处理，例如下面的例子，如果一个网络发生临时性错误，那么将会sleep 1秒之后重试：
	//
	//if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
	//	time.Sleep(1e9)
	//	continue
	//}
	//if err != nil {
	//	log.Fatal(err)
	//}




}
func errDemo1() {

	//error类型是一个接口类型，这是它的定义：
	//
	//type error interface {
	//	Error() string
	//}
	//error是一个内置的接口类型，我们可以在/builtin/包下面找到相应的定义。而我们在很多内部包里面用到的 error是errors包下面的实现的私有结构errorString
	//
	//// errorString is a trivial implementation of error.
	//type errorString struct {
	//	s string
	//}
	//
	//func (e *errorString) Error() string {
	//	return e.s
	//}
	//你可以通过errors.New把一个字符串转化为errorString，以得到一个满足接口error的对象，其内部实现如下：
	//
	//// New returns an error that formats as the given text.
	//func New(text string) error {
	//	return &errorString{text}
	//}

	fmt.Println(errors.New("我不管 我就像报错了  咋地 不服气么 "))

	f,err:= errSqrt(-1)
	if err != nil {
		fmt.Println(err)
		fmt.Println(f)
	}


}
func errSqrt(f float64) (float64,error)  {

	if f<0 {
		//  todo fmt.Println(fmt包在处理error时会调用Error方法)被调用，以输出错误
		return 0, errors.New("不能小于0啊")
	}
	return f,nil
}
