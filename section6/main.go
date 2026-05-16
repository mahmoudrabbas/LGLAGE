package main

import (
	"errors"
	"fmt"
)

type Address struct {
	Street string
	City   string
}

func (a Address) GetFullAddress() string {
	return fmt.Sprintf("The City: %s, and the street:%s \n", a.City, a.Street)
}

type Person struct {
	Name string
	Address
}

// bank system project

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

// Deposit
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("The Amount Is Invalid.")
	}
	a.Balance += amount
	fmt.Printf("The Amount %0.2f is Deposited and your Balance became:%0.2f \n", amount, a.Balance)
	return nil
}

// Withdraw

func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return errors.New("The amount: %0.2f is insufficient to Withdraw")
	}
	a.Balance -= amount
	fmt.Printf("The Amount %0.2f is Withdrawed and your Balance became:%0.2f \n", amount, a.Balance)
	return nil

}

// get Balance
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// string()

func (a *Account) String() string {
	return fmt.Sprintf(` --- Account Data ---\n 
	Account Number:%s\n
	Account Owner:%s\n
	Account Balance%0.2f\n
	`, a.AccountNumber, a.OwnerName, a.Balance)
}

type SavingAccount struct {
	Account
	InterestsRate float64
}

func (sa *SavingAccount) AddInterests() {
	interests := sa.InterestsRate * sa.Balance

	sa.Deposit(interests)
	fmt.Printf("Interests: %0.2f is Deposited and Balance became %0.2f \n", interests, sa.Balance)
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Invalid Amount.")
	}

	if oa.OverdraftLimit+oa.Balance < amount {
		return fmt.Errorf("Withdraw with %0.2f exceeds the limit for Account Number: %s", amount, oa.String())

	}

	oa.Balance -= amount

	fmt.Println("Done Withdraw for Overdraft account.")

	return nil
}

func main() {

	fmt.Println("=== Bank Account System ===")

	account := Account{AccountNumber: "123", OwnerName: "Mahmoud", Balance: 1000}

	fmt.Printf("%+v\n", account)

	err := account.Deposit(200)
	if err != nil {
		fmt.Println(err.Error())
	}

	sa := SavingAccount{
		Account:       account,
		InterestsRate: 0.25,
	}

	sa.AddInterests()

	fmt.Printf("%+v\n", account)

	od := OverdraftAccount{
		Account:        account,
		OverdraftLimit: 100,
	}

	err = od.Deposit(10)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = od.Withdraw(500)
	if err != nil {
		fmt.Println(err.Error())
	}

	// od := OverdraftAccount{
	// 	Account:        account,
	// 	OverdraftLimit: 200,
	// }

	// addr := Address{
	// 	City:   "Minia",
	// 	Street: "44 St",
	// }

	// Person := Person{
	// 	Name:    "Mahmoud",
	// 	Address: addr,
	// }

	// fmt.Println(Person)

}
