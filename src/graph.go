package main

import (
	"fmt"
)

type Node struct {
	id        int
	neighbors map[string][]*Node
}

type Graph struct {
	nodes    map[int]*Node
	count    int
	start    *Node
	terminal *Node
	current  *Node
	inputs   Set
}

func (n *Node) addNeighbor(symbol string, dest *Node) {
	n.neighbors[symbol] = append(n.neighbors[symbol], dest)
}

func (n Node) transitionOn(input string) []*Node {
	return n.neighbors[input]
}

func (n Node) printNode() {
	fmt.Printf("%d", n.id)
	if len(n.neighbors) == 0 {
		fmt.Print("\tEMPTY\n")
	} else {
		for i := range n.neighbors {
			fmt.Print("\t[")
			if i == "" {
				fmt.Print("ε")
			} else {
				fmt.Print(i)
			}
			fmt.Print(" -> ")
			for j := 0; j < len(n.neighbors[i]) - 1; j++ {
				fmt.Printf("%d, ", n.neighbors[i][j].id)
			}
			fmt.Printf("%d]\n", n.neighbors[i][len(n.neighbors[i]) - 1].id)
		}
	}
}

func initGraph() Graph {
	return Graph{make(map[int]*Node), 0, nil, nil, nil, initSet()}
}

func (g *Graph) addNode() int {
	g.count += 1
	n := Node{g.count, make(map[string][]*Node)}
	g.nodes[n.id] = &n

	return n.id
}

func (g *Graph) addEdge(src int, dest int, symbol string) {
	g.nodes[src].addNeighbor(symbol, g.nodes[dest])
	if symbol != "" {
		g.inputs.add(symbol)
	}
}

func (g *Graph) setStart(id int) {
	g.start = g.nodes[id]
	g.current = g.start
}

func (g Graph) findTerminal() int {
	for node := range g.nodes {
		if len(g.nodes[node].neighbors) == 0 {
			return node
		}
	}

	return 0
}

func (g *Graph) setTerminal(id int) {
	g.terminal = g.nodes[id]
}

func (g *Graph) transition(symbol string) {
	next, present := g.current.neighbors[symbol]

	if present {
		g.current = next[0]
	} else {
		fmt.Printf("No transition on %s from state %d\n", symbol, g.current.id)
	}
}
