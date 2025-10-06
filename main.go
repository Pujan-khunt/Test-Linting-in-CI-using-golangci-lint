package main

import (
	"errors"
)

func main() {
	returnDataAndError()
}

func returnDataAndError() error {
	return errors.New("simple error")
}
