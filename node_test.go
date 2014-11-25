package go_singly_linked_list

import "testing"

func TestNext(t *testing.T) {
  s := New()
  n1 := &Node{Value: "test1"}
  n2 := &Node{Value: "test2"}

  s.Append(n1)
  s.Append(n2)

  if n1.Next() != n2 {
    t.Error("Expected ", n2, ", got ", n1.next)
  }
}