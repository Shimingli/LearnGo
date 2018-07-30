package Old

import "fmt"
/**
Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。
 */
func init() {
   fmt.Println("init 开始执行   Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。  " )
   structDemo()
}
func structDemo() {
	//第一种初始化办法
	var p  person
	p.name="shiming"
	p.age=18
	fmt.Println("这个人的信息",p)//{shiming 18}  如果按照 java的理解  就是  重写了toString的方法

	//第二种初始化办法  1.按照顺序提供初始化值
	P1:=person{"shiming",20}
	fmt.Println("第二种初始化的办法",P1)

	//第三种初始化的方法  //可以没有顺序 2.通过field:value的方式初始化，这样可以任意顺序
	P2:=person{age:15,name:"shiming"}
	fmt.Println("么有顺序的第三种的初始化的方法",P2)

	//第四种也可以通过new 函数分配一个指正，此处P2的类型为 *person
	P3:=new(person)
	fmt.Println("第四总初始化的方法",P3)//第四总初始化的方法 &{ 0}
    //上面的输出的结果   &{ 0}
	//demo
	P4:=new(person)
	P4.name="shiming"
	P4.age=22
	fmt.Println("demo 的演示的方法----->",P4)//第四总初始化的方法 &{shiming 22}


	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{age:25, name:"Bob"}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}
    //第一值是那个更加的大  ，第二个值，是相差多少岁
	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)


	//struct的匿名字段
	//实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。
	//当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。


	 //初始化  Student   感觉有点像继承啊
	 var s Student
	 s.age=12
	 s.name="shiming"
	 s.speciality="ge"
	 s.weight=180
	 fmt.Println("打印出来的感觉像继承啊",s)



	// 我们初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	fmt.Println("没有修改它的信息之前为  ",mark)

	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
	fmt.Println("修改所有完了的信息了",mark)

	//  Student组合了Human struct和string基本类型    注意看  2.4.student_struct.png 的图片
	//我们看到Student访问属性age和name的时候，就像访问自己所有用的字段一样，对，匿名字段就是这样，能够实现字段的继承。是不是很酷啊？还有比这个更酷的呢，那就是student还能访问Human这个字段作为字段名。请看下面的代码，是不是更酷了。

	mark.Human=Human{"shiming",10,50}
	fmt.Println("简写方法之前的值",mark)
	mark.Human.age=15
	fmt.Println("简写玩了 ，修改了其中的值",mark)


	//通过匿名访问和修改字段相当的有用，但是不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段的。
	// 初始化学生Jane
	jane := Student1{Human1:Human1{"Jane", 35, 100}, speciality:"Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
	//就是呢 如果不是标准的类型的话，这个就必须指定你给的类型 ，如果是的话  ，比如说是 int  string 那么久直接指定它 就行了
	P222  := Student1{Human1:Human1{"shiming",10,545},Skills:Skills{"ddd"},int:int(10),speciality:string("dddd")}
    //struct不仅仅能够将struct作为匿名字段，自定义类型、内置类型都可以作为匿名字段，而且可以在相应的字段上面进行函数操作（如例子中的append）。
    P222.Skills=append(  P222.Skills,"shiming","baobei")
	fmt.Println("我自己定义的类型",P222)
	//如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？
	 //     todo    和java 是一样的    想访问之类的变量
	//Go里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段。
	//这样就允许我们去重载通过匿名字段继承的一些字段，当然如果我们想访问重载后对应匿名类型里面的字段，可以通过匿名字段名来访问
	Bob := Employee{Human2{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human2.phone)

}

type Human2 struct {
	name string
	age int
	phone string  // Human类型拥有的字段
}

type Employee struct {
	Human2  // 匿名字段Human
	speciality string
	phone string  // 雇员的phone字段
}

type Skills []string
type Human1 struct {
	name string
	age int
	weight int
}
type Student1 struct {
	Human1  // 匿名字段，struct
	Skills // 匿名字段，自定义的类型string slice
	int    // 内置类型作为匿名字段
	speciality string
}


type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}


// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age>p2.age {  // 比较p1和p2这两个人的年龄
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}


//Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。
type person struct {
	name string//一个string类型的字段name，用来保存用户名称这个属性
	age int//一个int类型的字段age,用来保存用户年龄这个属性
}
/**
如果这个是个main 包的下面的 ，这个main 的方法必须要
 */
func main()  {
	fmt.Println("<---------------------------------------------->")
	fmt.Println("main 后面执行")
}
