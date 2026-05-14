package main

import (
	// "fmt"
	// problemsolving "lgge/problem_solving"
	"lgge/concepts"
)

func main() {

	mysql := &concepts.MySqlRepo{}
	postgres := &concepts.PostgresRepo{}

	service1 := concepts.NewUserService(mysql)
	service2 := concepts.NewUserService(postgres)

	service1.PrintUser(5)
	service2.PrintUser(5)

}
