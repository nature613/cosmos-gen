package main

import (
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/jinzhu/inflection"
)

// Module specifies the data needed to construct a module
type Module struct {
	Name string
}

func main() {
	funcMap := template.FuncMap{
		"title":     strings.Title,
		"pluralize": inflection.Plural,
	}

	name := os.Args[1]
	m := Module{name}

	box := packr.NewBox("./templates")

	files := []string{"codec.go", "genesis.go", "module.go", "handler.go", "keeper.go", "querier.go", "types.go"}

	for _, moduleFilename := range files {
		moduleTemplateFilename := moduleFilename + ".tmpl"
		filename, err := box.FindString(moduleTemplateFilename)
		if err != nil {
			log.Fatal(err)
			return
		}
		t, err := template.New(moduleTemplateFilename).Funcs(funcMap).Parse(filename)
		if err != nil {
			log.Fatal(err)
			return
		}

		modulePath := path.Join("x", name)
		err = os.MkdirAll(modulePath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}

		f, err := os.Create(path.Join(modulePath, moduleFilename))
		if err != nil {
			log.Fatal(err)
			return
		}

		err = t.Execute(f, m)
		if err != nil {
			log.Fatal("Execute: ", err)
			return
		}

		err = f.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}
}
