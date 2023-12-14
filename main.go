package main

import (
	"fmt"
	"go-math-island/animal"
	"go-math-island/island"
)

func main() {
	cat := animal.Cat{
		Animal: *animal.NewAnimal("Jiji", 3, true),
		Color:  "black",
	}

	dog := animal.Dog{
		Animal: *animal.NewAnimal("Clifford", 8, true),
		Kind:   "labradoodle",
	}

	islandService := island.Island{
		Cat: &cat,
		Dog: &dog,
	}

	intList := []int{}
	for j := 1; j <= 1000000; j++ {
		intList = append(intList, j)
	}

	fmt.Println(islandService.SumIntegersList(intList))
}
