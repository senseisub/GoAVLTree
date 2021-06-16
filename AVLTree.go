package main

import(
	"bufio"
  	"fmt"
  	"os"
  	"strings"
	"strconv"
)

type Generic interface {
	GetData()
}

type TreeNode struct {
	data int
	left *TreeNode
	right *TreeNode
	height int
}

type AVLTree struct {
	root *TreeNode
	size int

}

// height of the tree
func height(currentNode *TreeNode) int{
	if currentNode == nil {
		return 0;
	}
	return currentNode.height;
}

// A utility function to get maximum
// of two integers
func max(a, b int) int{
	if a > b {
		return a
	} 
	return b;
}

func (tree *AVLTree) Insert(curr *TreeNode, x int) {
	if curr == nil {
		fmt.Println("insert <nul>");
		tree.root = &TreeNode{data : x}
		curr = tree.root
		fmt.Println(tree.root);
	} else if curr.data > x {
		if curr.left == nil {
			curr.left = &TreeNode{data : x}
			if height(curr.left) - height(curr.right)  > 1 {
				if x < curr.left.data {
					case1(curr);
				} else {
					case2(curr);
				}
			}
		} else {
			tree.Insert(curr.left, x)
		}
	} else if curr.data < x {
		if curr.right == nil {
			curr.right = &TreeNode{data : x}
			if height(curr.right) - height(curr.left)  > 1 {
				if x > curr.right.data {
					case4(curr);
				} else {
					case3(curr);
				}
			}
		} else {
			tree.Insert(curr.right, x)
		}
	} else {
		return
	}
	tree.size++
	curr.height = max(height(curr.left), height(curr.right))+1;
}

//prints the values in the AVLTree in order
func (tree *AVLTree) Inorder(node *TreeNode){
	if node == nil {
		return;
	}
	tree.Inorder(node.left);
	fmt.Println(node.data); // TODO: getID works but not what we want
	tree.Inorder(node.right);
}

func (tree *AVLTree) Exists(curr *TreeNode, x int) bool {
	if curr == nil {
		fmt.Println("<nul>");
		return false
	}
	if curr.data > x{
		return tree.Exists(curr.left, x)
	} else if curr.data < x {
		return tree.Exists(curr.right, x)
	} 
	return true;
}

//left left rotation
func case1(k2 *TreeNode){
	k1 := k2.left;
	k2.left = k1.right;
	k1.right = k2;
	k2.height = max(height(k2.left), height(k2.right)) + 1;
	k1.height = max(height(k1.left), k2.height) + 1;

	k2 = k1;
}

//right right rotation
func case4(k1 *TreeNode){
	k2 := k1.right;
	k1.right = k2.left;
	k2.left = k1;
	k1.height = max(height(k1.left), height(k1.right)) + 1;
	k2.height = max(height(k2.left), k1.height) + 1;

	k1 = k2;
}

//left right rotation
func case2(k3 *TreeNode){
	case4(k3.left);
	case1(k3);
}

//right left rotation
func case3(k1 *TreeNode){
	case1(k1.right);
	case4(k1);
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
			fmt.Println(tree.Exists(tree.root, i));
		} else {
			break
		}
	}
	tree.Inorder(tree.root)
}