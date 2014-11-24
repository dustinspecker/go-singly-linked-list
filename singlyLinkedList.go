// Package singlyLinkedList implements a singly linked list
package singlyLinkedList

// SinglyLinkedList is a singly linked list implementation
type SinglyLinkedList struct {
  front *Node

  length int
}

// Init initializes an empty list
func (s *SinglyLinkedList) Init() *SinglyLinkedList {
  s.length = 0
  return s
}

// New returns an initialized list
func New() *SinglyLinkedList {
  return new(SinglyLinkedList).Init()
}

// Front returns the first node in list s
func (s *SinglyLinkedList) Front() *Node {
  return s.front
}

// Back returns the last node in list s
func (s *SinglyLinkedList) Back() *Node {
  currentNode := s.front
  for currentNode != nil && currentNode.next != nil {
    currentNode = currentNode.next
  }
  return currentNode
}

// Append appends node n to list s
func (s *SinglyLinkedList) Append(n *Node) {
  if s.front == nil {
    s.front = n
  } else {
    currentNode := s.front

    for currentNode.next != nil {
      currentNode = currentNode.next
    }

    currentNode.next = n
  }

  s.length++
}

// Prepend prepends node n to list s
func (s *SinglyLinkedList) Prepend(n *Node) {
  if s.front == nil {
    s.front = n
  } else {
    n.next = s.front
    s.front = n
  }

  s.length++
}

// InsertBefore inserts node insert before node before in list s
func (s *SinglyLinkedList) InsertBefore(insert *Node, before *Node) {
  if s.front == before {
    insert.next = before
    s.front = insert
    s.length++
  } else {
    currentNode := s.front
    for currentNode != nil && currentNode.next != nil && currentNode.next != before {
      currentNode = currentNode.next
    }

    if currentNode.next == before {
      insert.next = before
      currentNode.next = insert
      s.length++
    }
  }
}

// InsertAfter inserts node insert after node after in list s
func (s *SinglyLinkedList) InsertAfter(insert *Node, after *Node) {
  currentNode := s.front
  for currentNode != nil && currentNode != after && currentNode.next != nil {
    currentNode = currentNode.next
  }

  if currentNode == after {
    insert.next = after.next
    after.next = insert
    s.length++
  }
}

// Remove removes node n from list s
func (s *SinglyLinkedList) Remove(n *Node) {
  if s.front == n {
    s.front = n.next
    s.length--
  } else {
    currentNode := s.front

    // search for node n
    for currentNode != nil && currentNode.next != nil && currentNode.next != n {
      currentNode = currentNode.next
    }

    // see if current's next node is n
    // if it's not n, then node n wasn't found in list s
    if currentNode.next == n {
      currentNode.next = currentNode.next.next
      s.length--
    }
  }
}

// RemoveBefore removes node before node before
func (s *SinglyLinkedList) RemoveBefore(before *Node) {
  if s.front != nil && s.front != before {
    if s.front.next == before {
      s.front = before
    } else {
      currentNode := s.front
      for currentNode.next.next != nil && currentNode.next.next != before {
        currentNode = currentNode.next
      }
      if currentNode.next.next == before {
        currentNode.next = before
      }
    }
  }
}

// RemoveAfter removes node after node after
func (s *SinglyLinkedList) RemoveAfter(after *Node) {
  if s.front != nil && s.front.next != nil {
    currentNode := s.front
    for currentNode != after && currentNode.next != nil {
      currentNode = currentNode.next
    }

    if currentNode == after {
      currentNode.next = currentNode.next.next
    }
  }
}

// GetAtPos returns the node at index in list s
func (s *SinglyLinkedList) GetAtPos(index int) *Node {
  currentNode := s.front
  count := 0
  for count < index && currentNode != nil && currentNode.next != nil {
    currentNode = currentNode.next
    count++
  }

  if count == index {
    return currentNode
  } else {
    return nil
  }
}

// Find returns the node with matching value or nil if not found
func (s *SinglyLinkedList) Find(value interface{}) *Node {
  currentNode := s.front
  for currentNode != nil && currentNode.Value != value && currentNode.next != nil {
    currentNode = currentNode.next
  }

  return currentNode
}

// Length returns the amount of nodes in list s
func (s *SinglyLinkedList) Length() int {
  return s.length
}
