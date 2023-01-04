package fuzalianbiao

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		 return nil
	}
	m := make(map[*Node]*Node)
	for cur := head ; cur != nil ; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}
	// 新节点组装
	for cur := head ; cur != nil ; cur = cur.Next {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
	}
     return m[head]
}