package main

import "fmt"

/*
Go Commands:
	go build: compiles the file and creates the executable.
	go run: compiles and runs the file.
	go fmt: format all the code in all the files in the project.
	go install: compiles and installs a package.
	go get: gets the raw code of someone's else package.
	go test: runs all the test in the project.


Go Packages:
	Go have two types of packages 1. Exeuctable and 2. Reusable
	main is the executable package which must have function with name main and package with any other name is reusable package.
*/

var cards int

func main() {
	isValid := true
	var times int
	times = 10
	cards = 52
	fmt.Println("Hello world")
	fmt.Println(isValid)
	fmt.Println(times)
	fmt.Println(cards)
}
