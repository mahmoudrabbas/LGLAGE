package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TShirt": 55.50,
	"Book":   22.50,
	"Mobile": 88.50,
}

func CalculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]

	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItem := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found := productPrices[originalItem]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf("item: %s price: %0.2f orignal price: %0.2f\n", originalItem, salePrice, basePrice)
				return salePrice, true
			}
		}

		fmt.Printf("item: %s not found", itemCode)
		return 0.0, false

	}

	fmt.Printf("item: %s price: %0.2f\n", itemCode, basePrice)
	return basePrice, true
}

func main() {

	fmt.Println("=============== Sales Order Processor ==============")

	orderItems := []string{"TShirt", "Mobile", "Book_SALE"}

	totalPrice := 0.0

	for _, item := range orderItems {
		price, found := CalculateItemPrice(item)

		if found {
			totalPrice += price
		}
	}

	fmt.Println("totalPrice: ", totalPrice)

	// var x interface{} = 5

	// value, ok := x.(string)

	// if ok {
	// 	fmt.Println("ok:", value)
	// } else {
	// 	fmt.Println(value)
	// }

	// checkType := func(i interface{}) {
	// 	switch v := i.(type) {
	// 	case int:
	// 		fmt.Println("Integer:", v)
	// 	case string:
	// 		fmt.Println("String:", v)
	// 	case float64:
	// 		fmt.Println("float64:", v)
	// 	}
	// }

	// checkType("55.55")

	// switch hour := time.Now().Hour(); {
	// case hour < 12:
	// 	fmt.Println(time.Now().Second())
	// 	fmt.Println("Morning")
	// case hour < 17:
	// 	fmt.Println("Afternoon")
	// default:
	// 	fmt.Println("Evening")
	// }

	// day := "Thursday"

	// switch day {
	// case "Saturday", "Sunday":
	// 	fmt.Println("Weekend")
	// case "Monday", "Tuesday":
	// 	fmt.Println("Starting planning")
	// default:
	// 	fmt.Println("Development")

	// }

	// userAccess := map[string]bool{"Mahmoud": true, "Ali": false}

	// name := "Ali"
	// if val, ok := userAccess[name]; ok && val {
	// 	fmt.Println("this user:", name, "exists and has Access")
	// } else if ok {
	// 	fmt.Println("this user:", name, "exists and doesnt have an Access")
	// } else {
	// 	fmt.Println("this user:", name, "doesnt exist and has no Access")

	// }

}
