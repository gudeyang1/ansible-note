1. go 执行顺序
	 全局变量 --> init 函数- -> main 函数

	 包1 import 包2 , 包2 import 包3 : 执行顺序.. 
	 	包3中(全局变量 --> init 函数- -> main 函数) --> 
	 						包2中(全局变量 --> init 函数- -> main 函数) --> 
	 											包1中(全局变量 --> init 函数- -> main 函数)
 
 2.  不调用 import 的 package

	 import (
		_ "../second"
	)