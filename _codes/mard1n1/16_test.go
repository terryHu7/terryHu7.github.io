package mard1n1

import (
	"fmt"
	"os/user"
	"testing"
)

func TestMerge(t *testing.T) {
	var userinfo, _ = user.Current()
	fmt.Println(userinfo.HomeDir + "/")
}

func mergeTreesCom(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 != nil && t2 == nil {
		return &TreeNode{
			Val:   t1.Val,
			Left:  mergeTrees(t1.Left, nil),
			Right: mergeTrees(t1.Right, nil),
		}
	}
	if t1 == nil && t2 != nil {
		return &TreeNode{
			Val:   t2.Val,
			Left:  mergeTrees(nil, t2.Left),
			Right: mergeTrees(nil, t2.Right),
		}
	}
	if t1 != nil && t2 != nil {
		return &TreeNode{
			Val:   t1.Val + t2.Val,
			Left:  mergeTrees(t1.Left, t2.Left),
			Right: mergeTrees(t1.Right, t2.Right),
		}
	}
	return nil
}

//simpilification
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}

}
