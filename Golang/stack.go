package main

import("fmt")

func pushStack(mypushStack [] int,value int) [] int{
	
	push:= append(mypushStack, value)
	return push
	
	
		
		
		
	}

func popStack (mypopStack [] int) (int, [] int){
		if len(mypopStack)  == 0 {
			fmt.Println("The stack is empty")
		}
	
		top:= mypopStack[len(mypopStack)-1]
	
		mypopStack = mypopStack[:len(mypopStack)-1]
		return top, mypopStack
	
	}

func sortStack (mysortStack [] int) [] int{
	for i:=0; i<len(mysortStack)-1; i++{
		for j:= 0 ; j<len(mysortStack)-1; j++{
			if mysortStack[j] > mysortStack[j+1]{
				mysortStack[j],mysortStack[j+1] = mysortStack[j+1] , mysortStack[j]
			}
		}
	}
	return mysortStack

}

func main(){
	myStack:= [] int {}
	myStack = pushStack(myStack , 2)
	myStack = pushStack(myStack , 3)
	myStack = pushStack(myStack , 4)
	myStack = pushStack(myStack , 5)
	myStack = pushStack(myStack , 3)
	myStack = pushStack(myStack , 8)
	myStack = pushStack(myStack , 14)
	myStack = pushStack(myStack , 9)
	
	fmt.Println(myStack)
	top, myStack2 := popStack(myStack)
	fmt.Println("Popped element:", top)
	fmt.Println("Remaining stack:", myStack2)
	mysortStack := sortStack(myStack)
	fmt.Println("The sorted stack would be:", mysortStack)

	var ask int
	fmt.Println("Enter the number you want to search?.... ") 
 	fmt.Scanln(&ask)
	found := false
	for i:=0 ; i<len(myStack); i++{
		if (myStack)[i] == ask{
			found = true
			break
		}


	}
	if found {
		fmt.Println("Your searched number is present in stack")
	} else {
		fmt.Println("Your searched number is not present in stack")
			}

	
}






