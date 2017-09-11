0.
	printf 是和标准输出文件(stdout)关联的,fprintf 则没有这个限制.
	 
	fprintf是用于文件操作的，原型是int fprintf( FILE *stream, const char *format [, argument ]...);

	sprintf是格式化输出到一个字符串，fprintf是格式化输出到一个stream，通常是到文件。


1. 变量默认值
	var  (
		a int   // 默认 0
		b string  // 默认 空""
		c bool  // 默认false
		d = 8
	)

2. 值类型和引用类型
	2.1 值类型
		变量直接存值
		int float bool string struct  
	2.2 引用类型
		变量存的是一个地址,这个地址存的是最终的值
		slice  map chan ,指针 


		func main() {
			a := 5
			b := make(chan int ,1)

			fmt.Println(a)
			fmt.Println(b)
			modify(&a)   // 传递内存地址
			fmt.Println(a)
		}

		func modify(a *int)   {   //*  表示指针
			*a = 10    // 给指定赋值
			return
		 }

3. 类型转换
	var a int8 = 10
	var b int32 = int32(a)  //将a 转换成int32 类型
	int8 int32 虽然都是整数类型,但是不是同一种类型, 不能做运算

4. 生成随机数
	func main()  {
		rand.Seed(time.Now().Unix())  // 获取当前秒时间
		for i:=0; i< 10 ;i++{
			fmt.Println(rand.Intn(10))
		}

		for i := 0 ;i < 10; i++ {
			fmt.Println(rand.Intn(100))
		}

		for i:= 0; i< 10; i++ {
			fmt.Println(rand.Float32())
		}
	}


5. 读取标准输入
		var a int
		var b int
		fmt.Scanf("%d %d", &a, &b)   // 必须要取内存地址
		fmt.Println(a, b)

6. 水仙花判断
		func main() {

			for i := 100 ; i< 999; i++{
				if isnumber(i) == true{
					fmt.Println(i)
				}
			}
		}
		func isnumber(n int) bool{
			var a,b,c  int

			a = n/100
			b = (n%100) /10
			c = n%10
			sum := a*a*a + b*b*b + c*c*c
			return sum == n
		}

7. 求阶乘 !5
		func main() {

			var a int
			fmt.Scanf("%d", &a)
			fmt.Println(sum(a))

		}

		func sum(n int) uint64 {
			var s uint64 = 1
			var all_sum uint64  = 0
			
			for i := 1 ; i< n; i++{
				s = s * uint64(i)
				all_sum += s
			}
			return all_sum
		}


8. 时间类型
		func main() {
			now := time.Now()
			fmt.Println(now.Format("2006/01/02 15:04:05"))  // 时间格式必须是 2006/01/02 15:04
		}

		func main() {
		start := time.Now().UnixNano()  // 获取1970到现在的纳秒数
		now := time.Now()
		fmt.Println(now.Format("2006/01/02 15:04:06"))
		end := time.Now().UnixNano()  

		fmt.Printf("cost: %d \n", (end - start)/1000/1000)  // 程序执行完毕消耗的毫秒数
		}

9. 指针

		func main() {
			var a string = "ss"

			fmt.Println(&a)   // 取a 的内存地址

			var p *string = &a   // p 为指针类型, 指向的是a 的内存地址
			*p = "ppppp"	// *p 为p 的内存地址对应的值
			
			fmt.Println(p)
			fmt.Println(a)
		}

		输出:
		0xc042008240   // a 的内存地址
		ppppp    // *p 对应的值
		ppppp   // a内存地址对应的值被p 修改, a 的值就会修改

10. 转换
		func main() {
			var a string
			fmt.Scanf("%s", &a)
			convert(a)
		}

		func convert(n string)   {

			num , err := strconv.Atoi(n)
			if err != nil{
				fmt.Printf("%s can not covert to int \n", n)
				return   // 报错就返回, 不在继续
			}
			fmt.Println(num)
		}
11. case 语句

		func main() {
			var a int = 10

			switch a {
			case 1, 10:
				fmt.Println("1 or 10")
			case 2:
				fmt.Println("2")
			default:
				fmt.Println("default")

			}
		}		

		// 多条件
		func main() {
		var a int = 101
		switch  {
		case a >10 && a < 100:
		fmt.Println("sss")
		case a >100:
			fmt.Println("dayu")
		case a < 10:
			fmt.Println("xiao yu")

		default:
			fmt.Println("default")
		}
	}
12. type 
		type  add_func func(int,int) int //两个参数必须要用, 分隔

		func main() {
			//c := add  // c 是内存地址
			fmt.Println(operator(add, 200,300))
		}
		func add(a,b int) int{
			return a + b
		}

		func operator(op add_func, a,b int) int  {  // a b 对应的是type 里面定义的两个 int 参数
			return op(a,b)
		}

13. 可变长参数
		func Add(arg ... int) int  {    // 传入的arg参数是个slice [1,2,3,4,....]
			sum := 0
			for i := 0; i< len(arg); i++{
				sum += arg[i]
			}
			return sum
		}

