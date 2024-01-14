package main

import (
	"fmt"
)

func main() {
	app := NewApplication()
	err := app.Start(":8080")
	fmt.Print(err)
}
