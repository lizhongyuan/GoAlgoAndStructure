package list


import "fmt"

type Node struct {
  data int
  next *Node
}

func ShowList() {
  var node1 Node
  var node2 Node
  node1.data = 1
  node2.data = 2
  node1.next = &node2

  fmt.Printf("node1's data:%d\n", node1.data)
  fmt.Printf("node2's data:%d\n", node1.next.data)
}
