// ? 1.check if string is palindrome or not.
package practice

import "fmt"

func isPalindrome( str string ) bool {
    // ? a b c d a b c  --> c b a d c b a
	// ? n i t i n -> n i t i n
 	//  if 	originalString == reverse(originalString) ? "Yes, Palindrome" : "Not a Palindrome" 

	for startPointer := 0; startPointer  < len(str) ; startPointer++ {
		endPointer :=  len(str) - 1 - startPointer; //? expresion len(str) - 1 - startPointer
		if  (startPointer <= (endPointer)) && (str[startPointer] != str[endPointer])  {
			return false;
		}
	}
	return true;
}


// ? 2. func checkSetBits in a decimal number
func countSetBits(decimalNumber int) int {
	countOfSetBits := 0;
	// for decimalNumber > 0 {
	// 	if ((decimalNumber & 1) >= 1) {
	// 		countOfSetBits++
	// 	}
	// 	decimalNumber = decimalNumber >> 1; 
	// }
	for pos := 0; pos <= 31; pos++ {
		if ((decimalNumber & (1 << pos)) >= 1) {
			countOfSetBits++
		} 
	}
	// ? 10 --> 101  count := 0;  101 --> 10 (1 or 0) and (at what position) (010)
	// ? 5 --> 10  count = 1;
	// ? 2 --> 1
	// ? 1 -> 0  count = 2;
	return countOfSetBits;
}

func Practice() {
	// str := "abcdabc";
	// str = "nitin";
	// str = "a";
	// fmt.Println(isPalindrome(str));
	// str = "aa"
   	// fmt.Println(isPalindrome(str));

    fmt.Println(countSetBits(10)); // 2
	fmt.Println(countSetBits(7)); // 3
	fmt.Println(countSetBits(15)); // 4
	fmt.Println(countSetBits(128));
}