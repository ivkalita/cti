package task3

import "github.com/ivkalita/cti/internal/chapter2/structs"

// RemoveNode – O(1) memory, O(1) time
func RemoveNode(n *structs.Node) {
	n.Data = n.Next.Data
	n.Next = n.Next.Next
}
