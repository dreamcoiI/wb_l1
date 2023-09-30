package main

import "fmt"

type Human struct {
	Name string
	Pet  string
}

type Owner struct {
	Name Human
}

func (h *Human) bestPet() string {
	return h.Pet
}

func main() {
	owner := Human{Name: "Masha", Pet: "dog"}
	animal := Owner{Name: owner}
	animal.petHuman()
}

func (a *Owner) petHuman() {
	fmt.Printf("%s the %s owner", a.Name.Name, a.Name.Pet)
}
