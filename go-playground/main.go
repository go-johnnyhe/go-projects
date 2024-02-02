package main

import (
"runtime"
"fmt" 
"github.com/go-johnnyhe/go-projects/go-playground/versionPack"
)
func main() {
	fmt.Println("I am happy today:)")
	fmt.Println("hahaha")
	fmt.Println("yay")

	fmt.Println(runtime.Version())

	versionPack.Version()
}
