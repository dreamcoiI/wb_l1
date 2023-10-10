package main

import "fmt"

type Human struct {
	Name string
	Pet  string
}

type Action struct {
	Name Human
}

func (h *Human) bestPet() string {
	return h.Pet
}

func main() {
	owner := Human{Name: "Masha", Pet: "dog"}
	animal := Action{Name: owner}
	animal.petHuman()
}

func (a *Action) petHuman() {
	fmt.Printf("%s the %s owner", a.Name.Name, a.Name.Pet)
}
