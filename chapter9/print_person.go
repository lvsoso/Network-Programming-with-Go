package main



import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age int
	Emails []string
	Jobs []*Job
}


type Job struct {
	Employer string
	Role string
}

const templ = `The name is {{.Name}}.
The age is {{.Age}}.

{{range .Emails}} An email is {{.}} {{end}}

{{with .Jobs}}
	{{range .}}An employer is {{.Employer}} and the role is {{.Role}}{{end}}
{{end}}
`

func main(){
	job1 := Job{Employer: "Mo", Role: "Ho"}
	job2 := Job{Employer: "Box", Role: "Head"}

	person := Person{
		Name: "jan",
		Age: 50,
		Emails: []string{"j@163.com", "j.nm@qq.com"},
		Jobs: []*Job{&job1, &job2},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
