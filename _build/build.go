package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Quiz *struct {
	Key  string `yaml:"key"`
	Name string `yaml:"name"`

	Question string
	Code     string
	Main     string
	Answer   string
}

type Sections struct {
	Sections []struct {
		Key     string `yaml:"key"`
		Name    string `yaml:"name"`
		Quizzes []Quiz `yaml:"quizzes"`
	} `yaml:"sections"`
}

func loadDirectory(directory string) *Sections {
	bs, err := ioutil.ReadFile(directory)
	if err != nil {
		log.Fatalf("Loading %s: %v", directory, err)
	}

	sections := Sections{}
	err = yaml.Unmarshal(bs, &sections)
	if err != nil {
		log.Fatalf("Parsing directory.yaml: %v", err)
	}
	return &sections
}

func buildHTML(directory, tmpl string) {
	sections := *loadDirectory(directory)
	for _, section := range sections.Sections {
		for _, quiz := range section.Quizzes {
			var bs []byte
			prefix := "directory/" + section.Key + "/" + quiz.Key
			bs, _ = ioutil.ReadFile(prefix + "/question.txt")
			quiz.Question = string(bs)
			bs, _ = ioutil.ReadFile(prefix + "/code.txt")
			quiz.Code = string(bs)
			bs, _ = ioutil.ReadFile(prefix + "/main.go")
			quiz.Main = string(bs)
			bs, _ = ioutil.ReadFile(prefix + "/answer.txt")
			quiz.Answer = string(bs)
		}
	}

	bs, err := ioutil.ReadFile(tmpl)
	if err != nil {
		log.Fatalf("Loading %s: %v", tmpl, err)
	}
	t := template.Must(template.New("directory").Parse(string(bs)))
	if err = t.Execute(os.Stdout, sections); err != nil {
		fmt.Fprintln(os.Stderr, "%v", err)
	}
}

func buildREADME(directory string) {
	sections := *loadDirectory(directory)
	for _, section := range sections.Sections {
		fmt.Fprintf(os.Stdout, "[%s](https://goquiz.github.io/#%s)\n", section.Name, section.Key)
		for _, quiz := range section.Quizzes {
			fmt.Fprintf(os.Stdout, "- [%s](https://goquiz.github.io/#%s)\n", quiz.Name, quiz.Key)
		}
		fmt.Fprintln(os.Stdout)
	}
}

func main() {
	directory := flag.String("d", "directory.yaml", "directory.yaml")
	tmpl := flag.String("t", "index.tmpl", "index.tmpl")
	cmd := flag.String("cmd", "html", "html|readme")

	flag.Parse()
	switch *cmd {
	case "html":
		buildHTML(*directory, *tmpl)
	case "readme":
		buildREADME(*directory)
	default:
		log.Fatalf("Please specify a valid cmd (e.g. html|readme)")
	}
}
