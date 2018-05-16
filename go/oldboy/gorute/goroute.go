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
				sum *=i  //计算阶乘
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
			test = make(chan int,10)  // 需要make 初始化, 10 是元素个数
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

///////////////
channel 与goroute 简单配合示例
		package main

		import (
			"fmt"
			"time"
		)
	//写入0-9到channel
		func Write(ch chan int)  {
			for i:= 0 ;i<10; i++{
				ch <- i
			}
		}
	// 读取channel内容
		func Read(ch chan int)  {
		for {
			a := <- ch
			fmt.Println(a)
		}
		}

		func main()  {
			testChan := make(chan int)
			go Write(testChan)
			go Read(testChan)
			time.Sleep(time.Second )
		}
2.2.2 channel 阻塞
		package main

		import (
			"fmt"
			"time"
		)

		func Write(ch chan int)  {
			for i:= 0 ;i<10; i++{
				ch <- i
				fmt.Println(i)
			}
		}
		func Read(ch chan int)  {
		for {
			a := <- ch
			fmt.Println(a)
		}
		}

		func main()  {
			testChan := make(chan int,5)  //定义了channel的长度,会迅速写入5个进去,但是没有从channel里面读取,所有channel会阻塞,
			go Write(testChan)
			//go Read(testChan)
			time.Sleep(time.Second *10 )
		}

////////////////////////
判断100以为质数
		package main

		import (
			"fmt"
			"time"
		)

		func cacl(sourceChan chan int, resultChan chan int)  {
			for v:= range sourceChan{  //遍历channel内容
				flag := true
				for i :=2; i<v; i++{  // 判断是否为质数
					if v%i == 0{
						flag = false  //给质数打上flag
						break
					}
				}
				if flag{  //将质数写入result channel
					resultChan <- v
				}
			}
		}

		func main()  {
			intChan := make(chan int,100)
			resultChan := make(chan int,100)
			// 将1-100所有数塞入管道
			for i := 0 ; i<100; i ++{
				intChan <- i
			}
			//4核cpu 跑4个goroute
			for i:=0;i<4;i++{
				go cacl(intChan,resultChan)
			}
			//读取resultChannel 质数
			for i:= range resultChan{
				//i =<- resultChan
				fmt.Println(i)
			}
			time.Sleep(time.Second *10)
		}

//////////或者
		package main

		import (
			"fmt"
			"time"
		)

		func cacl(sourceChan chan int, resultChan chan int)  {
			for v:= range sourceChan{  //遍历channel内容
				flag := true
				for i :=2; i<v; i++{  // 判断是否为质数
					if v%i == 0{
						flag = false  //给质数打上flag
						break
					}
				}
				if flag{  //将质数写入result channel
					resultChan <- v
				}
			}
		}

		func main()  {
			intChan := make(chan int,100)
			resultChan := make(chan int,100)
			// 将1000000所有数塞入管道,每次读取4个,读完之后动态塞入剩下的数
			go func() {
				for i := 0 ; i<10000000; i ++{
					intChan <- i
				}
				close(intChan)
			}()

			//4核cpu 跑4个goroute
			for i:=0;i<4;i++{
				go cacl(intChan,resultChan)
			}
			//读取resultChannel 质数
			for i:= range resultChan{
				//i =<- resultChan
				fmt.Println(i)
			}
			time.Sleep(time.Second *10)
		}
