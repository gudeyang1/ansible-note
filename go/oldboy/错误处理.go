1. 定义错误
		package main

		import (
			"errors"
			"fmt"
		)

		var errNotFound = errors.New("Not found error")

		func main() {

			fmt.Printf("error, %s\n", errNotFound)
		}
