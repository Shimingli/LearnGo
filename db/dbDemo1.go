package main

import (
	"fmt"
)


func init() {
	fmt.Println("Go官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动，这样做有一个好处，只要是按照标准接口开发的代码， 以后需要迁移数据库时，不需要任何修改")

}

func main() {
	//  https://github.com/Shimingli/build-web-application-with-golang/blob/master/zh/05.1.md  有点迷
	fmt.Println("https://github.com/Shimingli/build-web-application-with-golang/blob/master/zh/05.1.md")
}
//sql.Register  数据库的注册
//这个存在于database/sql的函数是用来注册数据库驱动的，当第三方开发者开发数据库驱动时，都会实现init函数，在init里面会调用这个Register(name string, driver driver.Driver)完成本驱动的注册。
//
//mymysql、sqlite3的驱动里面都是怎么调用
//https://github.com/mattn/go-sqlite3驱动
//func init() {
//	sql.Register("sqlite3", &SQLiteDriver{})
//}

//https://github.com/mikespook/mymysql驱动
//// Driver automatically registered in database/sql
//var d = Driver{proto: "tcp", raddr: "127.0.0.1:3306"}
//func init() {
//	//Register("SET NAMES utf8")
//	//sql.Register("mymysql", &d)
//}


