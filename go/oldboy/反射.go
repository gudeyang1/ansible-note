1. 示例
		package main

		import (
			"fmt"
			"reflect"
		)

		type Student struct {
			Name string
			Age int
		}
		func test(b interface{}){
			fmt.Println(reflect.TypeOf(b))  //main.Student
			fmt.Println(reflect.ValueOf(b))  //{jj 12}
			fmt.Println(reflect.ValueOf(b).Kind()) //struct

		}

		func main()  {
		 a := Student{
			Name:"jj",
			Age:12,
		}
			test(a)
		}
 


////////////////////////////////////////////////
获取字段数量
获取字段方法数量
////////////////////////////////
		package main

		import (
			"reflect"
			"fmt"
		)

		type Student struct {
			name string
			age int
			score float32
		}
		//定义struct 方法
		func (p Student)Set(name string ,age int, score float32 )  {
			p.name = name
			p.age = age
			p.score = score
		}

		func TestStruct(a interface{})  {
			value := reflect.ValueOf(a)
			kind := value.Kind()
			//判断传入的参数类型是否为struct
			if kind != reflect.Struct{
				fmt.Println("expect struct")
				return
			}
			fmt.Println(value.NumField()) // 字段数量 3个
			fmt.Println(value.NumMethod())  //struct的方法数量  


		}

		func main()  {
			var a  = Student{
				name:"stu01",
				age: 18,
				score: 93.2,
			}
			TestStruct(a)
		}


//////////////////
用反射调用struct 方法
		package main

		import (
			"reflect"
			"fmt"
		)

		type Student struct {
			name string
			age int
			score float32
		}
		//定义struct 方法
		func (p Student)Set(name string ,age int, score float32 )  {
			p.name = name
			p.age = age
			p.score = score
		}

		func (p Student)Print()  {
			fmt.Println("--------start-------")
			fmt.Println("--------stop-------")

		}

		func TestStruct(a interface{})  {
			value := reflect.ValueOf(a)
			kind := value.Kind()
			//判断传入的参数类型是否为struct
			if kind != reflect.Struct{
				fmt.Println("expect struct")
				return
			}
			fmt.Println(value.NumField()) // 字段数量 3个
			fmt.Println(value.NumMethod())  //struct的方法数量

			//调用struct 方法
			var params  []reflect.Value
			value.Method(0).Call(params) // 0表示调用第一个方法


		}

		func main()  {
			var a  = Student{
				name:"stu01",
				age: 18,
				score: 93.2,
			}
			TestStruct(a)
		}
	输出:
		3
		2
		--------start-------
		--------stop-------		.

/////////
打印每个字段的值 value.Field(i), 类型:value.Field(i).Kind()		
		package main

		import (
			"reflect"
			"fmt"

		)

		type Student struct {
			name string
			age int
			score float32
		}
		//定义struct 方法
		func (p Student)Set(name string ,age int, score float32 )  {
			p.name = name
			p.age = age
			p.score = score
		}

		func (p Student)Print()  {
			fmt.Println("--------start-------")
			fmt.Println("--------stop-------")

		}

		func TestStruct(a interface{})  {
			value := reflect.ValueOf(a)
			kind := value.Kind()
			//判断传入的参数类型是否为struct
			if kind != reflect.Struct{
				fmt.Println("expect struct")
				return
			}
			fmt.Println(value.NumField()) // 字段数量 3个
		// 打印每个字段的值 value.Field(i), 类型:value.Field(i).Kind()
			num := value.NumField()
			for i := 0 ; i< num; i++{
				fmt.Printf("%d %v %s\n",i,value.Field(i),value.Field(i).Kind())
			}

			
		}

		func main()  {
			var a  = Student{
				name:"stu01",
				age: 18,
				score: 93.2,
			}
			TestStruct(a)
		}



////////
修改struct 字段值
		package main

		import (
			"reflect"
			"fmt"
		)

		type Student struct {
			Name string
			Age int
			Score float32
		}
		//定义struct 方法


		func TestStruct(a interface{})  {
			value := reflect.ValueOf(a)
			kind := value.Kind()
			// 传入 的参数是指针, 且指针指向的为struct
			if kind != reflect.Ptr && value.Elem().Kind() == reflect.Struct{
				fmt.Println("expect struct")
				return
			}
			//fmt.Println(value.NumField()) // 字段数量 3个
			value.Elem().Field(0).SetString("CCCCJJ")

		}

		func main()  {
			var a  = Student{
				Name:"stu01",
				Age: 18,
				Score: 93.2,
			}
			TestStruct(&a)
			fmt.Println(a)
		}