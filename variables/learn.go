package variables

import "fmt"

var fullname string //global
var email string = "example@example.com"
var c, python bool = true, false
const vat int = 7

func Learn() {
	fullname = "John Doe"
	fmt.Println(fullname)
	fmt.Println(email)
	fmt.Printf("Hello %s Email %s \n", fullname, email)
	fmt.Println(c, python)

	countryName := "Thailand"
	fmt.Println(countryName)

	fmt.Println(vat)
}