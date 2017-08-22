package main

import (
	"fmt"
	"io"
	"text/template"
)

func main() {
	learningTemplating()
}

const status = 10

func templater() {

	b := Bug(400)
	b.startServer()
}

type Bug int

func (b *Bug) startServer() {
	i := int(*b)
	fmt.Println(i)
}

// do some http stuff

func learningTemplating() {

	type Choord map[string]int
	type Email []*map[string]string
	type Human struct {
		Name string
		Age  uint8
	}

	render("file.html")

}

func render(view string, data ...interface{}) (func(writer io.Writer), error) {
	const base_dir = "templates/"
	tmp, err := template.ParseGlob(base_dir + view)
	if err == nil {
		tmp.ParseFiles()
		return func(w io.Writer) {
			tmp.Execute(w, data)
		}, nil
	} else {
		fmt.Errorf(err.Error())
		return nil, err
	}
}
