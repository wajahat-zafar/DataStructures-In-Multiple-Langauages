function Graph() {
    this.vertices = [];
  
    
    this.addVertex = function(vertex) {
      if (!this.vertices.includes(vertex)) {
        this.vertices.push(vertex);
        this[vertex] = [];
      }
    };
  
    
    this.removeVertex = function(vertex) {
      let index = this.vertices.indexOf(vertex);
      if (index !== -1) {
        this.vertices.splice(index, 1);
        while (this[vertex].length) {
          const adjacentVertex = this[vertex].pop();
          this.removeEdge(vertex, adjacentVertex);
        }
      }
    };
  
    
    this.addEdge = function(vertex1, vertex2) {
      if (this[vertex1] && this[vertex2]) {
        this[vertex1].push(vertex2);
        this[vertex2].push(vertex1);
      }
    };
  
    
    this.removeEdge = function(vertex1, vertex2) {
      if (this[vertex1] && this[vertex2]) {
        this[vertex1] = this[vertex1].filter(v => v !== vertex2);
        this[vertex2] = this[vertex2].filter(v => v !== vertex1);
      }
    };
  
    
    this.sort = function() {
      this.vertices.sort();
    };
  
    
    this.search = function(vertex) {
      if (this.vertices.includes(vertex)) {
        return `Vertex ${vertex} found in graph`;
      }
      return `Vertex ${vertex} not found in graph`;
    };
  

    this.print = function() {
      let graphString = "";
      this.vertices.forEach(vertex => {
        graphString += vertex + " -> ";
        graphString += this[vertex].join(", ") + "\n";
      });
      console.log(graphString);
    };
  }
  let myGraph = new Graph();
  myGraph.addVertex("A");
  myGraph.addVertex("B");
  myGraph.addVertex("C");
  myGraph.addEdge("A", "B");
  myGraph.addEdge("B", "C");
  myGraph.addEdge("C", "A");
  myGraph.print(); 
  
  myGraph.removeVertex("C");
  myGraph.print(); 
  
    