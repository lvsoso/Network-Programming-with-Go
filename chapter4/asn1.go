package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"time"
)

func main() {
	mdata, err := asn1.Marshal(13)
	checkError(err)

	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("After marshal/unmarshal int: ", n)

	s := "hello"
	smdata, _ := asn1.Marshal(s)

	var newstr string
	asn1.Unmarshal(smdata, &newstr)

	fmt.Println("After marshal/unmarshal string: ", smdata)

	t := time.Now()
	tmdata, err := asn1.Marshal(t)

	var newtime = new(time.Time)
	asn1.Unmarshal(tmdata, newtime)

	fmt.Println("After marshal/unmarshal time: ", tmdata)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
