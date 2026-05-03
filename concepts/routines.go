package concepts

import (
	"fmt"
	"time"
)

func name(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(2 * time.Second)

	}
}

func main() {

	go name("Mahmoud")
	go name("Ra")

	time.Sleep(4 * time.Second)

}
