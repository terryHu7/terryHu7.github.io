package main

import "fmt"

type DAG struct {
	Vertexes []*Vertex
}

type Vertex struct {
	Value interface{}

	Parent  []*Vertex
	Childen []*Vertex
}

func (d *DAG) AddEdge(from, to *Vertex) {
	from.Childen = append(from.Childen, to)
	to.Parent = append(to.Parent, from)
}

//undone
func main() {
	var dag DAG
	va := &Vertex{Value: "a"}
	vb := &Vertex{Value: "b"}
	vc := &Vertex{Value: "c"}
	vd := &Vertex{Value: "d"}
	dag.AddEdge(va, vb)
	dag.AddEdge(va, vc)
	dag.AddEdge(vb, vc)
	dag.AddEdge(vb, vd)
	dag.AddEdge(vc, vd)
	//nil dag
	for _, v := range dag.Vertexes {
		fmt.Printf("%v->", v.Value)
	}
	//https://toutiao.io/posts/keetjdp/preview

}
