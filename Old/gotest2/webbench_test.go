package gotest

import (
	"testing"
)

// todo  测试的最好分开两个包 --才可以测试完成

/*
压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：

压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
	func BenchmarkXXX(b *testing.B) { ... }
go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数
在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行
文件名也必须以_test.go结尾
 */
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

/*

E:\new_code\GoDemo\gotest2>go test -test.bench=".*"
goos: windows
goarch: amd64
Benchmark_Division-4                    2000000000               1.19 ns/op
Benchmark_TimeConsumingFunction-4       2000000000               0.75 ns/op
PASS
ok      _/E_/new_code/GoDemo/gotest2    4.105s
 */

//  todo  上面的结果显示我们没有执行任何TestXXX的单元测试函数，显示的结果只执行了压力测试函数，第一条显示了Benchmark_Division执行了2000000000次，每次的执行平均时间是1.19纳秒，第二条显示了Benchmark_TimeConsumingFunction执行了2000000000 ，每次的平均执行时间是0.75纳秒。最后一条显示总共的执行时间。 4.105s