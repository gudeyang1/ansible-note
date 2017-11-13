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

