package grind75

// type Node struct {
// 	Val       int
// 	Neighbors []Node
// }

// var st = []Node{}

// func push(n *Node) {
// 	st = append(st, n)
// }
// func pop() Node {
// 	len = len(st)
// }

// }

//   Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	hm := map[int]*Node{}
	newNode := helper(node, hm)
	return newNode

}

func helper(node *Node, hm map[int]*Node) *Node {
	if _, ok := hm[node.Val]; ok {
		return hm[node.Val]
	}

	new_node := &Node{node.Val, []*Node{}}
	hm[node.Val] = new_node
	for i := range node.Neighbors {
		new_node.Neighbors = append(new_node.Neighbors, helper(node.Neighbors[i], hm))
	}
	return new_node
}
