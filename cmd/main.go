package main

import (
	"Test/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		_ = fmt.Errorf("Ошибка! Аргумент не передан")
		return
	}

	if len(os.Args) != 2 {
		_ = fmt.Errorf("Ошибка! Аргумент должен быть 1")
		return
	}

	path := os.Args[1]

	proto, err := os.ReadFile(path)
	if err != nil {
		_ = fmt.Errorf("Ошибка чтения файла: %v\n", err)
		return
	}

	var parser internal.Parse = &internal.Parser{}

	symbols := parser.ParseProto(string(proto))

	for _, s := range symbols {
		fmt.Println(s)
	}
}
