package main

import (
	"Test/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("Ошибка: вы должны предоставить ровно один файл в качестве аргумента, вы предоставили %d\n", len(os.Args)-1))
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
