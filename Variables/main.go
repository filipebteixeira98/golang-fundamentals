package main

import "fmt"

func main() {
	// var phrase string = "Golang is awesome"
	// fmt.Println(keyword)

	// var (
	// 	character string = "G"
	// 	password  string = "s3cr3t"
	// )
	// fmt.Println(character, password)

	// word, quote := "Excited", "The quiter you become, the more you are able to hear"
	// fmt.Println(word, quote)

	// const dictation string = "The most men live lives in quiet desperation"
	// fmt.Println(dictation)

	change, value := "Swap", "Content"
	change, value = value, change
	fmt.Println(change, value)
}
