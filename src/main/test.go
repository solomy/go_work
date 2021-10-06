package main

import (
	"fmt"
	"main/myhelloworld"
	"time"
)

func main() {
	fmt.Println("你好，世界！")
	myhelloworld.Hello()
	time.Sleep(5 * time.Second)
}
