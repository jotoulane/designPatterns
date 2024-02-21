package extends

type IPerson interface {
	Eat()
}

type Person struct {
	Name string
}

func (a *Person) Eat() string {
	return a.Name + " is eating"
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

type Worker struct {
	*Person
}

func NewWorker(name string) *Worker {
	return &Worker{Person: NewPerson(name)}
}
