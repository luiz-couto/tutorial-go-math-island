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

func (i Island) SumIntegersListTogether(intList []int) int {
	if len(intList)%2 == 1 {
		intList = append(intList, 0)
	}

	numbersToEachAnimal := len(intList) / 2

	totalSum := 0

	go func() {
		for j := 0; j < numbersToEachAnimal; j++ {
			totalSum = i.Dog.DoMath(totalSum, intList[j])
		}
	}()

	go func() {
		for j := numbersToEachAnimal; j < len(intList); j++ {
			totalSum = i.Cat.DoMath(totalSum, intList[j])
		}
	}()

	return totalSum
}

func (i Island) SumIntegersListTogetherWithChannels(intList []int) int {
	if len(intList)%2 == 1 {
		intList = append(intList, 0)
	}

	numbersToEachAnimal := len(intList) / 2

	sumChannel := make(chan int)

	go func() {
		sum := 0
		for j := 0; j < numbersToEachAnimal; j++ {
			sum = i.Dog.DoMath(sum, intList[j])
		}
		sumChannel <- sum
	}()

	go func() {
		sum := 0
		for j := numbersToEachAnimal; j < len(intList); j++ {
			sum = i.Cat.DoMath(sum, intList[j])
		}
		sumChannel <- sum
	}()

	dogSum, catSum := <-sumChannel, <-sumChannel
	totalSum := dogSum + catSum

	return totalSum
}
