package animal

type Animal struct {
	name         string
	age          int
	isGoodAtMath bool
}

type Cat struct {
	Animal
	Color string
}

type Dog struct {
	Animal
	Kind string
}

type AnimalInterface interface {
	TellMyName() string
	DoMath(x int, y int) int
	Sound() string
}

func NewAnimal(name string, age int, isGoodAtMath bool) *Animal {
	animal := Animal{
		name:         name,
		age:          age,
		isGoodAtMath: isGoodAtMath,
	}

	return &animal
}

func (a Animal) TellMyName() string {
	return a.name
}

func (a Animal) DoMath(x int, y int) int {
	return x + y
}

func (c Cat) Sound() string {
	return "Meow!"
}

func (d Dog) Sound() string {
	return "Bark!"
}
