package main

import (
	"fmt"

	v1 "github.com/zhanjingjie/buf-testing/protos/gen/go/zhanjingjie/buftesting/buftesting/v1"
)

func main() {
	fmt.Println("hello world")

	_ = v1.Echo{}
}
