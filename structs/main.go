package structs

import "fmt"

type Person struct {
	firstName string
	LastName  string
	Email string
	Age uint
}
func NewPerson(firstName string, lastName string, age uint) Person {
	return Person {
		firstName: firstName,
		LastName: lastName,
		Email: "abc@mail.com",
		Age: 40,
	}
}
// Reciever function
func (person Person) GetFirstName() string {
	return person.firstName
}


func (person *Person) SetFirstName (updatedName string) {
	person.firstName = updatedName
}
func Structs() {

	rahul := Person { firstName: "Rahul", LastName: "Chhabra", Email: "chhabra@mail.com", Age: 23 };
	fmt.Println(rahul.GetFirstName());
	john := Person { firstName: "John", LastName: "Doe", Email: "john@mail.com", Age: 55 };
	fmt.Println(john.GetFirstName());
	person1 := NewPerson("steve","kinney",56);

	fmt.Println(person1.firstName)

	fmt.Println(rahul.firstName)
	fmt.Println(john.LastName);
}



// func (person *Person) changeName(newName string) {
	
// }