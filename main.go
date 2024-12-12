package main

import (
	"fmt"
	"rc/structs"
)

func main() {
	// datatype.Datatypes();
	// fmt.Println("Hello World");
	// practice.Practice();
	// arrayandslices.Demo()
	structs.Structs()
	steve := structs.NewPerson("steve","kinney",50);
	fmt.Println(steve.GetFirstName());

	steve.SetFirstName("Brain");

	fmt.Println(steve.GetFirstName());
}