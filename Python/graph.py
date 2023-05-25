class Graph:
    def __init__(self):
        self.adj_list = {}

    def add_node(self, node):
        if node not in self.adj_list:
            self.adj_list[node] = []

    def add_edge(self, node1, node2):
        if node1 not in self.adj_list:
            self.add_node(node1)
        if node2 not in self.adj_list:
            self.add_node(node2)
        self.adj_list[node1].append(node2)
        self.adj_list[node2].append(node1)

    def remove_node(self, node):
        if node in self.adj_list:
            del self.adj_list[node]
            for adj_nodes in self.adj_list.values():
                if node in adj_nodes:
                    adj_nodes.remove(node)

    def remove_edge(self, node1, node2):
        if node1 in self.adj_list and node2 in self.adj_list:
            if node2 in self.adj_list[node1]:
                self.adj_list[node1].remove(node2)
                self.adj_list[node2].remove(node1)

    def search(self, node):
        if node in self.adj_list:
            return True
        return False

    def sort(self):
        for adj_nodes in self.adj_list.values():
            adj_nodes.sort()

    def insert_node(self, node, adj_nodes):
        if node in self.adj_list:
            return
        self.adj_list[node] = []
        for adj_node in adj_nodes:
            if adj_node in self.adj_list:
                self.adj_list[node].append(adj_node)
                self.adj_list[adj_node].append(node)

    def __str__(self):
        nodes = []
        for node, adj_nodes in self.adj_list.items():
            nodes.append(f"{node}: {', '.join(str(x) for x in adj_nodes)}")
        return "\n".join(nodes)
graph = Graph()
graph.add_edge(1, 2)
graph.add_edge(2, 3)
graph.add_edge(3, 1)
print(graph)

graph.sort()
print(graph)

graph.remove_edge(2, 3)
print(graph)


