package week209

func isEvenOddTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		if root.Val%2 == 0 {
			return false
		}
		return true
	}
	if root.Val%2 == 0 {
		return false
	}
	stack := make([]*TreeNode, 0)
	if root.Left != nil {
		stack = append(stack, root.Left)
	}
	if root.Right != nil {
		stack = append(stack, root.Right)
	}
	level := 1
	var node *TreeNode
	for len(stack) > 0 {
		length := len(stack)
		var pre int
		if level%2 == 0 {
			pre = 0
		} else {
			pre = 10 * 10 * 10 * 10 * 10 * 10
		}
		for i := 0; i < length; i++ {
			node, stack = stack[0], stack[1:]
			if node.Val%2 == 0 {
				if level%2 == 0 {
					// fmt.Println("even level even value", node.Val, level)
					return false
				}
			} else if node.Val%2 == 1 {
				if level%2 == 1 {
					// fmt.Println("odd level odd value", node.Val, level)
					return false
				}
			}
			if level%2 == 0 {
				if node.Val <= pre {
					return false
				}
			} else {
				if node.Val >= pre {
					return false
				}
			}
			pre = node.Val
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		level++
	}
	return true
}
