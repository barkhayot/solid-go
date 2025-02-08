package main

import "fmt"

/*
Liskov Substitution Principle (LSP):
- Objects of a superclass should be replaceable with objects of a subclass **without altering** the correctness of the program.
- Subtypes must not break the behavior expected by their parent type.
*/

// Define a generic Bird interface.
type Bird interface {
	MakeSound() string
}

// Separate FlyingBird interface for birds that can fly.
type FlyingBird interface {
	Bird
	Fly() string
}

// Pigeon is a bird that can fly.
type Pigeon struct{}

func (p Pigeon) MakeSound() string {
	return "Coo"
}

func (p Pigeon) Fly() string {
	return "Pigeon is flying."
}

// Penguin is a bird but cannot fly.
type Penguin struct{}

func (p Penguin) MakeSound() string {
	return "Squawk"
}

// Function that works correctly for all birds.
func DescribeBird(b Bird) {
	fmt.Println("Sound:", b.MakeSound())
}

// Function that only works for flying birds.
func DescribeFlyingBird(fb FlyingBird) {
	fmt.Println("Sound:", fb.MakeSound())
	fmt.Println("Flying:", fb.Fly())
}

// func main() {
// 	pigeon := Pigeon{}
// 	penguin := Penguin{}

// 	DescribeBird(pigeon)
// 	DescribeBird(penguin)

// 	DescribeFlyingBird(pigeon) // ✅ Works fine.
// 	// DescribeFlyingBird(penguin) // ❌ This will not compile, preventing misuse.
// }
