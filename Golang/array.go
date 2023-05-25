package main

import (
	"fmt"
)

// func Array(myArr [6]int) {
// 	var ask int
// 	fmt.Println("Enter the number you want to search?.... ") 
// 	fmt.Scanln(&ask)
// 	var addNumber int
// 	fmt.Println("Enter the number you want to add in the array?....")
// 	fmt.Scanln(&addNumber)
// // Searching number
// 	found := false
// 	for i := 0; i < len(myArr); i++ {
// 		if (myArr)[i] == ask {
// 			found = true
// 			break
// 		}
// 	}

// 	if found {
// 		fmt.Println("Your searched number is present in the array")
// 	} else {
// 		fmt.Println("Your searched number is not present in the array")
// 	}
// // Adding number
// 	myArr[5] = addNumber
// 	fmt.Println("Your updated array is:", myArr)
// }

// func main() {
// 	arr := [6]int{1, 3, 5, 7, 9, 0}
// 	Array(arr)
// }


//remove element

// func removeElement(myArr []int, index int) []int {
	
// 	return append(myArr[:index], myArr[index+1:]...)
// }

// func main() {
// 	my_Array := []int{1, 3, 5, 7, 9}


// 	var askRemove int
// 	fmt.Println("Enter the index number of the element you want to remove?.... ") 
// 	fmt.Scanln(&askRemove)
// 	my_Array = removeElement(my_Array, askRemove)

// 	fmt.Println("Updated array:", my_Array)
// }

//Sorting Array
// func sort(myArray [] int) []int{
// 	for i:=0;  i<len(myArray)-1; i++{
// 		for j:=0 ; j<len(myArray)-1; j++{
// 			if myArray[j]> myArray[j+1]{
// 				myArray[j],myArray[j+1] = myArray[j+1],myArray[j]

// 			}
// 		}
// 	}
// 	return myArray
// }

// func main(){
// 	myArray := [5]int{15,43,2,6,1}
// 	result := sort(myArray[:])
// 	fmt.Println(result)
// }


//Inserting in a array

func Insert( originalArray [] int , index int , value int) [] int{
	if len(originalArray) == index{
		result:= append(originalArray, value)
		return result
	} else{

		originalArray = append(originalArray[:index+1],originalArray[index:]...)
		originalArray[index] = value
		return originalArray
	}
	
}

func main(){
	originalArray := []int {10,30,40}
	originalArray = Insert(originalArray,1,20)
	fmt.Println(originalArray)

}