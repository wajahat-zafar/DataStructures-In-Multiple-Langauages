class Node {
    constructor(value) {
      this.value = value;
      this.left = null;
      this.right = null;
    }
  }
  
  class Tree {
    constructor() {
      this.root = null;
    }
  
    add(value) {
      const node = new Node(value);
      if (this.root === null) {
        this.root = node;
      } else {
        this.insert(node, this.root);
      }
    }
  
    insert(node, current) {
      if (node.value < current.value) {
        if (current.left === null) {
          current.left = node;
        } else {
          this.insert(node, current.left);
        }
      } else {
        if (current.right === null) {
          current.right = node;
        } else {
          this.insert(node, current.right);
        }
      }
    }
  
    search(value) {
      return this._search(value, this.root);
    }
  
    _search(value, current) {
      if (current === null) {
        return false;
      } else if (current.value === value) {
        return true;
      } else if (value < current.value) {
        return this._search(value, current.left);
      } else {
        return this._search(value, current.right);
      }
    }
  
    remove(value) {
      this.root = this._remove(value, this.root);
    }
  
    _remove(value, current) {
      if (current === null) {
        return null;
      } else if (value === current.value) {
        if (current.left === null && current.right === null) {
          return null;
        } else if (current.left === null) {
          return current.right;
        } else if (current.right === null) {
          return current.left;
        } else {
          const tempNode = this._findMin(current.right);
          current.value = tempNode.value;
          current.right = this._remove(tempNode.value, current.right);
          return current;
        }
      } else if (value < current.value) {
        current.left = this._remove(value, current.left);
        return current;
      } else {
        current.right = this._remove(value, current.right);
        return current;
      }
    }
  
    _findMin(current) {
      while (current.left !== null) {
        current = current.left;
      }
      return current;
    }
  
    sort() {
      const result = [];
      this._traverseInOrder(this.root, (node) => result.push(node.value));
      return result;
    }
  
    _traverseInOrder(node, fn) {
      if (node !== null) {
        this._traverseInOrder(node.left, fn);
        fn(node);
        this._traverseInOrder(node.right, fn);
      }
    }
  }
  
  const tree = new Tree();

  tree.add(5);
  tree.add(3);
  tree.add(7);
  tree.add(1);
  tree.add(9);
  
  console.log(tree.sort()); 
  
  console.log(tree.search(3)); 
  console.log(tree.search(8)); 
  
  tree.remove(5);
  
  console.log(tree.sort()); 
    