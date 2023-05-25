class Array:
    def __init__(self,arrayElements):
        self.arrayElements = arrayElements

    def  printAllarrayElements(self):
        for  i in self.arrayElements:

            print(i)
    def sort(self):
        for  i in range(len(self.arrayElements)):
            for j in range(i+1, len(self.arrayElements)):
                if self.arrayElements[i] > self.arrayElements[j]:
                    self.arrayElements[i] , self.arrayElements[j] = self.arrayElements[j] , self.arrayElements[i]
        print(self.arrayElements)

    def search(self):
        number = int(input('Enter the number you want to search? = '))
        if number in self.arrayElements:

            print("Your searched number" ,number ,'is existed in the present dataset.')
        else:
            print("The number you asked is not existed in the present dataset")

    def add(self):
        addNumber = int(input("Enter the number which you want to add?"))
        if addNumber != None:
            self.arrayElements.append(addNumber)

            print(self.arrayElements)

    def remove(self):
        removeNumber  = int(input("Enter the number you want to remove from the list?"))
        if removeNumber in self.arrayElements:
            self.arrayElements.remove(removeNumber)

            print(self.arrayElements)

    def insert(self):
        insertNumber = int(input("Enter the number you want to insert in the array?"))
        definePosition = int(input("Define the position where you want to insert the number?"))
        if insertNumber != None and definePosition != None:
            self.arrayElements.insert(definePosition,insertNumber)

            print(self.arrayElements)
            
    

a1 = Array([5,8,9,3,2,5,6,1])

a1.printAllarrayElements()
a1.sort()
a1.search()
a1.add()
a1.remove()
a1.insert()
