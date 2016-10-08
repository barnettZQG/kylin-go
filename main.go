package main

import (
	"fmt"
	_ "kylin-orm/conf"
	"kylin-orm/kylin"
)

func main() {
	code, body, err := kylin.Login()
	if err != nil {
		panic(err)
	}
	fmt.Println(code, string(body))
}
