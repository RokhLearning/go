package main

import "fmt"

type contactInfo struct {
	email    string
	zipconde int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var me person = person{"Behnam", "Aslani"}
	// me := person{firstName: "Behnam", lastName: "Aslani"}

	// var me person
	// fmt.Println(me.firstName)
	// %+v prints all field name and values from variable only works with printf
	// fmt.Printf("%+v", me)

	// var max person
	// max.firstName = "max"
	// max.lastName = "Anderson"

	// fmt.Printf("\n%+v", max)

	jim := person{
		firstName: "jim",
		lastName:  "party",
		contactInfo: contactInfo{
			zipconde: 9400,
			email:    "jimparty@gmail.com",
		},
	}
	jimPointer := &jim
	jimPointer.updateFirsName("japanalian")
	jim.Print()
}

func (p *person) updateFirsName(newName string) {
	(*p).firstName = newName
}

func (p person) Print() {
	fmt.Printf("%+v", p)
}
