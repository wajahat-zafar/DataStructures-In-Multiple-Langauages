class Node:
    def __init__(self, data=None):
        self.data = data
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None

    def add(self, data):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
        else:
            current = self.head
            while current.next is not None:
                current = current.next
            current.next = new_node

    def remove(self, data):
        if self.head is None:
            return
        if self.head.data == data:
            self.head = self.head.next
            return
        current = self.head
        while current.next is not None:
            if current.next.data == data:
                current.next = current.next.next
                return
            current = current.next

    def insert(self, data, position):
        new_node = Node(data)
        if position == 0:
            new_node.next = self.head
            self.head = new_node
            return
        current = self.head
        for i in range(position - 1):
            if current is None:
                return
            current = current.next
        new_node.next = current.next
        current.next = new_node

    def search(self, data):
        current = self.head
        while current is not None:
            if current.data == data:
                return True
            current = current.next
        return False

    def sort(self):
        if self.head is None:
            return
        current = self.head
        while current.next is not None:
            next_node = current.next
            while next_node is not None:
                if current.data > next_node.data:
                    temp = current.data
                    current.data = next_node.data
                    next_node.data = temp
                next_node = next_node.next
            current = current.next

    def __str__(self):
        if self.head is None:
            return "Empty List"
        current = self.head
        nodes = []
        while current is not None:
            nodes.append(str(current.data))
            current = current.next
        return "->".join(nodes)
linked_list = LinkedList()
linked_list.add(3)
linked_list.add(1)
linked_list.add(2)
print(linked_list)  
linked_list.sort()
print(linked_list)  
linked_list.remove(2)
print(linked_list)  
linked_list.insert(2, 1)
print(linked_list)  
print(linked_list.search(2))  
print(linked_list.search(4))  