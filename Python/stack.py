class Stack:
    def __init__(self,stackElements):
        self.stackElements = stackElements

    def push(self):
        askElement = int(input("How  many elemnents you want to insert in Stack?"))
        while len(self.stackElements) < askElement:
            addElement = int(input("Enter the elements you want to insert in Stack?"))
            self.stackElements.append(addElement)
            print(self.stackElements)
    def sort(self):
        for  i in range(len(self.stackElements)):
            for j in range(i+1, len(self.stackElements)):
                if self.stackElements[i] > self.stackElements[j]:
                    self.stackElements[i] , self.stackElements[j] = self.stackElements[j] , self.stackElements[i]
        print(self.stackElements)
    def search(self):
        number = int(input('Enter the number you want to search? = '))
        if number in self.stackElements:

            print("Your searched number" ,number ,'is existed in stack.')
        else:
            print("The number you asked is not existed in the stack")

    def Pop(self):
            

            print(self.stackElements.pop())
            print(self.stackElements)

    def Insert(self):
        insertNumber = int(input("Enter the number you want to insert in the stack?"))
        definePosition = int(input("Define the position where you want to insert the number?"))
        if insertNumber != None and definePosition != None:
            self.stackElements.insert(definePosition,insertNumber)

            print(self.stackElements)
    
    



s = Stack([])
s.push()
s.sort()
s.search()
s.Pop()
s.Insert()