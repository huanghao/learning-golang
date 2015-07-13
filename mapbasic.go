package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var PersonDB map[string]PersonInfo
	PersonDB = make(map[string]PersonInfo)

	PersonDB["1"] = PersonInfo{"1", "Bob", "Room 2003"}
	PersonDB["2"] = PersonInfo{"2", "Alice", "Room 101"}

	person, ok := PersonDB["1"]
	if ok {
		fmt.Println("Found", person.Name, person.Address)
	} else {
		fmt.Println("Did not find")
	}

	fmt.Println(PersonDB["3"])

	var m = make(map[int]int)
	m[1] = 0
	fmt.Println("value is: ", m[1], "\nnot exist:", m[2])
}
