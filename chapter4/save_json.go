package main



import (
	"encoding/json"
	"fmt"
	"os"
)



type Person struct {
	Name Name
	Email []Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}


func main(){
	person := Person {
		Name : Name {
				Family : "NM",
				Personal:"john"},
		Email: []Email{
			Email{Kind: "home", Address:"john@nm.name"},
			Email{Kind: "work", Address: "j.nm@163.com"}}}

	saveJSON("person.json", person)
}


func saveJSON(fileName string, key interface{}){
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}


func checkError(err error){
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

