package main

import (
	"fmt"

	"github.com/mahmoudrabbas/module-example/app"
)

func main() {

	user := app.User{
		ID:   1,
		Name: "Mahmoud",
	}

	fmt.Println("Started")

	fmt.Println(user)
}
