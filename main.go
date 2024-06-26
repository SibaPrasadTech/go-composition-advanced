package main

import "fmt"

type Animal struct {
	Name string
	Weight int64
}

func (a Animal) Sleep() {
	fmt.Printf("Animal : %v is sleeping. \n",a.Name)
}

func (a Animal) Eat() {
	fmt.Printf("Animal : %v is sleeping. \n",a.Name)
}

type Bird struct {
	Animal // Embedded Struct - can be used for code reuse
	Migratory bool
	Color string
}

func (b Bird) Fly() {
	fmt.Printf("Bird : '%v' with Color : '%v' is flying. \n",b.Name,b.Color)
}

func main() {
	fmt.Println("Go composition basics.")
	a := Animal{
		Name: "Tiger",
		Weight: 2000,
	}
	a.Eat()
	a.Sleep()
	b := Bird{
		Animal: Animal{
			Name: "Crow",
			Weight: 90,
		},
		Color: "Black",
		Migratory: false,
	}
	b.Fly()
	b.Eat()
	b.Sleep()
}