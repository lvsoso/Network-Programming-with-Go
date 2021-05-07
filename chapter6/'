package main

import (
	"unicode/utf16"
)


func main(){
	str := "北京：Hello World !"

	println("String length", len([]rune(str)))
	println("Byte length", len(str))

	runes := utf16.Encode([]rune(str))
	ints := utf16.Decode(runes)

	str = string(ints)

	println(str)
}
