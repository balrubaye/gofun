package main

import "fmt"

var p = fmt.Println

func main() {
	var str string
	p(len(str), str)

	str = "another nice string"
	p(len(str), str)

	str += "يس"
	p(len(str), str, str[len(str)-1])

}
