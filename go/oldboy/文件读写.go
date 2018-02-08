1. 格式化字符串

		package main

		import "fmt"

		type Student struct {
			Name string
			Age int
			Score float32
		}


		func main()  {
			var str = "stu01 18 98.22"
			var stu Student
			fmt.Sscanf(str, "%s %d %f", &stu.Name,&stu.Age, &stu.Score ) // 格式化str 输出到stu
			fmt.Println(stu)
		}

2. 带缓存区的读
		package main

		import (
			"bufio"
			"os"
			"fmt"
		)

		func main()  {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("please input something:")
			str ,err := reader.ReadString('\n') //传递的是byte类型

			if err != nil{
				fmt.Println("read string failed")
				return
			}

			fmt.Println(str)
		}


3. 读写文件
		package main

		import (
			"bufio"
			"os"
			"fmt"
		)

		func main()  {

			file,err1:= os.Open("C:/test/test.log")
			if err1 != nil{
				fmt.Println("read file error:",err1)
				return
			}
			defer file.Close()  //关闭文件
			reader := bufio.NewReader(file)
			str ,err := reader.ReadString('\n')
			if err != nil{
				fmt.Println("read string failed",err)
				return
			}

			fmt.Println(str)
		}

4. 统计文件中字母,数字, 空格及其它字符的数量

		package main

		import (
			"bufio"
			"os"
			"fmt"
			"io"

		)

		type Count struct {
			ChrCount int
			SpaceConut int
			NumConut int
			OtherConut int
		}


		func main()  {
			var count Count

			file,err1:= os.Open("C:/test/test.log")
			if err1 != nil{
				fmt.Println("read file error:",err1)
				return
			}
			defer file.Close() //一定要关闭文件

			reader := bufio.NewReader(file)
			for {   

				str ,err := reader.ReadString('\n')   //读取一行统计一行
				if err == io.EOF{ //如果是行尾, 跳出循环
					break
				}
				if err != nil{
					fmt.Println("read string failed",err)
					return
				}

				
				RuneArry := []rune(str) // 解决中文字符问题
				
				for _,v := range RuneArry{
					switch  {
					case v >= 'a' && v <= 'z':
						fallthrough
					case v >='A' && v <= 'Z':
						count.ChrCount ++
					case v == ' '|| v == '\t':
						count.SpaceConut ++
					case v >='0' && v <='9':
						count.NumConut ++
					default:
						count.OtherConut ++

					}
				}

			}

			fmt.Printf("char num is %d\n", count.ChrCount)
			fmt.Printf("num count:%d\n",count.NumConut)
			fmt.Printf("space conut:%d\n",count.SpaceConut)
			fmt.Printf("other conut:%d\n",count.OtherConut)
		}

5. 读取整个文件
		package main

		import (
			"io/ioutil"
			"fmt"
			"os"
		)

		func main()  {
			inputfile := "C:/test/test.log"
			outputfile := "C:/test/out.log"

			buf ,err := ioutil.ReadFile(inputfile)
			if err != nil {
				fmt.Fprintf(os.Stderr,"file error: %s",err)
				return
			}
			fmt.Printf("%s\n",string(buf))

			err = ioutil.WriteFile(outputfile,buf,0664) //权限0664
			if err != nil{
				panic(err.Error())
			}
		}

6. 读取压缩文件
		package main

		import (
			"bufio"
			"os"
			"fmt"
			"compress/gzip"
			"io"
		)

		func main()  {
			fName := "C:/test/test.log.gz"
			var r *bufio.Reader
			fi , err := os.Open(fName)
			defer fi.Close()
			if err != nil{
				fmt.Printf("%v ,can't open %s: error: %s\n",os.Args[0],fName,err)
				os.Exit(1)
			}

			fz, err := gzip.NewReader(fi)
			if err != nil{
				fmt.Printf("open gzip failed, err: %v\n",err )
				return
			}
			r = bufio.NewReader(fz)
			for{
				line ,err := r.ReadString('\n')
				if err == io.EOF{
					break
				}
				if err != nil{
					fmt.Println("done read file")
					os.Exit(1)
				}
				fmt.Println(line)
			}
		}


7. 文件写入
os.Open(fName)   // 打开文件只能读
os.OpenFile(“output.dat”,  os.O_WRONLY|os.O_CREATE, 0666)
第二个参数：文件打开模式：
1. os.O_WRONLY：只写
2. os.O_CREATE：创建文件
3. os.O_RDONLY：只读
4. os.O_RDWR：读写
5. os.O_TRUNC ：清空


		package main

		import (
			"os"
			"fmt"
			"bufio"
		)

		func main()  {
			outputFile, outputError:= os.OpenFile("C:/test/test.log",os.O_WRONLY|os.O_CREATE,0644)  
			if outputError != nil{
				fmt.Printf("an error occurred with file creation\n")
				return
			}
			defer outputFile.Close()

			outputWriter := bufio.NewWriter(outputFile)
			outputString := "hello world\n"
			for i :=0; i <10 ; i++{
				outputWriter.WriteString(outputString)
			}
			outputWriter.Flush()  //刷新内存数据到磁盘

		}

8. 拷贝文件
		package main

		import (
			"os"
			"io"
		)

		func CopyFile(destName, srcName string)(written int64,err error)  {
			src,err := os.Open(srcName)
			if err != nil{
				return
			}
			defer src.Close()

			dest, err := os.OpenFile(destName,os.O_CREATE|os.O_WRONLY,0644)
			if err != nil{
				return
			}
			defer dest.Close()
			return io.Copy(dest,src)
		}
		func main()  {
			CopyFile("C:/test/test_copy.log","C:/test/test.log")
		}

9. 命令行参数 flag
		package main

		import (
			"flag"
			"fmt"
		)

		func main()  {

			var confPath string
			var LogLevel int

			flag.StringVar(&confPath,"c","","path of the config file")
			flag.IntVar(&LogLevel,"d",3,"log level")
			flag.Parse()   //使 flag 生效

			fmt.Printf("conf path: %s\n",confPath)
			fmt.Printf("log level:%d\n",LogLevel)
		}