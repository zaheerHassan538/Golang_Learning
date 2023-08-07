package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("js shorts for :", languages["js"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	//loops are intersting in golang

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}

}
