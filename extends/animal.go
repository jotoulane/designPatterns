package extends

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Eat() string {
	return fmt.Sprintf("%s is eating", a.Name)
}

type Cat struct {
	*Animal
}
