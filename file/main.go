package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err!=nil{
		fmt.Println("Error creating file:",err)
	}
	defer file.Close()
	fmt.Println("File Created:",file)
}