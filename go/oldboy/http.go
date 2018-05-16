1. http server
		package main

		import (
			"net/http"
			"fmt"
		)

		func main()  {
			http.HandleFunc("/",Hello)   //匹配 /
			http.HandleFunc("/login",login)  //匹配http://xxx/login 到login 函数
			err := http.ListenAndServe("0.0.0.0:8880",nil)
			if err != nil{
				fmt.Println("http listen failed")
			}
		}

		func Hello(w http.ResponseWriter,r *http.Request)  {
			fmt.Println("handle hello")
			fmt.Fprintf(w,"hello" )
		}

		func login(w http.ResponseWriter,r *http.Request)  {
			fmt.Println("handle login")
			fmt.Fprintf(w,"login" )
		}
2. http 客户端
		package main

		import (
			"net/http"
			"fmt"

			"io/ioutil"
		)

		func main()  {
			res, err := http.Get("http://www.baidu.com")
			if err != nil{
				fmt.Println("get err:",err)
				return
			}
			data,err := ioutil.ReadAll(res.Body)
			if err != nil{
				fmt.Println("get data err:",err)
				return
			}
			fmt.Println(string(data))
		}

3. http 请求
3.1 HEAD
		package main

		import (
			"net/http"
			"fmt"
		)

		var url  = []string{
			"http://www.baidu.com",
			"http://google.com",
			"http://www.taobao.com",
		}
		func main()  {

			for _,v := range url{
				resp,err := http.Head(v)
				if err != nil{
					fmt.Printf("head %s failed,err: %v\n",v,err)
					continue
				}
				fmt.Printf("head success, status: %v" ,resp.Status)
			}
		}