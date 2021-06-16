package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Generic interface {
	GetData()
}

type TreeNode struct {
	data   interface{}
	left   *TreeNode
	right  *TreeNode
	height int
}

type AVLTree struct {
	root *TreeNode
	size int
}

// height of the tree
func height(currentNode *TreeNode) int {
	if currentNode == nil {
		return 0
	}
	return currentNode.height
}

// A utility function to get maximum
// of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (node *TreeNode) GreaterThan(data2 interface{}) bool {
	data := node.data
	switch (node.data).(type) {
	case int:
		if data.(int) > data2.(int) {
			return true
		} else {
			return false
		}
	case int8:
		if data.(int8) > data2.(int8) {
			return true
		} else {
			return false
		}
	case int16:
		if data.(int16) > data2.(int16) {
			return true
		} else {
			return false
		}
	case int32:
		if data.(int32) > data2.(int32) {
			return true
		} else {
			return false
		}
	case int64:
		if data.(int64) > data2.(int64) {
			return true
		} else {
			return false
		}
	case uint8:
		if data.(uint8) > data2.(uint8) {
			return true
		} else {
			return false
		}
	case uint16:
		if data.(uint16) > data2.(uint16) {
			return true
		} else {
			return false
		}
	case uint32:
		if data.(uint32) > data2.(uint32) {
			return true
		} else {
			return false
		}
	case uint64:
		if data.(uint64) > data2.(uint64) {
			return true
		} else {
			return false
		}
	case float64:
		if data.(float64) > data2.(float64) {
			return true
		} else {
			return false
		}
	case float32:
		if data.(float32) > data2.(float32) {
			return true
		} else {
			return false
		}
	case string:
		if data.(string) > data2.(string) {
			return true
		} else {
			return false
		}
	}
	return false
}

func (node *TreeNode) LesserThan(data2 interface{}) bool {
	data := node.data
	switch (node.data).(type) {
	case int:
		if data.(int) < data2.(int) {
			return true
		} else {
			return false
		}
	case int8:
		if data.(int8) < data2.(int8) {
			return true
		} else {
			return false
		}
	case int16:
		if data.(int16) < data2.(int16) {
			return true
		} else {
			return false
		}
	case int32:
		if data.(int32) < data2.(int32) {
			return true
		} else {
			return false
		}
	case int64:
		if data.(int64) < data2.(int64) {
			return true
		} else {
			return false
		}
	case uint8:
		if data.(uint8) < data2.(uint8) {
			return true
		} else {
			return false
		}
	case uint16:
		if data.(uint16) < data2.(uint16) {
			return true
		} else {
			return false
		}
	case uint32:
		if data.(uint32) < data2.(uint32) {
			return true
		} else {
			return false
		}
	case uint64:
		if data.(uint64) < data2.(uint64) {
			return true
		} else {
			return false
		}
	case float64:
		if data.(float64) < data2.(float64) {
			return true
		} else {
			return false
		}
	case float32:
		if data.(float32) < data2.(float32) {
			return true
		} else {
			return false
		}
	case string:
		if data.(string) < data2.(string) {
			return true
		} else {
			return false
		}
	}
	return false
}

func (tree *AVLTree) Insert(curr *TreeNode, x int) {

	if curr == nil {
		fmt.Println("insert <nul>")
		tree.root = &TreeNode{data: x}
		curr = tree.root
		fmt.Println(tree.root)
	} else if curr.GreaterThan(x) {
		if curr.left == nil {
			curr.left = &TreeNode{data: x}
			if height(curr.left)-height(curr.right) > 1 {
				if curr.left.GreaterThan(x) {
					case1(curr)
				} else {
					case2(curr)
				}
			}
		} else {
			tree.Insert(curr.left, x)
		}
	} else if curr.LesserThan(x) {
		if curr.right == nil {
			curr.right = &TreeNode{data: x}
			if height(curr.right)-height(curr.left) > 1 {
				if curr.right.LesserThan(x) {
					case4(curr)
				} else {
					case3(curr)
				}
			}
		} else {
			tree.Insert(curr.right, x)
		}
	} else {
		return
	}
	tree.size++
	curr.height = max(height(curr.left), height(curr.right)) + 1
}

//prints the values in the AVLTree in order
func (tree *AVLTree) Inorder(node *TreeNode) {
	if node == nil {
		return
	}
	tree.Inorder(node.left)
	fmt.Println(node.data) // TODO: getID works but not what we want
	tree.Inorder(node.right)
}

func (tree *AVLTree) Exists(curr *TreeNode, x int) bool {
	if curr == nil {
		fmt.Println("<nul>")
		return false
	}
	if curr.GreaterThan(x) {
		return tree.Exists(curr.left, x)
	} else if curr.LesserThan(x) {
		return tree.Exists(curr.right, x)
	}
	return true
}

//left left rotation
func case1(k2 *TreeNode) {
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	k2.height = max(height(k2.left), height(k2.right)) + 1
	k1.height = max(height(k1.left), k2.height) + 1

	k2 = k1
}

//right right rotation
func case4(k1 *TreeNode) {
	k2 := k1.right
	k1.right = k2.left
	k2.left = k1
	k1.height = max(height(k1.left), height(k1.right)) + 1
	k2.height = max(height(k2.left), k1.height) + 1

	k1 = k2
}

//left right rotation
func case2(k3 *TreeNode) {
	case4(k3.left)
	case1(k3)
}

//right left rotation
func case3(k1 *TreeNode) {
	case1(k1.right)
	case4(k1)
}

func main() {
	tree := &AVLTree{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input values")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("q", text) != 0 {
			i, _ := strconv.Atoi(text)
			tree.Insert(tree.root, i)
		} else {
			break
		}

	}

	for {
		fmt.Print("Check Exists:  ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("q", text) != 0 {
			i, _ := strconv.Atoi(text)
			fmt.Println(tree.Exists(tree.root, i))
		} else {
			break
		}
	}
	tree.Inorder(tree.root)
}
