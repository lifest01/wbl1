//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от
//родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	email     string
	Age       int
	Action    Action
}

func NewHuman(firstName, lastName, email string, age int, action Action) *Human {
	return &Human{
		FirstName: firstName,
		LastName:  lastName,
		email:     email,
		Age:       age,
		Action:    action,
	}
}

type Action struct {
	WalkingSpeed
	RunningSpeed
}

type WalkingSpeed struct {
	speed int
}

type RunningSpeed struct {
	speed int
}

func (h *Human) GetIngo() {
	fmt.Printf("Firts name: %s\n Last name: %s\n Age: %d\n Email: %s\n", h.FirstName, h.LastName, h.Age, h.email)
}

func (a *Action) GetWalkingSpeed() int {
	return a.WalkingSpeed.speed
}

func (a *Action) GetRunningSpeed() int {
	return a.RunningSpeed.speed
}

func (h *Human) GetAllSpeed() string {
	return fmt.Sprintf("Скорость у %s:\n ходьбы %d\n бега %d", h.FirstName, h.Action.GetWalkingSpeed(), h.Action.GetRunningSpeed())
}

func main() {
	Boris := NewHuman("Boris", "ivanov", "email@email.com", 30, Action{
		WalkingSpeed: WalkingSpeed{5},
		RunningSpeed: RunningSpeed{20},
	})
	fmt.Println(Boris.Action.GetRunningSpeed())
	fmt.Println(Boris.GetAllSpeed())
}