///////
不使用sleep 退出channel

		package main

		import (
			"fmt"

		)

		func cacl(sourceChan chan int, resultChan chan int,exitChan chan bool)  {
			for v:= range sourceChan{  //遍历channel内容
				flag := true
				for i :=2; i<v; i++{  // 判断是否为质数
					if v%i == 0{
						flag = false  //给质数打上flag
						break
					}
				}
				if flag{  //将质数写入result channel
					resultChan <- v
				}
			}
			exitChan <- true
		}

		func main()  {
			intChan := make(chan int,100)
			resultChan := make(chan int,100)
			exitChan := make(chan bool,4)
			// 将1000000所有数塞入管道,每次读取4个,读完之后动态塞入剩下的数
			go func() {
				for i := 0 ; i<200; i ++{
					intChan <- i
				}
				close(intChan)
			}()

			//4核cpu 跑4个goroute
			for i:=0;i<4;i++{
				go cacl(intChan,resultChan,exitChan)
			}
			//添加exitchannel ,当协程干完活,会在exitChannel 里面写入"true",当写入满4个时关闭resultChannel,
			go func() {
				for i :=0;i<4 ;i++{
					<- exitChan
				}
				close(resultChan) // 关闭之后, for range 不会阻塞
			}()

			for i:= range resultChan{
				fmt.Println(i)
			}

		}


2.2.3 关闭channel

		package main

		import ("fmt")
		func main()  {

			ch := make(chan int,3)

			for i:= 0; i<3 ;i++{
				ch <- i
			}
			close(ch)

			for {
				b,ok := <- ch
				if ok == false{
					fmt.Println("channel closed")
					break
				}
				fmt.Println(b)
			}
		}
	输出:
		0
		1
		2
		channel closed

		///for range channel
		package main

		import (

			"fmt"

		)

		func main()  {

			ch := make(chan int,3)

			for i:= 0; i<3 ;i++{
				ch <- i
			}
			close(ch)  // 如果不关闭管道,for range 时会阻塞
			//for {
			//	b,ok := <- ch
			//	if ok == false{
			//		fmt.Println("channel closed")
			//		break
			//	}
			//	fmt.Println(b)
			//}
			for v := range ch{
				fmt.Println(v)
			}
		}


////////
示例:
		package main

		import (
			"fmt"
		)

		func send(ch chan int,exitChan chan bool)  {

			for i :=0;i<10;i ++{
				ch <- i
			}
			close(ch)
			exitChan <- true
		}

		func revice(ch chan int,exitChan chan bool)  {
			for {
				v,ok:= <- ch
				if !ok {
					break
				}
				fmt.Println(v)
			}
			exitChan <- true
		}

		func main()  {

			ch := make(chan int,10)
			exitCh:= make(chan bool,2)  // send recive 两个协程跑完之后在exitChan做标记
			go send(ch,exitCh)
			go revice(ch,exitCh)
			// 主程序需要等待 exitchannel 显示sent recive 两个协程跑完毕
			for i:=0;i<2;i++{
				<- exitCh
			}

		}

2.2.4 定时器
// 每隔一秒输出一次, 无限循环
		package main

		import (
			"time"
			"fmt"
		)

		func main()  {
			t := time.NewTicker(time.Second)

			for v:= range t.C{  //	C是个channel
				fmt.Println("hello",v)
			}
			t.stop()// 关闭定时器
		}

// 一次性定时器

		package main

		import (
			"time"
			"fmt"
		)

		func main()  {

			ch := make(chan int,10)
			ch2 := make(chan int,10)
		go func() {
			var i int
			for{
				ch <- i
				time.Sleep(time.Second)
				ch2 <- i*i
				time.Sleep(time.Second)
				i++
			}
		}()

		for {
			select {
			case v := <- ch:
				fmt.Println(v)
			case v := <- ch2:
				fmt.Println(v)
			case <- time.After(time.Second):  //超时
				fmt.Println("time out")
				return
			}
			}
		}


2.2.6  捕获goroute panic

		package main

		import (
			"fmt"
			"time"
		)

		func test()  {
			defer func() {
				if err := recover(); err != nil{
					//等于 两条语句分开写
					// err := recover()
					// if err != nil
					fmt.Println("panic error")
				}
			}()
			var m map[string]int

			m["stu"] = 100
		}

		func working()  {
			for {
				fmt.Println("i'm working")
				time.Sleep(time.Second)
			}
		}

		func main()  {

			go test()   // panic 不会影响下面代码运行
			go working()
			time.Sleep(time.Second *10)
		}
	输出:
	i'm working
	panic error
	i'm working
	i'm working