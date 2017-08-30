package conf

import (
	"fmt"
	"os"
)

// Envs
const (
	Dev  string = "development"
	Prod string = "production"
	Test string = "test"
)

// Env ..
var Env = Dev

func setENV(e string) {
	if len(e) > 0 {
		Env = e
	}
}

func init() {
	setENV(os.Getenv("ENV"))
	fmt.Println("ENV is: " + Env)
}
