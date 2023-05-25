
//Sorting

let numbers= [9,2,8,4,5,6];
sortNumbers = numbers.sort()
console.log("The sorted array would be" + " " + sortNumbers)

//Searching
const prompt = require("prompt-sync")({sigint:true});
var askValue = prompt("Enter your required value to search? = ")
let arrayElement = [7,4,5,8,11,2,1,66];
if (arrayElement === (askValue)){
  console.log("Your searched value is existed in an array")
} else{
  console.log("Your searched Value is not  existed in an array")
}


//Adding
// const prompt = require("prompt-sync")({sigint:true});
var addValue = prompt("Enter the value that you want to add? = ");
let arrayElement2 = [4,6,7,12,33,2,12,3];
if (addValue !=  null){
  arrayElement2.push(addValue)
  console.log("Your element is"+ " "+ addValue + " " +  "added in an array")
  console.log(arrayElement2)
}else{
  console.log("we cant add your element")
}

//Removing

let arrayElement3 = [1,3,5,7,11,23,43];
for ( var i=0 ; i< arrayElement3.length;  i++){
  if (arrayElement3[i] === 7 ){
    arrayElement3.splice(i,1);
    i--;
    console.log(arrayElement3)
    
  }
}

//Inserting

let arrayElement4 = [4,6,7,8,12];
arrayElement4.splice(2,0,6)
console.log(arrayElement4)

