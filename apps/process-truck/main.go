package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return ErrorNotFoundTruck
}

var (
	ErrorNotImplemented = errors.New("Not Implemented")
	ErrorNotFoundTruck  = errors.New("Truck Not Found")
)

func processTruck(truck Truck) error {
	fmt.Printf("Processing Truck %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error Loading Cargo %w", ErrorNotFoundTruck)
	}
	return ErrorNotImplemented
}

func main() {

	trucks := []Truck{
		{id: "truck_1"},
		{id: "truck_2"},
		{id: "truck_3"},
		{id: "truck_4"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived\n", truck.id)
		if err := processTruck(truck); err != nil {
			// log.Fatalf("Errors happened with:%s", err)
			if errors.Is(err, ErrorNotFoundTruck) {
				log.Fatalf("Errors happened with:%s", ErrorNotFoundTruck)
			}
		}
	}

}
