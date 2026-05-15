package main

import (
	"fmt"
)

type Person interface {
	GetName() string
	GetId() int
}

type Employee struct {
	Id   int
	Name string
}
type BusinessEmployee struct {
	Id   int
	Name string
}

func (e Employee) GetName() string {
	return e.Name
}

func (e BusinessEmployee) GetName() string {
	return e.Name
}

func (e Employee) GetId() int {
	return e.Id
}

func (e BusinessEmployee) GetId() int {
	return e.Id
}

func DisplayPerson(p Person) {
	fmt.Println(p.GetName())
}

func NewEmployee(id int, fn string) Employee {
	return Employee{
		Id:   id,
		Name: fn,
	}
}

func (e Employee) String() string {
	return fmt.Sprintf("%s comes from string\n", e.Name)
}

// func (e *Employee) FullName() string {
// 	return e.FirstName + " " + e.LasstName
// }

// func (e *Employee) Deactivate() {
// 	e.IsActive = false
// }

type Number interface {
	int | float64 | string
}

func sum[T Number](args ...T) T {
	var sum T

	for _, val := range args {
		sum += val
	}

	return sum
}

// payroll program
type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("%s is Salaried Employee, Salary: %0.2f", se.Name, se.CalculatePay())
}

type HoursEmployee struct {
	Name       string
	HourRate   float64
	HourSalary float64
}

func (he HoursEmployee) CalculatePay() float64 {
	return he.HourRate * he.HourSalary
}

func (he HoursEmployee) String() string {
	return fmt.Sprintf("%s is Hours Employee, Salary: %0.2f", he.Name, he.CalculatePay())
}

type CommisionEmployee struct {
	Name          string
	BaseSalary    float64
	CommisionRate float64
	SalesAmount   float64
}

func (ce CommisionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommisionRate * ce.SalesAmount)
}

func (ce CommisionEmployee) String() string {
	return fmt.Sprintf("%s is Commision Employee, Salary: %0.2f", ce.Name, ce.CalculatePay())
}

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

func PrintEmployeeSummary[T fmt.Stringer](e T) {
	fmt.Printf("- Processing: %s\n", e)
}

func PaymentPayrollProcess(emps []Payable) {
	fmt.Println("Payment Payroll Processing Has Started:")

	var total float64 = 0.0

	for _, emp := range emps {
		PrintEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("Monthly Paid: %0.2f\n", pay)
		total += pay
	}

	fmt.Printf("Total Monthly Payroll: %0.2f\n", total)
}

func main() {

	se := SalariedEmployee{
		Name:         "Mahmoud",
		AnnualSalary: 150000,
	}

	he := HoursEmployee{
		Name:       "Khaled",
		HourRate:   40,
		HourSalary: 20,
	}

	ce := CommisionEmployee{
		Name:          "Hossam",
		BaseSalary:    20000,
		CommisionRate: 15,
		SalesAmount:   2,
	}

	payrollList := []Payable{
		se,
		he,
		ce,
		SalariedEmployee{Name: "Khalef", AnnualSalary: 250000},
	}

	PaymentPayrollProcess(payrollList)

	// fmt.Printf("%T \n", sum("1", "2", "3"))

	// emp1 := Employee{
	// 	Id:   1,
	// 	Name: "Mahmoud",
	// }

	// bEmp := BusinessEmployee{
	// 	Id:   2,
	// 	Name: "Abbas",
	// }

	// DisplayPerson(emp1)
	// DisplayPerson(bEmp)

	// fmt.Print(emp1)

}
