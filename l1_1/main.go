package main

import "fmt"

type Human struct {
	Age    uint8
	Name   string
	Gender string
}

func (h Human) showName() {
	fmt.Printf("Hello, my name is %s\n", h.Name)
}

func (h Human) showStats() {
	fmt.Printf("Hello, i am %s %s %d years old\n", h.Gender, h.Name, h.Age)
}

type Action struct {
	Human
	Action string
}

func (a Action) showStats() {
	fmt.Printf("Hello, i am %s %s %d years old with action '%s'\n", a.Gender, a.Name, a.Age, a.Action)
}

func (a Action) showAction() {
	fmt.Printf("Hello, i am %s and i can %s\n", a.Name, a.Action)
}

func main() {
	Alice := Human{
		Age:    18,
		Name:   "Alice",
		Gender: "Female",
	}

	Alice.showName()
	Alice.showStats()

	MoveForward := Action{
		Human:  Alice,
		Action: "Move Forward",
	}

	MoveForward.showName()
	MoveForward.showStats()
	MoveForward.showAction()
}
