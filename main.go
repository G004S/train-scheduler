package main

import "fmt"

func main() {

	//Validates setup: args, path to .map
	mapPath, err := ValidateSetup()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Get content from .map
	mapContent, err := GetMapContent(mapPath)
	fmt.Println(mapContent)

}
