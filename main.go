package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	conf, err := loadConfig(filePath)
	if err != nil {
		printError(err)
		return
	}

	fmt.Println(conf)
}
