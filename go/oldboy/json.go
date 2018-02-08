1. json 序列化	
		package main

		import (

			"encoding/json"
			"fmt"
		)

		type User struct {
			UserName string `json:"user_name"`
			NickName string `json:"nick_name"`
			Age int `json:"age"`
			Birthday string `json:"birthday"`
			Sex string `json:"sex"`
			Email	string `json:"email"`
		}

		func main()  {
			user1 := &User{
				UserName:"user1",
				NickName:"cjk",
				Age:15,
				Birthday:"1126",
				Sex:"boy",
				Email:"11@qq.com",
			}

			data, err := json.Marshal(user1)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Printf("%s\n",string(data ))
		}

2.map 序列化
		package main

		import (
			"encoding/json"
			"fmt"
		)

		func testmap()  {
			var m map[string]interface{}
			m = make(map[string]interface{})

			m["username"] = "user2"
			m["age"] =10
			data, err := json.Marshal(m)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Printf("%s\n",string(data ))
			}

		func main()  {
			testmap()
		}

3. slice 嵌套map 序列化
		package main

		import (
			"encoding/json"
			"fmt"
		)

		func testmap()  {
			var s []map[string]interface{} //定义map 组成的slice
			var m map[string]interface{}  // 定义map
			m = make(map[string]interface{})

			m["username"] = "user2"
			m["age"] =10
			s = append(s,m)

			m["username"] = "user3"
			m["age"] =10
			s = append(s,m)

			m["username"] = "user4"
			m["age"] =10
			s = append(s,m)

			data, err := json.Marshal(s)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Printf("%s\n",string(data ))
			}

		func main()  {
			testmap()
		}

4. unmarshal  json to struct

		package main

		import (
			"encoding/json"
			"fmt"
		)

		type User struct {
			UserName string `json:"user_name"`
			NickName string `json:"nick_name"`
			Age int `json:"age"`
			Birthday string `json:"birthday"`
			Sex string `json:"sex"`
			Email	string `json:"email"`
		}

		func ToJson() (json_result string,err error) {  // 生成josn
			user1 := &User{
				UserName:"user1",
				NickName:"cjk",
				Age:15,
				Birthday:"1126",
				Sex:"boy",
				Email:"11@qq.com",
			}

			data, err := json.Marshal(user1)
			if err != nil{
				err = fmt.Errorf("json marshal failed,err:%s",err)
				return
			}
			json_result = string(data)
			return

		}

		func test()  {
			transed_json ,err:= ToJson()
			var user1 User
			if err != nil{
				fmt.Println("trans struct to json failed")
				return
			}
			err = json.Unmarshal([]byte(transed_json),&user1)  //将json 转换成string 赋值给user1 ,因为要修改uuser1的值,所有要传递指针类型
			if err != nil{
				fmt.Println("unmarshal failed, err:",err)
			}
			fmt.Println(user1)
		}

		func main()  {
			test()
		}

	输出:
		{user1 cjk 15 1126 boy 11@qq.com}

5. unmarshal json to map

		package main

		import (
			"encoding/json"
			"fmt"

		)

		func testmap() (result string,err error) {

			var m map[string]interface{}  // 定义map
			m = make(map[string]interface{})

			m["username"] = "user2"
			m["age"] =10

			data, err := json.Marshal(m)
			if err != nil{
				err = fmt.Errorf("marsh map to json failed,err:%s\n",err)
				return
			}
			result = string(data )
			return
		}

		func Unmarsh_Map()  {
			data, err := testmap()
			//fmt.Println(data)
			if err != nil{
				fmt.Println(err)
				return
			}
			var m map[string]interface{}
			//m = make(map[string]interface{})

			err = json.Unmarshal([]byte(data),&m)
			if err != nil{
				fmt.Println("unmarshal map to json failed, ",err)
				return
			}
			fmt.Println(m)

		}

		func main()  {
		Unmarsh_Map()
		}