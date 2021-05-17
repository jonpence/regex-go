package main

import (
	"fmt"
)

type Node struct {
	id        int
	neighbors map[string][]*Node
}

type Graph struct {
	nodes   map[int]*Node
	count   int
	start   *Node
	current *Node
}

func (n *Node) addNeighbor(symbol string, dest *Node) {
	n.neighbors[symbol] = append(n.neighbors[symbol], dest)
}

func (g *Graph) addEdge(src int, dest int, symbol string) {
	g.nodes[src].addNeighbor(symbol, g.nodes[dest])
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

func (g *Graph) addNode() int {
	g.count += 1
	n := Node{g.count, make(map[string][]*Node)}
	g.nodes[n.id] = &n

	return n.id
}

func (g *Graph) setStart(id int) {
	g.start = g.nodes[id]
	g.current = g.start
}

func (g *Graph) transition(symbol string) {
	next, present := g.current.neighbors[symbol]

	if present {
		g.current = next[0]
	} else {
		fmt.Printf("No transition on %s from state %d\n", symbol, g.current.id)
	}
}
