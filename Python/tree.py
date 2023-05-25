class Node:
    def __init__(self, data):
        self.left = None
        self.right = None
        self.data = data

class BinaryTree:
    def __init__(self):
        self.root = None

    def add_node(self, data):
        node = Node(data)
        if self.root is None:
            self.root = node
        else:
            self._add_node(self.root, node)

    def _add_node(self, current_node, node):
        if node.data < current_node.data:
            if current_node.left is None:
                current_node.left = node
            else:
                self._add_node(current_node.left, node)
        else:
            if current_node.right is None:
                current_node.right = node
            else:
                self._add_node(current_node.right, node)

    def remove_node(self, data):
        self.root = self._remove_node(self.root, data)

    def _remove_node(self, current_node, data):
        if current_node is None:
            return None
        if data == current_node.data:
            if current_node.left is None and current_node.right is None:
                return None
            if current_node.left is None:
                return current_node.right
            if current_node.right is None:
                return current_node.left
            min_node = self._find_min(current_node.right)
            current_node.data = min_node.data
            current_node.right = self._remove_node(current_node.right, min_node.data)
        elif data < current_node.data:
            current_node.left = self._remove_node(current_node.left, data)
        else:
            current_node.right = self._remove_node(current_node.right, data)
        return current_node

    def _find_min(self, node):
        while node.left is not None:
            node = node.left
        return node

    def search(self, data):
        return self._search(self.root, data)

    def _search(self, current_node, data):
        if current_node is None:
            return False
        if current_node.data == data:
            return True
        if data < current_node.data:
            return self._search(current_node.left, data)
        else:
            return self._search(current_node.right, data)

    def sort(self):
        nodes = []
        self._inorder_traversal(self.root, nodes)
        self.root = None
        for node in nodes:
            self.add_node(node)

    def _inorder_traversal(self, current_node, nodes):
        if current_node is not None:
            self._inorder_traversal(current_node.left, nodes)
            nodes.append(current_node.data)
            self._inorder_traversal(current_node.right, nodes)

    def insert_node(self, data):
        if not self.search(data):
            self.add_node(data)

    def __str__(self):
        return self._print_tree(self.root)

    def _print_tree(self, current_node, level=0):
        if current_node is not None:
            output = self._print_tree(current_node.right, level + 1)
            output += "\n" + "  " * level + str(current_node.data)
            output += self._print_tree(current_node.left, level + 1)
            return output
        else:
            return ""

tree = BinaryTree()


tree.add_node(5)
tree.add_node(3)
tree.add_node(7)
tree.add_node(2)
tree.add_node(4)
tree.add_node(6)


print("Initial tree:")
print(tree)


print("Is 3 in the tree?", tree.search(3))
print("Is 10 in the tree?", tree.search(10))


tree.sort()
print("Sorted tree:")
print(tree)


tree.remove_node(4)
print("Tree after removing 4:")
print(tree)


tree.insert_node(2)
tree.insert_node(8)
print("Tree after inserting 2 and 8:")
print(tree)

