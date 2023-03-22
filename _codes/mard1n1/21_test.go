package mard1n1

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  *int   `json:"age,omitempty"`
	IsA  bool   `json:"isA,omitempty"`
}

func TestXxx(t *testing.T) {
	p := Person{Name: "A", Age: func() *int { v := 0; return &v }(), IsA: true}

	jsonBytes, _ := json.Marshal(p)
	fmt.Println(string(jsonBytes))

	// var up Person
	// pj := `{"name":"A"}`
	// json.Unmarshal([]byte(pj), &up)
	// fmt.Printf("up:%+v\n", up)

}

/*
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return pTraversal(root, p, q)
}

func pTraversal(node, p, q *TreeNode) *TreeNode {
	var res *TreeNode

	if node == nil {
		return res
	}

	l := pTraversal(node.Left, p, q)
	r := pTraversal(node.Right, p, q)

	if node == p || node == q {
		return node
	}

	if l != nil && r != nil {
		return node
	}
	return nil
}
*/
