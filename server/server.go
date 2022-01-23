package server

import (
	"algorithms-go/amazon"
	"algorithms-go/util"
	"fmt"
)

func RunAmazon() {
	fmt.Println("\n******** Amazon ********")
	var root *util.TreeNode
	root = &util.TreeNode{Val: 10}
	root.Left = &util.TreeNode{Val: 6}
	root.Right = &util.TreeNode{Val: 20}
	root.Left.Left = &util.TreeNode{Val: 1}
	root.Left.Right = &util.TreeNode{Val: 8}
	root.Right.Left = &util.TreeNode{Val: 11}
	root.Right.Right = &util.TreeNode{Val: 23}
	fmt.Printf("LevelOrder:%v\n", amazon.LevelOrder(root))
	fmt.Printf("GetRow:%v\n", amazon.GetRow(3))
	fmt.Printf("MaxProfit:%v\n", amazon.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
}
