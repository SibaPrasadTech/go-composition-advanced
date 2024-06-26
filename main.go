package main

import "fmt"

type SpeakingAnimal interface {
	Speak()
}

func makeAnimalSpeak(sa SpeakingAnimal) {
	sa.Speak()
}

type Animal struct {
	Name string
	Weight int64
}

func (a Animal) Sleep() {
	fmt.Printf("Animal : %v is sleeping. \n",a.Name)
}

func (a Animal) Eat() {
	fmt.Printf("Animal : %v is eating. \n",a.Name)
}

func (a Animal) Speak() {
	fmt.Printf("Generic Sound by Animal : %v\n",a.Name)
}

type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Printf("Woof Woof by Dog : %v\n", d.Name)
}

type Cat struct {
	Animal
}

func (c Cat) Speak() {
	fmt.Printf("Meow Meow by Dog : %v\n", c.Name)
}

func main() {
	fmt.Println("Go composition basics.")
	a := Animal{
		Name: "Simba",
		Weight: 2000,
	}
	a.Eat()
	a.Sleep()
	d := Dog{
		Animal: Animal{
			Name: "Tommy",
			Weight: 100,
		},
	}
	d.Eat()
	d.Sleep()

	//a.Speak()
	makeAnimalSpeak(a)
	makeAnimalSpeak(d)
	// d.Speak() // This works as a Overloaded method
	// d.Animal.Speak() // But still the Embedded Struct's (Parents) Speak is available
}