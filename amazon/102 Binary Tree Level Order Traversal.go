package amazon

import (
	"algorithms-go/util"
	"math"
)

func LevelOrder(root *util.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	h := height(root)
	result := make([][]int, int(h))
	for i := 0; i < int(h); i++ {
		temp := &[]int{}
		printLevel(root, i, temp)
		result[i] = *temp
	}
	return result
}

func height(node *util.TreeNode) float64 {
	if node == nil {
		return 0
	}
	return 1 + math.Max(height(node.Left), height(node.Right))
}

func printLevel(node *util.TreeNode, level int, result *[]int) {
	if node == nil {
		return
	}
	if level == 0 {
		*result = append(*result, node.Val)
		return
	}
	printLevel(node.Left, level-1, result)
	printLevel(node.Right, level-1, result)
}
