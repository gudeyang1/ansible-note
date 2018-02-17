1.   1.8 以前go设置程序跑的内核数量

		package main

		import (
			"fmt"
			"runtime"
		)

		func main()  {

			num := runtime.NumCPU()
			runtime.GOMAXPROCS(num)
			fmt.Println(num)
		}


2. 不通gotoutine r如何进行通讯

2.1 全局变量和锁同步
		package main

		import (
			"sync"
			"time"
			"fmt"
		)

		var m = make(map[int]int)
		var lock  sync.Mutex

		type task struct {
			n int
		}

		func cacl(t *task)  {
			var sum int
			sum = 1
			for i := 1; i<t.n; i++{
				sum *=i
			}
			lock.Lock()
			m[t.n] = sum
			lock.Unlock()
		}

		func main()  {
			for i :=0; i <10; i++{
				t := &task{n:i}
				go cacl(t)  // 不会等待这行代码返回,直接执行下面代码, 这行代码多线程同时执行
			time.Sleep(time.Second * 2)

			lock.Lock()
			for k, v := range m{
				fmt.Printf("%d!=%v\n",k,v)
			}
			lock.Unlock()

		}

2.2 channel

2.2.1 channel的声明
var 变量名 chan 类型
var test chan int
var test chan string
var test chan map[string]string

		package main

		import "fmt"

		func main()  {

			var test chan int
			test = make(chan int,10)  // 需要make 初始化, 10 是cap
			test <- 10  // 写入channel
			a := <- test  // 从channel 读取
			fmt.Println(a)
		}

////////////////////////////////
复杂类型channel 
		package main

		import (
			"fmt"

		)

		type student struct {
			name string
		}
		func main()  {

			var stuChan chan interface{}
			stuChan = make(chan interface{},10)   // channel 是interface 类型

			stu := &student{name:"jim"}  // 生成struct

			stuChan <- stu  // 给channel 传值
			
			// 读取channel的内容
			var  stu01 interface{}   
			stu01 = <- stuChan
			//fmt.Println(stu01)

			//从interface 转成成具体类型(指针类型的struct)
			var stu02 *student
			 stu02, ok := stu01.(*student)

			 if !ok {
			 	fmt.Println("convert interface{} to struct failed")
			 }
			 fmt.Println(stu02)

		}