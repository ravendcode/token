package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

const maxCount = 1000

var count = 32
var err error

func token(args []string) ([]byte, error) {
	if len(args) == 2 {
		count, err = strconv.Atoi(args[1])
		if err != nil {
			return nil, fmt.Errorf("аргумент должен быть числом")
		}
		if count > maxCount {
			return nil, fmt.Errorf(fmt.Sprintf("аргумент должен быть не более %v", maxCount))
		}
	}
	b := make([]byte, count)
	rand.Read(b)
	return b, nil
}
