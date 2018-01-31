1. 定义
interface 类型定义医嘱方法, 这些方法不需要实现, interface 不能包含任何变量

		type Carer interface {
			method1(参数列表)  返回值
			Name(a string) string  
			Run()
			DiDi()
		}

1.1 示例
// 指正类型的方法实现
	package main

	import "fmt"

	type Carer interface {
		GetName() string
		Run()
		DiDi()
	}

	type BMW struct {
		name	string
	}

	func (p *BMW) GetName() string  {   // 指正类型
		return p.name
	}
	func (p *BMW)Run()  {		// 指正类型
		fmt.Printf("%s is running \n", p.name)
	}
	func (p *BMW)DiDi()  {		// 指正类型
		fmt.Printf("%s is didi\n",p.name )
	}

	func main() {
		var car Carer

		var bmw BMW
		bmw.name="byd"

		car = &bmw    // 传递指针
		car.Run()
	}
	//////////main 另外一种写法
	func main()  {

	var a Car
	b := &BMW{
		Name:"benchi",
	}
	a = b
	a.Run()

	}
//
// 值类型方法实现
	package main

	import "fmt"

	type Carer interface {
		GetName() string
		Run()
		DiDi()
	}

	type BMW struct {
		name	string
	}

	func (p BMW) GetName() string  {   // 值类型
		return p.name
	}
	func (p BMW)Run()  {		// 值类型
		fmt.Printf("%s is running \n", p.name)
	}
	func (p BMW)DiDi()  {		// 值类型
		fmt.Printf("%s is didi\n",p.name )
	}

	func main() {
		var car Carer

		var bmw BMW
		bmw.name="byd"

		car = bmw   // 值类型
		car.Run()
	}



/////////////////////////////////////////////////
//	sort接口示例
		package main

		import (
			"fmt"
			"math/rand"
			"sort"
		)

		type Student struct {
			name	string
			id 		string
			age 	int

		}

		type StudentArry  []Student
	//sort 需要的接口 
	/*
	type Interface interface {
	    // Len方法返回集合中的元素个数
	    Len() int
	    // Less方法报告索引i的元素是否比索引j的元素小
	    Less(i, j int) bool
	    // Swap方法交换索引i和j的两个元素
	    Swap(i, j int)
	}*/

		func (p StudentArry ) Len() int  {
			return len(p)
		}

		func (p StudentArry)Less(i,j int) bool  {
			return p[i].name < p[j].name
		}

		func (p StudentArry) Swap(i,j int)  {
			p[i], p[j] = p[j], p[i]
		}



		func main()  {

//生成 数组
			var stulist   StudentArry
			for i :=0; i< 10; i++{
				stu  := Student{
					name:fmt.Sprintf("stu%d", rand.Intn(100)),
					id:fmt.Sprintf("1000%d",rand.Int()),
					age:rand.Intn(100),
				}
				stulist = append(stulist,stu)
			}
//打印数组
			for _,v := range stulist{
				fmt.Println(v)
			}
	//数组排序
			sort.Sort(stulist)
			fmt.Println()
			fmt.Println()

			for _,v := range stulist{
				fmt.Println(v)
			}

		}


1.2 类型断言
package main

import (
	"fmt"

)

func classifer(items ... interface{})  { // 可以传递任何数量任何类型的参数

	for _,v := range items{
		switch v.(type) {    // 显示类型
		case bool:
			fmt.Printf("valuse is %T",v)
		case int,int32,int64:
			fmt.Printf("valuse is %T",v)
		case float64,float32:
			fmt.Printf("valuse is %T",v)
		case string:
			fmt.Printf("valuse is %T",v)
		default:
			fmt.Printf("unknow type of %T",v)
		}
	}

}

func main()  {
	var b int
	classifer(2131,"sca",&b)


}