14.  defer
	defer 是先进后出

		func main() {

			i := 1
			defer fmt.Println(i)  // 先保存到栈里面,函数退出之前再执行 , i 是值拷贝, 下面对i 的修改不会影响defer 函数里面的i 的值
			defer fmt.printf("second")
			i = 10
			fmt.Println(i)   

		}
		输出:
		10
		second
		1

	defer 的用途:
		1. 关闭文件句柄
		2. 锁资源释放
		3. 数据库链接的释放

15. 匿名函数
		func main() {

			i := 1
			defer func() {   // 函数没有名字
				fmt.Println("xxxxx")
			}()   // 结尾加上() 表示调用匿名函数

			i = 10
			fmt.Println(i)
		}			
		输出:
		10
		xxxxx	
///////////////////////////////
		func main() {
			a := test(100,200)
			fmt.Println(a)
		}

		func test(a, b int) int  {
			result := func(a1,b1 int)int {
				return a1 + b1
			}(a,b)   // 匿名函数调用传值
			return result
		}
======或者========================================
		func test(a, b int) int  {
			result := func(a1,b1 int)int {
				return a1 + b1
			}    // 这里不调用

			return result(a,b)   
		}

16. 中文字符问题

		[]byte(string) 是字节, 一个中文字符占三个字节,
		[]rune(string) 是字符, 一个字符就是一个字符,不管是中文还是英文
				func main() {

				str :="laoYu老虞"
				r := []rune(str)    //中文字符需要转换成rune 类型之后才可以计算
				fmt.Println("rune=",r)
				for i:=0;i<len(r) ; i++ {
				       fmt.Println("r[",i,"]=",r[i],"string=",string(r[i]))
					fmt.Println(len(string(r[i])))
				}
				}


		func main() {

			reader := bufio.NewReader(os.Stdin)    //从输入读取的新方式
			name, _,err :=reader.ReadLine()			//
			if err != nil{
				fmt.Println("eror happens")
			}
			fmt.Println(name)    // 输出, [115 97 99]   字符的utf-8编码
			fmt.Printf("%c\n" , name)   // 格式化, 但是中文会失败 [s a c]
			fmt.Println(string(name))   //  将输入的[]byte 类型转换后才能string 类型   sac
			t := []rune(string(name))   // 将输入转成字符类型, 尤其针对中文, 注意要将输入的byte类型转换后才能string 作为其参数
			fmt.Println(t)    // 输出: [115 97 99]
			
				for _,v := range t{
				fmt.Println(string(v))   // 将rune 类型转换成string 即可输出输入 的字符串
			}
		}
		输出:
		sac   // 输入
		[115 97 99]
		[s a c]
		sac
		[115 97 99]



17. 内置函数

	17.1 new
	new 用来分配内存, 主要用来分配值类型, 比如int, struct ,返回的是指针
		func main() {
		var name int
		fmt.Println(name)

		new_name := new(int)
		fmt.Println(new_name)   //输出内存地址
		/*
		new_name := new(int)
		*new_name = 12     // 赋值
		fmt.Println(*new_name)		
		*/
		}
		输出:
		0
		0xc042008260    // new 输出的是指针
	17.2 make
	make 用来分配内存, 只要用来分配引用类型, 比如chan ,map, slice

	17.3 append
		var a  []int

		a  = append(a,10,22,34)
		fmt.Println(a)
		a = append(a,a...)   // ...表示数组a的展开
		fmt.Println(a)
	输出:
		[10 22 34]
		[10 22 34 10 22 34]	
	17.4 panic
		func test()  {
			
			defer func() {     //捕获异常
				if err := recover()	; err != nil{
					fmt.Println(err)  //打印报错
					debug.PrintStack()   //打印堆栈信息, 具体哪行代码出错
				}
			}() //()调用匿名函数

			var a int = 0
			b := 100 /a
			fmt.Println(b)
		}
//////自动触发panic
		func main() {
			err := InitConfig()
			if err != nil{
				panic(err)
			}

		}

		func InitConfig()error  {
			return errors.New("init config error")

		}


18. 递归 , 一个函数自己调用自己交递归
		func main() {
			recusive(9)

		}
 
		func recusive(n int)  {   // 添加n作为退出死循环条件  

			fmt.Println("hello" )
			time.Sleep(time.Second)
			if n > 10{   // 定义出口条件
				return
			}
			recusive(n + 1)

		}

19. 闭包
		func main() {
			f:= adder()
			fmt.Println(f(1))
			fmt.Println(f(200))
			fmt.Println(f(100))

		}

		func adder() func(int) int {
			var x int   //x 的值会保存起来
			return func(d int)int {
				x += d
				return x
			}
		}
	输出:
		1
		201
		301

///////////////////////////////////////////////////
	func makeSuffix(suffix string)func(string) string  {

		return func(name string) string {
			if strings.HasSuffix(name,suffix) == false{
				return name + suffix
			}
			return name
		}

	}


	func main() {
		f1 := makeSuffix(".bmp")
		fmt.Println(f1("test"))

		f2 := makeSuffix(".jpg")
		fmt.Println(f2("pic"))

	}
	输出:
	test.bmp
	pic.jpg

