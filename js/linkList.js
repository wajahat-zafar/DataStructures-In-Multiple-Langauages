function Node(value) {
    return {
      value: value,
      next: null,
    };
  }
  
  function LinkedList() {
    const list = {
      head: null,
      tail: null,
      length: 0,
    };
  
    function add(value) {
      const node = Node(value);
      if (list.head === null) {
        list.head = node;
        list.tail = node;
      } else {
        list.tail.next = node;
        list.tail = node;
      }
      list.length++;
    }
  
    function insert(value, index) {
      if (index < 0 || index > list.length) {
        throw new Error('Index out of range');
      }
  
      const node = Node(value);
      if (index === 0) {
        node.next = list.head;
        list.head = node;
        if (list.tail === null) {
          list.tail = node;
        }
      } else if (index === list.length) {
        list.tail.next = node;
        list.tail = node;
      } else {
        let current = list.head;
        for (let i = 0; i < index - 1; i++) {
          current = current.next;
        }
        node.next = current.next;
        current.next = node;
      }
      list.length++;
    }
  
    function search(value) {
      let current = list.head;
      while (current !== null) {
        if (current.value === value) {
          return true;
        }
        current = current.next;
      }
      return false;
    }
  
    function remove(value) {
      if (list.head === null) {
        return;
      }
      if (list.head.value === value) {
        list.head = list.head.next;
        if (list.head === null) {
          list.tail = null;
        }
        list.length--;
        return;
      }
      let current = list.head;
      while (current.next !== null && current.next.value !== value) {
        current = current.next;
      }
      if (current.next !== null) {
        current.next = current.next.next;
        if (current.next === null) {
          list.tail = current;
        }
        list.length--;
      }
    }
  
    function sort() {
      const result = [];
      let current = list.head;
      while (current !== null) {
        result.push(current.value);
        current = current.next;
      }
      result.sort((a, b) => a - b);
      return result;
    }
  
    return {
      add,
      insert,
      search,
      remove,
      sort,
    };
  }
  const linkedList = LinkedList();

  linkedList.add(5);
  linkedList.add(3);
  linkedList.add(7);
  linkedList.add(1);
  linkedList.add(9);
  
  console.log(linkedList.sort()); 
  
  linkedList.remove(7);
  
  console.log(linkedList.search(7)); 
  
  linkedList.insert(6, 2);
  
  console.log(linkedList.sort());
  