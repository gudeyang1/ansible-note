package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main()  {
	a := "999"
	b := "999"

	bigdatasum(a,b)

}

func bigdatasum(a,b string) (string ,string)  {
	// 位数少的, 高位补0
	if len(a) < len(b){
		a = strings.Repeat("0",len(b)- len(a)) +a
	}else {
		b = strings.Repeat("0", len(a)- len(b)) +b
	}
	fmt.Println(a)
	fmt.Println(b)

	addOne := false
	sum_list := make([]int, len(a))

	for i :=0; i< len(a);i++{
		var numa int  // 如果用num:= numa 是byte 类型 不方便转换成string类型
		var numb int
		numa = int(a[len(a) -i -1] - '0')
		numb = int(b[len(a) -i -1] - '0')
		per_sum := numa + numb

		// 高位进1
		if addOne{
			per_sum += 1
		}
	// 各位超过10 就减去10
		if per_sum > 9{
			per_sum -= 10
			addOne = true
		}else{
			addOne = false
		}
		sum_list[i] = per_sum   // 得出的list 是反的, 和是123的话, 列表里面存的是321
	}

	result := convertToSting(sum_list)

	if addOne == true{   // 最后的高位进1
		result = "1" + convertToSting(sum_list)
	}

	fmt.Println(result)
	return a,b
}

func convertToSting(a []int )  string {
	new_list := make([]string, len(a) )
	//new_list := []string{}   // 这样会报index out of range 错误
	for i :=0 ;i<len(a);i ++{
		new_list[i] =  strconv.Itoa(a[len(a) -i -1])
	}

	return strings.Join(new_list,"")
}