package main

import "fmt"

type Graph struct {
	nodes map[string]node
}

type node map[string]bool

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string]node)}
}

func (g *Graph) AddNode(name string) {
	if _, ok := g.nodes[name]; !ok {
		g.nodes[name] = make(node)
	}
}

func (g *Graph) AddEdge(from, to string) bool {
	f, ok := g.nodes[from]
	if !ok {
		return false
	}
	_, ok = g.nodes[to]
	if !ok {
		return false
	}
	f.addedge(to)
	return true
}

func (n node) addedge(to string) {
	n[to] = true
}

func (n node) edges() []string {
	var k []string
	for key := range n {
		k = append(k, key)
	}
	return k
}

func (g *Graph) topsort(start string) {
	result := NewOrderedSet()
	g.visit(start, result, nil)
	fmt.Println("order is :", result)
}

func (g *Graph) visit(n string, result *OrderedSet, visited *OrderedSet) error {
	if visited == nil {
		visited = NewOrderedSet()
	}

	already_added := visited.added(n)
	if already_added {
		fmt.Println("CYCLE found")
		return fmt.Errorf("Cycle Error")
	}
	node := g.nodes[n]
	for _, n := range node.edges() {
		err := g.visit(n, result, visited.copy())
		if err != nil {
			return err
		}
	}
	result.added(n)
	return nil
}

func (s *OrderedSet) copy() *OrderedSet {
	new := NewOrderedSet()
	for k, v := range s.idxs {
		new.idxs[k] = v
	}
	new.length = s.length
	return new
}

func (s *OrderedSet) added(node string) bool {
	_, ok := s.idxs[node]
	if !ok {
		s.idxs[node] = s.length
		s.length++
		s.items = append(s.items, node)
	}
	return ok
}

type OrderedSet struct {
	idxs   map[string]int
	items  []string
	length int
}

func NewOrderedSet() *OrderedSet {
	return &OrderedSet{idxs: make(map[string]int), items: []string{}, length: 0}
}

func main() {
	g := NewGraph()
	g.AddNode("a")
	g.AddNode("b")
	g.AddNode("c")

	g.AddEdge("a", "b")
	g.AddEdge("a", "c")
	g.AddEdge("b", "c")
	g.topsort("a")
	fmt.Println(g)
}
