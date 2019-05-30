package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	f, err := os.Create("module.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
