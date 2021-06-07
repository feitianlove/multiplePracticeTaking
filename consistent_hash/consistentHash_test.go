package consistent_hash

import (
	"fmt"
	"testing"
)

func TestConsistent_AddNode(t *testing.T) {
	c := NewConsistentHash(nil)
	c.AddNode([]Node{
		{
			IPAddr:   "127.0.0.1",
			replicas: 2,
		},
		{
			IPAddr:   "127.0.0.2",
			replicas: 2,
		},
		{
			IPAddr:   "127.0.0.3",
			replicas: 2,
		},
		{
			IPAddr:   "127.0.0.4",
			replicas: 2,
		},
	})
	fmt.Println(c.Nodes, "Nodes")
	fmt.Println(c.HashMap, "hashMap")
	fmt.Println(c.GetNode("ftfeng1"))
	fmt.Println("===============")
	_ = c.ChangeWeight(Node{
		IPAddr:   "127.0.0.1",
		replicas: 0,
	})
	_ = c.RemoveNode("127.0.0.2")
	fmt.Println(c.Weight)
	fmt.Println(c.Nodes, "Nodes")
	fmt.Println(c.HashMap, "hashMap")
}
