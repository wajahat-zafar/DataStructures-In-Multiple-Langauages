class Queue:
    def __init__(self,queueElements):
        self.queueElements = queueElements

    def push(self):
        askElement = int(input("How  many elemnents you want to insert in queue?"))
        while len(self. queueElements) < askElement:
            addElement = int(input("Enter the elements you want to insert in queue?"))
            self.queueElements.append(addElement)
            print(self.queueElements)
    def sort(self):
        for  i in range(len(self.queueElements)):
            for j in range(i+1, len(self.queueElements)):
                if self.queueElements[i] > self.queueElements[j]:
                    self.queueElements[i] , self.queueElements[j] = self.queueElements[j] , self.queueElements[i]
        print(self.queueElements)
    def search(self):
        number = int(input('Enter the number you want to search? = '))
        if number in self.queueElements:

            print("Your searched number" ,number ,'is existed in queue.')
        else:
            print("The number you asked is not existed in the queue")

    def Pop(self):
            

            print(self.queueElements.pop(0))
            print(self.queueElements)

    def Insert(self):
        insertNumber = int(input("Enter the number you want to insert in the queue?"))
        definePosition = int(input("Define the position where you want to insert the number?"))
        if insertNumber != None and definePosition != None:
            self.queueElements.insert(definePosition,insertNumber)

            print(self.queueElements)
    
    



s = Queue([])
s.push()
s.sort()
s.search()
s.Pop()
s.Insert()