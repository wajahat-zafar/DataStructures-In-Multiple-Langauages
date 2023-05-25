
//Adding or Pushing in Queue
let prompt = require("prompt-sync")({sigint:true});
let queueList = []
let inputNumbers =  prompt("How many elements you want to insert in the queue? = ")
while( queueList.length < inputNumbers ){
  ask = prompt("Enter the number you want to insert? = ");
  queueList.unshift(ask);
  
}

//Sorting Queue

let numbers= queueList;
sortNumbers = numbers.sort(function(a, b){return a - b})
console.log("The sorted queue would be" + " " + sortNumbers)

//Searching in Queue
var askValue = prompt("Enter your required value to search? = ")
if (queueList.includes(askValue)){
  console.log("Your searched value is existed in queue")
} else{
  console.log("Your searched Value is not  existed in queue")
}
//Removing in Queue
var removeValue = prompt("Enter the value which you want to remove? = ");
for ( var i=0 ; i< queueList.length;  i++){
    if (queueList[i] === removeValue ){
      queueList.splice(i,1);
      i--;
      console.log(queueList)
    }
  }

//Pop from Queue

queueList.pop(0)
console.log(queueList)

// insert in Queue
let position = prompt("Enter at which position you want to  insert the number? = ")
let insertNumber = prompt("Enter which number you want to insert? = ")
queueList.splice(position,0,insertNumber)
console.log(queueList)

