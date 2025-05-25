package main

import (
	"Test/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Второй аргумент пустой, пожалуйста, передайте его")
		return
	}

	path := os.Args[1]

	proto, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return
	}

	var parser internal.Parse = &internal.Parser{}

	symbols := parser.ParseProto(string(proto))

	for _, s := range symbols {
		fmt.Println(s)
	}
}
