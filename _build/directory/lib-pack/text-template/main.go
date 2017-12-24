package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	const tpl = `

` // A

	type Tree struct {
		Name     string
		Children []Tree
	}
	root := Tree{
		"A", []Tree{
			Tree{"B", []Tree{
				Tree{"C", []Tree{}},
			}},
			Tree{"D", []Tree{}},
		},
	}

	t := template.Must(template.New("tree").Parse(tpl))

	err := t.Execute(os.Stdout, root)
	if err != nil {
		log.Fatalf("executing template:", err)
	}

}
