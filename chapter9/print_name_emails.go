package main


import (
	"html/template"
	"os"
	"fmt"
)


type Person struct {
	Name string
	Emails []string
}

const templ = `{{$name := .Name}}
{{range .Emails}}
	Name is {{$name}}, email is {{.}}
{{end}}
`


func main(){
	person := Person{
		Name: "j",
		Emails: []string{"j@nm.n", "j.nm@163.com"}
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}


