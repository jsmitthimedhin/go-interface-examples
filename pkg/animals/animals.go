package animals

import "fmt"

type (
	dog struct {
		name string
	}
	cat struct {
		name string
	}
	walkable interface {
		walk()
	}
	bahtable interface {
		requiresBath() bool
	}
)

// Initial example (without passing interface in the parameter)
// func requiresBath(d dog) bool {
// 	fmt.Printf("%s, needs a bath!", d.name)
// 	return true
// }

func requiresBath(i walkable) bool {
	return true
}

func (d dog) walk() {
	fmt.Println("The dog is walking")
}

func (c cat) walk() {
	fmt.Println("The cat is walking")
}

func testFunction() {
	// Since requiresBath takes in anything that implements the walkable interface, we can pass both the dog and cat struct into the function without errors
	c := cat{}
	d := dog{}
	requiresBath(c)
	requiresBath(d)
}
