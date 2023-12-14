package main

import (
	"fmt"
	"go-math-island/animal"
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

	fmt.Println(WhatIsMyName(cat))
	fmt.Println(WhatIsMyName(dog))
}

func WhatIsMyName(myanimal animal.AnimalInterface) string {
	return myanimal.TellMyName()
}
