package arrayandslices

import (
	"fmt"
	"slices"
	"sort"
)

func Demo() {
	// ? fixed length DS -> we have to provide the size of the array / containers at the declaration [1,2,3,3,4]
	// ? [1,2,3,4] int = 4bytes -> 1000 , 1004, 1008, 1012.

	// ? dynamic length DS -> Length is incremented / changed at the runtime linkedlist, stack, queue, deque, priorityqueye
	// ? tree, bt, bst, avl


	// var hobbies [4] string; // ? var <identifier> [sizeofarray] <datatype>
	// hobbies = [4] string { "sports", "cooing", "reading","boxing" };

	var hobbies [5] string = [5] string { "sports", "cooking", "reading","boxing", "ab" };

	fmt.Println(hobbies[3]); // ? this value is by default value
	// ? shorthand way
	prices := [5] float64 { 10,20,30,40 } // ? { 10,30 } not substring but subsequence
	fmt.Println(prices[4]); // ? by default value

	fmt.Println(hobbies)
	fmt.Println(prices)

	// ? slices / substring
	hobbiesSlice := hobbies[0:];

	fmt.Println(hobbiesSlice)
	// var hobbies [4] string = [4] string { "sports", "cooking", "reading","boxing" };

	hobbiesSlice = hobbiesSlice[1:]

	fmt.Println(len(hobbiesSlice))

	fmt.Println(hobbiesSlice[0]) // sports
	fmt.Println(hobbiesSlice[1]) // cooking
	// fmt.Println(hobbiesSlice[2]) // reading
	//fmt.Println(hobbiesSlice[3]) // ? error ARRAY INDEX OUT OF BOUND 
	
	fmt.Println(hobbiesSlice);

	hobbiesSlice[0] = "walking";
	fmt.Println(hobbiesSlice)
	fmt.Println(hobbies)

	hobbies[0] = "sports";
	fmt.Println(hobbiesSlice)
	fmt.Println(hobbies)


	// ? Utility functions
	fmt.Println(len(hobbiesSlice));
	fmt.Println(cap(hobbiesSlice));


	// ? Dynamic Arrays -> slice
	productPrices := [] int64 { 20,10 };
	productPrices = append(productPrices, 30);
	
	
	// fmt.Println(productPrices);


	// fmt.Println(productPrices[2]); // ? IOB
	
	productPricesUpdated := append(productPrices, 30);
	
	fmt.Println(productPricesUpdated)
	fmt.Println(productPrices)  // ? Original productProces won't changed.


	for index := 0; index < len(productPricesUpdated); index++ {
		fmt.Println(productPricesUpdated[index]);
	}

	slices.Sort(productPricesUpdated);
	fmt.Println(productPricesUpdated)

	sort.Slice(productPricesUpdated, func(i, j int) bool {
		return productPricesUpdated[i] > productPricesUpdated[j];
	})

	fmt.Println(productPricesUpdated)
}