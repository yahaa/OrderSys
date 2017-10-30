package main

import (
	"./apis"
)

func main() {
	r := apis.Router
	r.Run(":7008")
}
