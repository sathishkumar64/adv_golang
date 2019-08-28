package old

import "fmt"

//Person is a person
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d", p.Name, p.Age)
}
