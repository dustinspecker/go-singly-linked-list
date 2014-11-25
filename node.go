// Package singlyLinkedList implements a singly linked list
package go_singly_linked_list

// Node is a node within a singly linked list
type Node struct {
  Value interface{}

  next *Node
}

// Next returns Node n's next node
func (n *Node) Next() *Node {
  return n.next
}
