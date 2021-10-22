package pointers

import "fmt"

func Learn() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x) //get memory address

	y := x
	fmt.Println(y)
	fmt.Println(&y)

	fmt.Println("-----------")
	// var p *int = &x
	p := &x // create pointer p
	fmt.Println(p)
	fmt.Println(*p) // read value x by pointer p (dereference)
	*p = 20 //assign new value
	fmt.Println(*p)

	s := student{name: "Bob"}
	fmt.Println(s)
	setName(&s)
	fmt.Println(s)

}

type student struct {
	name string
}

func setName(std *student) {
	std.name = "Earth"
}