package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if _, err := f(node.Data); err != nil {
			panic(err)
		}
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
}
