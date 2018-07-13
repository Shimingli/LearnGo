package main

import (
	"fmt"
	"os"
)

func main() {
	//在任何计算机设备中，文件是都是必须的对象，而在Web编程中,文件的操作一直是Web程序员经常遇到的问题,文件操作在Web应用中是必须的,非常有用的,我们经常遇到生成文件目录,文件(夹)编辑等操作,现在我把Go中的这些操作做一详细总结并实例示范如何使用。
	fmt.Println("在任何计算机设备中，文件是都是必须的对象，而在Web编程中,文件的操作一直是Web程序员经常遇到的问题,文件操作在Web应用中是必须的,非常有用的,我们经常遇到生成文件目录,文件(夹)编辑等操作,现在我把Go中的这些操作做一详细总结并实例示范如何使用。")

	//目录操作
	//文件操作的大多数函数都是在os包里面
    osDemo1()

	//文件操作
	//建立与打开文件
	//func Create(name string) (file *File, err Error)
	//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。
	//func NewFile(fd uintptr, name string) *File
	//根据文件描述符创建相应的文件，返回一个文件对象
	//通过如下两个方法来打开文件：
	//func Open(name string) (file *File, err Error)
	//该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
	//func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
	//打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限

	//写文件
	//写文件函数：
	//func (file *File) Write(b []byte) (n int, err Error)
	//写入byte类型的信息到文件
	//func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
	//在指定位置开始写入byte类型的信息
	//func (file *File) WriteString(s string) (ret int, err Error)
	//写入string信息到文件

	fileDemo()


	//读文件
	//读文件函数：
	//func (file *File) Read(b []byte) (n int, err Error)
	//读取数据到b中
	//func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
	//从off开始读取数据到b中
	readFileDemo()


}
func readFileDemo() {
	fl,error:=os.Open("shiming11.text")

	if error!=nil {
		fmt.Println("readFileDemo",error)
		return
	}
	defer fl.Close()
	buf:=make([]byte,1024)

	for   {
		n,_:=fl.Read(buf)
		if n==0 {
			break
		}
		fmt.Println("n===============",n)
		fmt.Println("buf==========",buf)
		//	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
		//	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
		os.Stdout.Write(buf[:n])

	}

}

func fileDemo() {
	//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。
     	fout,error:= os.Create("shiming.text")
	if error!=nil {
		fmt.Println("error=",error)
		return
	}
	defer fmt.Println("shiming defer")
	defer fout.Close()
	for i:=0;i<10 ;i++  {
		fout.WriteString("shiming wo aini  我是爱你哒哒哒哒哒哒多多多多多多多多多多多多多多的  \r\n")
		fout.Write([]byte("字节 baobei baobie \r\n"))
	}

}
func osDemo1() {
	//创建名称为name的目录，权限设置是perm，例如0777
	//ModePerm FileMode = 0777 // Unix permission bits
	os.Mkdir("handletext/shimingdddd1111",0777)
	//创建目录
	os.MkdirAll("Mkdir/shiming1/shiming2",0777)

	//移除

	err :=os.Remove("Mkdir")
	if err!=nil {
		fmt.Println(err)//remove Mkdir: The directory is not empty.
	}
   // 需要使用这种的方式 才能 全部的移除掉
	os.RemoveAll("Mkdir")

	}
