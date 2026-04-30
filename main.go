package main

import (
	"fmt"
	"strconv"

	// "strconv"
	"strings"
)

func main() {
	// n := 5555

	// str := strconv.Itoa(n)

	fmt.Println(thousandSeparator(51040))
}

func thousandSeparator(n int) string {

	if n == 0 {
		return "0"
	}

	str := strconv.Itoa(n)

	l := len(str)
	r := l % 3

	c := l / 3

	var sb strings.Builder

	start := str[0:r]
	if start != "" {
		sb.WriteString(start + ".")
	}

	// c :=1
	for i := r; i < l-3+1; i += 3 {
		if c != 1 {
			fmt.Print(str[i:i+3] + " i ")
			fmt.Println(i)
			sb.WriteString(str[i:i+3] + ".")
		} else {
			fmt.Print(str[i:i+3] + " i ")
			fmt.Println(i)
			sb.WriteString(str[i : i+3])
		}

		c--
	}

	return sb.String()

}

// func thousandSeparator(n int) string {

// 	var sb strings.Builder

// 	var slice []string

// 	for n != 0 {
// 		r := n % 1000
// 		slice = append(slice, strconv.Itoa(r))
// 		n /= 1000
// 	}

// 	for i := len(slice) - 1; i >= 0; i-- {
// 		fmt.Println(slice[i])
// 		if i != 0 {
// 			sb.WriteString(slice[i] + ".")
// 		} else {
// 			sb.WriteString(slice[i])
// 		}

// 	}
// 	return sb.String()
// }
