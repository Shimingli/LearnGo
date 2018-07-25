package gotest

import (
	"testing"
	"fmt"
)


	//gotest_test.go:这是我们的单元测试文件，但是记住下面的这些原则：
	//
	//文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
	//你必须import testing这个包
	//所有的测试用例函数必须是Test开头
	//测试用例会按照源代码中写的顺序依次执行
	//测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是测试状态
	//测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
	//函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。

//  不要多余的 main函数  要不然 会执行失败
func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}

	if i, e := Division(6, 0); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过 原因是b=0") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
	fmt.Println("Test_Division_1 这个方法执行完成了哦---")
}

func Test_Division_2(t *testing.T) {
	fmt.Println("执行这个方法 Test_Division_2   ---")
	//t.Error("就是不通过=====")
}



//我们在项目目录下面执行`go test`,就会显示如下信息：
//
//--- FAIL: Test_Division_2 (0.00 seconds)
//gotest_test.go:16: 就是不通过
//FAIL
//exit status 1
//FAIL	gotest	0.013s
//从这个结果显示测试没有通过，因为在第二个测试函数中我们写死了测试不通过的代码`t.Error`，那么我们的第一个函数执行的情况怎么样呢？默认情况下执行`go test`是不会显示测试通过的信息的，我们需要带上参数`go test -v`，这样就会显示如下信息：
//
//=== RUN Test_Division_1
//--- PASS: Test_Division_1 (0.00 seconds)
//gotest_test.go:11: 第一个测试通过了
//=== RUN Test_Division_2
//--- FAIL: Test_Division_2 (0.00 seconds)
//gotest_test.go:16: 就是不通过
//FAIL
//exit status 1
//FAIL	gotest	0.012s
//上面的输出详细的展示了这个测试的过程，我们看到测试函数1`Test_Division_1`测试通过，而测试函数2`Test_Division_2`测试失败了，最后得出结论测试不通过。接下来我们把测试函数3修改成如下代码：
func Test_Division_3(t *testing.T) {
	if _, e := Division(6, 0); e == nil { //try a unit test on function
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", e) //记录一些你期望记录的信息
	}
}