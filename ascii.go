package main

import "fmt"

func main(){
	codes := []int{104, 101, 108, 108, 111}
	var asciistring string
	for _, code := range codes{
		asciistring += string(code)

	}
	fmt.Println(asciistring)
}
