package main

import (
	"fmt"
)

func hello(ss []string) {
	fmt.Println("hello! your cmds are: ")
	for _, s := range ss {
		fmt.Println(s)
	}
}
