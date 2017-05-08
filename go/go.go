1. go 基础知识


[01:36] Go内置关键字和注释方法
[02:08] Go程序的一般结构
[07:25] 包的导入
[09:43] package别名与省略调用(不建议用)
[11:45] 可见性规则


		package main

		import "fmt"

		定义别名
		import alias "os"

		const pt  =  "sssc"

		全局变量的声明和赋值
		var name  = 123

		type nas int
		type jie struct {
		}
		type inte interface {
		}

		func main()  {
			a := 1
			b := 2
			fmt.Println(a)
			fmt.Println(b)
			别名调用
			fmt.Println(alias.TempDir())
		}

1.2可见性规则
	go 使用大小写来决定是否能被外部调用, 首字母大写即为public , 首字母小写即为priviate


3. 类型与变量
[04:57] 基本类型
[11:25] 类型零值
[16:30] 类型别名
[19:20] 变量的声明与赋值
[29:20] 类型转换

3.1 基本类型
	- 布尔  true false
	- 整形  int/uint
	- 8位整形  int8/uint8(别名byte)
		取值范围: -128~127 / 0~255
	- 16位整形 int16/uint16
		取值范围: -2^16/2 ~2^16/2-1  / 0~2^16
	- 32位整形
	- 64位整形
	- 浮点型

4. 常量与运算符
[04:35] 常量的定义
[11:30] 常量的枚举
[17:05] 运算符

4.1 int string 的相互转换
package main

import "fmt"
import "strconv"

func main()  {
	a := 65
	fmt.Println(a)
	//将int 装换成string
	b:= strconv.Itoa(a)
	//将int 转换成string ,因为fmt.Println(strconv.Atoi(b))  出来的结果是65 <nil> ,要取第一个,所以用c,_忽略第二个输出
	c,_:= strconv.Atoi(b)

	fmt.Println(b)
	fmt.Println(c)

}

4.2 常量的定义
func main()  {
	if a,b:=132,322; a > 10{
		fmt.Println(a+b)
	}

}


func main()  {
	a:=1
	for i:=1; i<3; i++{
	a++
	fmt.Println(a)
	}

}
/////////////////////////////////////////
case 语句
5.1.
	func main()  {
		depth:=121
		switch depth {
		case 1:
			fmt.Println(depth)
		case 2:
			fmt.Println(depth)
		//默认值
		default:
			fmt.Println("default")
		}
	}
5.2. fallthrough
	//加了fallthrough 之后, 匹配到第一条规则之后还会继续往下匹配
	func main()  {
		switch depth:=2; {
		case depth >1:
			fmt.Println(">1")
			fallthrough
		case depth >= 2:
			fmt.Println(">=2")
		default:
			fmt.Println("default")
		}
	}
	//输出
	>1
	>=2	

5.3. LABEL
	// 可以指定break 到指定的label , 注意冒号
	func main()  {
		LABLE1:
			for {
				for i:=1; i<10 ;i++{
					if i > 3{
						fmt.Println(i)
						//直接跳出循环到LABEL1 那一层
						break LABLE1
					}
				}
			}
	}
5.4. goto  
	调整程序的执行位置,注意LABEL 必须在goto 代码后面, 不然会进入死循环
	func main()  {
			for {
				for i:=1; i<10 ;i++{
					if i > 3{
						fmt.Println(i)
						goto LABLE1
					}
				}
			}
		LABLE1:
		fmt.Println("goto 代码")
	}	

5.5. continue 
跳出当前循环进入下一次循环
	func main()  {
	LABEL1:
		for i:=1; i<10; i++{
			for {
				continue LABEL1
				fmt.Println("inside")
			}
		}
		fmt.Println("outside")
	}

6. 数组
[04:23] array的定义
[10:02] 指向数组的指针和指针数组
[12:30] 数组之间的比较
[14:22] 使用new创建数组
[16:09] 多维数组
[18:45] 冒泡排序

6.1 定义
[...]表示自动计算数组长度
	func main()  {
		a :=[...]string{"sacascac","wdqwdw"}
		b:=[]int{1,2,3,45}
		fmt.Println(a)
		fmt.Println(b)
	}
 输出:
	[sacascac wdqwdw]
	[1 2 3 45]

6.2定义数组长度为10, 第六个数=12 ,其他的默认都是0值
	func main()  {
		a:=[10]int{5:12}
		fmt.Println(a)
	}
	输出:
	[0 0 0 0 0 12 0 0 0 0]
6.3不指定数组长度, 会自动把定义的某个index作为最后一个, 下面指定了第6个数组=12, name数组长度就是6
	func main()  {
		a:=[...]int{5:12}
		fmt.Println(a)
	}
	[0 0 0 0 0 12]

6.4多维数组 ,
	注意int{}大口号的位置,

		func main()  {
			a:=[3][2]int{
				{1,2},
				{2,3},
				{3,5}} //{ 必须在这一行 ,换行的话就要加, 
			fmt.Println(a)
		}

		func main()  {
			a:=[3][2]int{
				{1,2},
				{2,3},
				{3,5},
			}
			fmt.Println(a)
		}
	以上两个输出都是  [[1 2] [2 3] [3 5]]
指定某个位置的值
		func main()  {
		a:=[3][2]int{
			{1,1:4},
			{2,3},
			{3,5},
		}
		fmt.Println(a)
	}

7. 切片
[03:15] slice 概述
[05:20] 创建 slice
[18:25] reslice 概述
[23:00] append()与slice
[29:55] copy()与slice

7.1 切片
	func main()  {
		a:=[]int{1,2,3,4,5,6,7}
		fmt.Println(a)
		fmt.Println(len(a))
		s1:=a[:]   // 全部数组长度
		fmt.Println(s1)
	}

	func main()  {
		a:=[]int{1,2,3,4,5,6,7}
		fmt.Println(a)
		s1:=a[:5]
		fmt.Println(s1) 
		fmt.Println(s1[:6])  // s1本身并没有6这个索引,但是任然可以获取到值,这个值是原始数组a 的值
	}
	输出:
	[1 2 3 4 5 6 7]
	[1 2 3 4 5]   //s1的长度
	[1 2 3 4 5 6]  // 超过s1的长度 ,取数组a的值

7.2 append
func main() {
	a:=[]int{1,2,3,4}
	b:=append(a,123,4,5,6)  // 在a 后面append 
	fmt.Println(b)
}

7.3 copy

func main() {
	s1 := []int{1,2,3,4,5,6}
	s2:= []int{7,8,9}
	copy(s1,s2) //把s2copy到s1
	fmt.Println(s1)
}
输出:
[7 8 9 4 5 6]

func main() {
	s1 := []int{1,2,3,4,5,6}
	s2:= []int{7,8,9}
	copy(s2,s1)
	fmt.Println(s2)
}
输出:
[1 2 3]
指定位置copy
func main() {
	s1 := []int{1,2,3,4,5,6}
	s2:= []int{7,8,9}
	copy(s2[1:3],s1[3:5])
	fmt.Println(s2)
}
输出:
[7 4 5]