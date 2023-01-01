package main

import (
	"fmt"
	pb_api_gateway "github.com/f-robo/pb-api-gateway"
)

func main() {
	fmt.Println("main!")

	a := pb_api_gateway.GetApi()
	fmt.Println(a)
}
