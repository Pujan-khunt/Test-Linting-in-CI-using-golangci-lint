package main

import (
	"errors"
	"fmt"
)

func main() {
	err := returnDataAndError()
	if err != nil {
		fmt.Println("error occured.")
	}
}

func returnDataAndError() error {
	return errors.New("simple error")
}
