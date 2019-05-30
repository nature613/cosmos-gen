package main

import (
	"log"
	"os"
	"text/template"
)

type Module struct {
	Name string
}

func main() {
	m := Module{"Post"}

	// create a new template with data
	tmpl := template.New("module")
	tmpl, err := tmpl.Parse("Hello {{.Name}}")
	if err != nil {
		log.Fatal("Parse:", err)
		return
	}

	// merge template with contents of m
	err = tmpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}

	// fmt.Println("Hello")
	// f, err := os.Create("module.go")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// l, err := f.WriteString("Hello World")
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }
	// fmt.Println(l, "bytes written successfully")
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
