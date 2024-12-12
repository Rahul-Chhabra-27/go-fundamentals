package datatype

import (
	"fmt"
)

var helloFirstWay int = 100;
var helloSecondWay = 100;

var (
	t1 int = 35;
	t2 int = 40
	t3, t4,t5 int = 1,2,3
)

// helloShorthandWay := 100; ? we can't use it outside the function.

func Datatypes() {
	// Datatypes Available
	// 1. Number --> [  int, uint, int32]
	// 2. Boolean 
	// 3. string
	
	// syntax of declaring a variable

	// var identifier int = 10;
	// 1st way
	var a int;
	a = 20
	// ? most preferable one
	var b int = 20;

	var c int = a + b;

	var d,e int = 10,20
	var f int = d + e;


// ? 1st way Type ---> value
// ? 2nd way value ---> Type.

	// 2nd way
	// var second = 10.45;
	// second = 34.56

	var first,second = 10,20;
	var third = first + second;
	
	// var example int;
	// example = 20;

	// fmt.Println(example)
	// shorthand way

	firstShorthand := 10;
	secondShorthand := 90;
	thirdShorthand := firstShorthand + secondShorthand;

	// ? <identifier> := <value>
	age := 100;

	fmt.Println(c,f,third, age);

	fmt.Println(thirdShorthand);

	x,y,z := 1,2,3;
	
	fmt.Println(x,y,z);

	var str string = "some string";
	fmt.Println(str);

	strShorthandWay := "someString"
	fmt.Println(strShorthandWay)


	var bool1 bool = true;
	var bool2 = true;
	boolShorthandWay := true;
	varA := 10;
	fmt.Println(bool1,bool2, boolShorthandWay, varA);

	fmt.Println(sum(10,20))
	// var character string = "a";

	// ? varA = 10;
	// ? varA := 10; 
	// ? 1. var varA = 10;


	err,result :=  divideWithException(0,0);
	
	if err == "" {
		fmt.Print("Result ")
		fmt.Println(result)
	} else {
		fmt.Print("Error ")
		fmt.Println(err);
	}


	returnedFunction := multiplier(8);
	fmt.Println(returnedFunction(3));

	// IIFE
	func(){
		fmt.Println("Inside IIFE")
	}();

	// IIFE
	values := func(a int, b int) int {
		return a + b;
	}(10,80);
	fmt.Println(values);
}

// func <identifier> () <Reserved for the return type> {
// 	<body>
// }

func sum(first int, second int) int {
	return first + second
}

func divideWithException (first int , second int) (string, int) {

	if (first == 0 && second == 0) {
		return "Cannot Perform this operation", -1;
	}
	return "", first / second;
}


func multiplier(someValue int) func(int) int {
	return func (shiftHandler int) int {
		return someValue << shiftHandler;
	}
} 
