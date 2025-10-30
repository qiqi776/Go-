package main

import (
	"GoStudy/routers"
	"fmt"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		fmt.Println("err:", err)
	}
}
