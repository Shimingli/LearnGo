package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

func init() {
//	Go中支持MySQL的驱动目前比较多，有如下几种，有些是支持database/sql标准，而有些是采用了自己的实现接口,常用的有如下几种:
//
//https://github.com/go-sql-driver/mysql 支持database/sql，全部采用go写。
//https://github.com/ziutek/mymysql 支持database/sql，也支持自定义的接口，全部采用go写。
//https://github.com/Philio/GoMySQL 不支持database/sql，自定义接口，全部采用go写。
//	接下来的例子我主要以第一个驱动为例(我目前项目中也是采用它来驱动)，也推荐大家采用它，主要理由：
//
//	这个驱动比较新，维护的比较好
//	完全支持database/sql接口
//	支持keepalive，保持长连接,虽然星星fork的mymysql也支持keepalive，但不是线程安全的，这个从底层就支持了keepalive。
}
func main() {
//	user@unix(/path/to/socket)/dbname?charset=utf8
//user:password@tcp(localhost:5555)/dbname?charset=utf8
//user:password@/dbname
//user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname


	//spring.datasource.url =jdbc:mysql://localhost:3306/test
	//spring.datasource.username = root
	//#  注意密码的问题 ，数据的密码的
	//spring.datasource.password = App123
    // sql.Open()函数用来打开一个注册过的数据库驱动，go-sql-driver中注册了mysql这个数据库驱动，第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息   如上的图片

	db, err := sql.Open("mysql", "root:App123@tcp(localhost:3306)/godbdemo?charset=utf8")
	checkErr(err)


	//插入数据 db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)
	//var s []string
	//for i:=1;i<100 ;i++  {
	//	//stmt.Exec()函数用来执行stmt准备好的SQL语句
	//	//s:="研发部门--i"
	//	var s []string
	//	s = append(s, strconv.Itoa(i),"研发部门--i====")
	//	res, err := stmt.Exec("shiming",strings.Join(s,strconv.Itoa(i)) , "20180709")
	//	checkErr(err)
	//	//shiming i== 9727 res=== {0xc0420a0000 0xc0422f5000}
	//	fmt.Println("shiming i==",i,"res===",res)
	//
	//}
	res, err := stmt.Exec("shiming", "研发部门", "20180706")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据  传入的参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("shimingupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据  db.Query()函数用来直接执行Sql返回Rows结果。
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
    stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
     //stmt.Exec()函数用来执行stmt准备好的SQL语句
	res, err = stmt.Exec(id)
	checkErr(err)
	//RowsAffected 返回受更新、插入或删除影响的行数。并非每个数据库或数据库驱动程序都可以支持这一点。
	affect, err = res.RowsAffected()
	checkErr(err)
	//
	//fmt.Println(affect)
    //  todo  删除数据库
	res, err = db.Exec("DELETE  FROM userinfo WHERE  username = ?", "shiming")
	fmt.Println(res)
	id,err = res.LastInsertId()
	fmt.Println("删除数据的id==",id)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err) //大概相当于  Java中的异常
	}
}
