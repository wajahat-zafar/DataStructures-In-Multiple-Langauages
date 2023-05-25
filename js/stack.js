
//Adding or Pushing in Stack
let prompt = require("prompt-sync")({sigint:true});
let stackList = []
let inputNumbers =  prompt("How many elements you want to insert in the stack? = ")
while( stackList.length < inputNumbers ){
  ask = prompt("Enter the number you want to insert? = ");
  stackList.unshift(ask);
  
}

//Sorting stack

let numbers= stackList;
sortNumbers = numbers.sort(function(a, b){return a - b})
console.log("The sorted stack would be" + " " + sortNumbers)

//Searching
var askValue = prompt("Enter your required value to search? = ")
if (stackList.includes(askValue)){
  console.log("Your searched value is existed in stack")
} else{
  console.log("Your searched Value is not  existed in stack")
}
//Removing
var removeValue = prompt("Enter the value which you want to remove? = ");
for ( var i=0 ; i< stackList.length;  i++){
    if (stackList[i] === removeValue ){
      stackList.splice(i,1);
      i--;
      console.log(stackList)
    }
  }

//Pop from stack

stackList.pop()
console.log(stackList)

// insert in Stack
let position = prompt("Enter at which position you want to  insert the number? = ")
let insertNumber = prompt("Enter which number you want to insert? = ")
stackList.splice(position,0,insertNumber)
console.log(stackList)

