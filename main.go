package main

import (
	"fmt"

	"github.com/DecodeWorms/cicd-example/greeting"
)

func main() {
	res := greeting.Greeting()
	fmt.Println(res)
}
