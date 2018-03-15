package main

import (
	"fmt"
	"os"
)

func main() {
	if token, err := token(os.Args); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%x\n", token)
	}
}
