package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//first define the struct "ruleset" - what fields does it have?
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // we can also just say contactInfo here, and it declares a field called contactInfo that has the type contactInfo!
}

// you can put custom types into structs, even other structs!

func main() {
	//then create a value of that struct
	// option 1:
	// alex := person{"Alex", "Anderson"} //go infers the order of definition of fields!

	// option 2:
	// alex := person{firstName: "Alex", lastName: "Anderson"} //less room for error this way!
	// fmt.Println(alex)

	// option 3:
	// var alex person
	// fmt.Println(alex)       // at this point, this would print {  } because of zero values => {""""}
	// fmt.Printf("%+v", alex) // %+v prints out all the field names and their values from alex!

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// note that in multiline declarations all lines must have a comma, even if it's the last line!
	// fmt.Printf("%+v", jim)
	jim.print()

	// jimPointer := &jim // store the memory address of the variable jim in jimPointer
	// jimPointer.updateName("Bob")

	// shorter way to do this:
	jim.updateName("Bob") // this works too even when the receiver function type is *person
	// go pointer shortcut: we can just pass a *person receiver function a person!!!
	// go automatically turns the person into its pointer! (yay!)

	jim.print()
}

// in this case we would call it as jim.updatePerson("Bob")
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName // when we do this, we are updating a copy of the struct, not the original struct!!!
// }

// for this function to work properly, we need to pass it the POINTER!!!
// so the receiver type is a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //fetch the value stored at the memory address this points to
}

// receiver function with structs
func (p person) print() {
	fmt.Printf("%+v", p)
}
