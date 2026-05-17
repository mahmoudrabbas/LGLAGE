package main

import (
	"fmt"

	"github.com/mahmoudrabbas/module-example/internal/model"
)

func main() {

	text := "Hello Go"

	fmt.Println(model.Text(text, model.Red, model.Underline))
}
