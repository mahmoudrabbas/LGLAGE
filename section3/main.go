package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact

var contactIndexByName map[string]int

var nextId = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)

}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Println("Contact is Already exists..")
		return
	}

	newContact := Contact{
		ID:    nextId,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	contactList = append(contactList, newContact)
	nextId++

	fmt.Println("Contact Created..")
	contactIndexByName[name] = len(contactList) - 1

}

func findContact(name string) *Contact {
	if idx, exists := contactIndexByName[name]; exists {
		return &contactList[idx]
	} else {
		return nil
	}

}

func ListContacts() {
	if len(contactList) == 0 {
		fmt.Println("No Contacts..")
	} else {
		for _, val := range contactList {
			fmt.Println(val)
		}
	}
}

func main() {

	ListContacts()
	addContact("Mahmoud", "m@f.com", "255555")
	addContact("Zezo", "z@f.com", "255555")
	addContact("Mostafa", "mos@f.com", "255555")
	addContact("Diaa", "d@f.com", "255555")

	ListContacts()

	contact := findContact("Alaa")

	if contact != nil {
		fmt.Print(contact.Email)

	} else {
		fmt.Println("Doesnt exist.")
	}

	// // arr := [5]string{"Mahmoud", "Mohamed", "Ahmed", "Ali"}

	// // var arr [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	// var arr [3][2]int
	// // arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	// // arr[0] = 1
	// // arr[1] = 2

	// fmt.Printf("%+v\n", arr)

	// slices

	// slice := []int{}

	// fmt.Printf("%+v\n", slice)
	// slice = append(slice, 5)
	// fmt.Printf("%+v\n", slice)

	// slice = append(slice, 6)
	// fmt.Printf("%+v\n", slice)

	// slice := make([]int, 0, 3)
	// slice = append(slice, 1)
	// fmt.Printf("%+v size: %d cap:%d \n", slice, len(slice), cap(slice))
	// slice = append(slice, 2)
	// fmt.Printf("%+v size: %d cap:%d \n", slice, len(slice), cap(slice))
	// slice = append(slice, 3)
	// fmt.Printf("%+v size: %d cap:%d \n", slice, len(slice), cap(slice))
	// slice = append(slice, 4)
	// fmt.Printf("%+v size: %d cap:%d \n", slice, len(slice), cap(slice))

	// slice[3] = 5

	// fmt.Printf("%+v size: %d cap:%d \n", slice, len(slice), cap(slice))

	// myMap := map[int]string{
	// 	0: "Mahmoud",
	// 	1: "Ali",
	// }

	// fmt.Printf("%+v\n", myMap[0])

	// name, ok := myMap[0]

	// if ok {
	// 	fmt.Println(name)
	// }

	// slice := []int{0, 1, 2, 3, 4, 5}

	// fmt.Printf("original: %v len %v cap %v\n", slice, len(slice), cap(slice))

	// slice = append(slice, 5)
	// fmt.Printf("original: %v len %v cap %v\n", slice, len(slice), cap(slice))

}
