package mapandstruct

import "fmt"

func Learn() {
	// Maps (key, value)
	m := map[string]int{"token": 20, "access": 100}
	m["token"] = 100
	fmt.Println(m)

	//loop
	for k, v := range m {
		fmt.Printf("%v => %v \n", k, v)
	}

	//delete map
	delete(m, "access")
	for k, v := range m {
		fmt.Printf("%v => %v \n", k, v)
	}

	//add map
	m["sum"] = 200
	for k, v := range m {
		fmt.Printf("%v => %v \n", k, v)
	}

	// struct
	type User struct {
		id       int
		username string
		password string
	}

	john := User{
		id:       1,
		username: "John Doe",
		password: "1234",
	}

	fmt.Println(john.username)
	john.password = "12345678"
	fmt.Println(john.password)

	users := []User{
		{id: 1, username: "Mary", password: "1234"},
		{id: 2, username: "Bob", password: "1234"},
	}

	fmt.Println(users[0].username)

}