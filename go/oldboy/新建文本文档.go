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