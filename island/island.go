package island

import (
	"go-math-island/animal"
)

type Island struct {
	Cat *animal.Cat
	Dog *animal.Dog
}

func (i Island) SumIntegersList(intList []int) int {
	sum := 0
	for j := 0; j < len(intList); j++ {
		sum = i.Dog.DoMath(sum, intList[j])
	}
	return sum
}
