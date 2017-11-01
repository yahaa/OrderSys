package main

import (
	"./apis"
	"time"
	"fmt"
)

func main() {
	t:=time.Now().Unix()
	fmt.Println(t)
	r := apis.Router
	r.Run(":7008")
}
