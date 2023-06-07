package main

import (
	"fmt"
	casbin "github.com/casbin/casbin/v2"
	"log"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Println("can")
	} else {
		fmt.Println("can not")
	}
}

func main() {
	e, err := casbin.NewEnforcer("./file-demo/model.conf", "./file-demo/policy.csv")
	if err != nil {
		log.Fatalln("new casbin failed:", err)
	}
	check(e, "jack", "data1", "read")
	check(e, "jack", "data2", "write")
	check(e, "tom", "data1", "read")
}
