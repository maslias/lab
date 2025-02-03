package graph

import (
	"fmt"
)

type verticeWeight struct {
	vertice *vertice
	weight  int
}

type vertice struct {
	value    int
	adjacent []*verticeWeight
}

type graph struct {
	vertices []*vertice
	path     []int
}

func (g *graph) addVertice(v *vertice) {
	g.vertices = append(g.vertices, v)
}

func (g *graph) printStruct() {
	for _, v := range g.vertices {
		fmt.Printf("%d: ", v.value)
		for _, a := range v.adjacent {
			fmt.Printf("to=%d w=%d | ", a.vertice.value, a.weight)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *graph) addEdgesToVerticeByValue(from int, to int, weight int, both bool) {
	vFrom, errFrom := g.getVerticeByValue(from)
	vTo, errTo := g.getVerticeByValue(to)

	if errFrom != nil || errTo != nil {
		return
	}

	vFrom.adjacent = append(vFrom.adjacent, &verticeWeight{vertice: vTo, weight: weight})

	if both == true {
		vTo.adjacent = append(vTo.adjacent, &verticeWeight{vertice: vFrom, weight: weight})
	}
}

func (g *graph) getVerticeByValue(value int) (*vertice, error) {
	for i, v := range g.vertices {
		if v.value == value {
			return g.vertices[i], nil
		}
	}

	return nil, fmt.Errorf("vertice with value %d not exists", value)
}

func containsVerticeByValue(vertices []*vertice, value int) bool {
	for _, v := range vertices {
		if v.value == value {
			return true
		}
	}

	return false
}

func (g *graph) dfs(start int, end int) {
	v, err := g.getVerticeByValue(start)
	if err != nil {
		return
	}

	seen := make([]bool, len(g.vertices))
	g.path = []int{}
	fmt.Println(seen)

	g.dfsProcedure(v, end, seen)
	fmt.Printf("path %v \n", g.path)
}

func (g *graph) dfsProcedure(v *vertice, end int, seen []bool) bool {
	fmt.Printf("dfsprocedure value: %d \n", v.value)
	if seen[v.value] == true {
		return false
	}
	seen[v.value] = true

	g.path = append(g.path, v.value)
	if v.value == end {
		return true
	}

	for _, a := range v.adjacent {
		if g.dfsProcedure(a.vertice, end, seen) {
			return true
		}
	}

	g.path = g.path[:len(g.path)-1]

	return false
}

func TestStructGraph() {
	g := graph{}
	g.addVertice(&vertice{value: 0})
	g.addVertice(&vertice{value: 1})
	g.addVertice(&vertice{value: 2})
	g.addVertice(&vertice{value: 3})
	g.addVertice(&vertice{value: 4})
	g.addVertice(&vertice{value: 5})

	g.printStruct()

	g.addEdgesToVerticeByValue(0, 1, 10, false)
	g.addEdgesToVerticeByValue(1, 2, 10, false)
	g.addEdgesToVerticeByValue(1, 5, 10, false)
	g.addEdgesToVerticeByValue(3, 4, 10, false)
	g.addEdgesToVerticeByValue(4, 5, 10, false)
	g.addEdgesToVerticeByValue(3, 5, 10, false)
	g.printStruct()

	g.dfs(0, 5)
}