20. 数组
	20.1 定义
		func main() {
			var a = [...]int{3:3}   // 数组有长度, slice 没有长度
			var b = []int32    // 这个是切片
			fmt.Println(a)

		}
		//
		[0 0 0 3]

	20.2 多维数组
	func main() {
		var a  = [...][3]int{{1,2,3},{3,4,5}}

		for k,v := range a{
			for index, value := range v{
				fmt.Printf("a[%d][%d] = %d", k,index,value)
				fmt.Println()
			}
		}

	}
	输出:
	a[0][0] = 1
	a[0][1] = 2
	a[0][2] = 3
	a[1][0] = 3
	a[1][1] = 4
	a[1][2] = 5

21. 切片
	var slice []int   //切片
	var arr =  [5]int{1,2,3,4,5}   // 数组
 	var b [5]int  = [...]int{1,2,3,4,5}   //数组
 	var b [5]int  = [5]int{1,2,3,4,5}   //数组



 23. 排序

	23.1 整数排序
	func main() {
		list := []int{1,2,3,434,65,32,12}   //切片

		sort.Ints(list)   // 需要传递切片
		fmt.Println(list)
	}

	23.2 字符串排序
	func main() {
		list := []string{"asa","scsac","sc","a"}
		sort.Strings(list)
		fmt.Println(list)
	}

	23.3 排序搜索

	func main() {
		list := []int{23,43,12,43}
		sort.Ints(list)
		fmt.Println(list)
		fmt.Println(sort.SearchInts(list, 43))   // 返回索引 2   排序后搜索必须是排好序的切片
	}

	输出:
	[12 23 43 43]
2

24. map
	24.1 定义
	func TestMap()  {
		
		a := make(map[string]string,10)   // map[key 类型]value 类型

		a["a"] = "21"
		fmt.Println(a)
	}
	输出:
	map[a:21]
	// 或者
	func TestMap()  {
		a := map[string]string{
			"a":"1",
			"b":"2"}
		fmt.Println(a)
	}

	24.2 map 嵌套map
	func TestMap()  {
		a := make(map[string]map[string]string,10)
		a["1"] = make(map[string]string)   // 需要对a[1] 的map make 分配内存
		a["1"]["a"] = "sss"
		a["1"]["b"] = "ssssss"
		a["1"]["c"] = "sssssssss"
		a["1"]["d"] = "sssqqsss"
		fmt.Println(a)
	}
	输出:
	map[1:map[a:sss b:ssssss c:sssssssss d:sssqqsss]]

	//示例:

	func TestMap()  {
		a := make(map[string]map[string]string)
		
		_,v := a["zhangsan"]
		if !v{
			a["zhangsan"] = make(map[string]string)   //如果不存在, 就初始化, 如果存在就更新记录

		}
		a["zhangsan"]["passowr"] = "123"
		a["zhangsan"]["nickname"] = "jack"
		fmt.Println(a)
	}

25. 包
线程同步 , 
build 加上 --race  . go build --race main.go

import sync
 互斥锁: sync.Mutex   一次只能有一个人干活
 读写锁: sync.RWMutex  无线读,一个写
		 var lock sync.Mutex
		func TestMap()  {

			a := make(map[int]int,4)

			a[1] = 1
			a[2] = 2
			a[3] = 3
			a[4] = 4
			for i := 0; i < 2 ;i ++{
				go func(b map[int]int) {
					lock.Lock()
					b[1] = rand.Intn(100)
					lock.Unlock()
				}(a)
			}
			lock.Lock()
			fmt.Println(a)
			lock.Unlock()
			time.Sleep(time.Second)
		}
		// 读写锁
		var rwlock sync.RWMutex
		func TestMap()  {

			a := make(map[int]int,4)

			a[1] = 1
			a[2] = 2
			a[3] = 3
			a[4] = 4
			for i := 0; i < 2 ;i ++{
				go func(b map[int]int) { 
					rwlock.Lock()    // 写锁
					b[1] = rand.Intn(100)
					rwlock.Unlock()
				}(a)
			}
			for i := 0; i < 100 ;i ++{
				go func(b map[int]int) {
					rwlock.RLock()   //读锁
					fmt.Println(a)
					rwlock.RUnlock()
				}(a)
			}
			time.Sleep(time.Second * 3)

		}


		测试性能, 3秒内读取次数
		var lock sync.Mutex
		var rwlock sync.RWMutex
		func TestMap()  {
			var conut int32
			a := make(map[int]int,4)

			a[1] = 1
			a[2] = 2
			a[3] = 3
			a[4] = 4
			for i := 0; i < 2 ;i ++{
				go func(b map[int]int) {
					rwlock.Lock()
					b[1] = rand.Intn(100)
					rwlock.Unlock()
				}(a)
			}

			for i := 0; i < 100 ;i ++{
				go func(b map[int]int) {
					for  {
						rwlock.RLock()
						time.Sleep(time.Millisecond)  //一毫秒
						rwlock.RUnlock()
						atomic.AddInt32(&conut,1)   // 原子性操作
					}
				}(a)
			}
			time.Sleep(time.Second * 3)
			fmt.Println(atomic.LoadInt32(&conut)) 
		}
