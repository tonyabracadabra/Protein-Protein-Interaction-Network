package main

import (
	"fmt"
	"math"
	"strconv"
	"bufio"
	"os"
)

var distTo []float64
var edgeTo []Edge
var pq MinPQ

func dijkstra(G Graph, s int) {
    for _, e := range G.Edges() {
        if e.Weight() < 0 {
        	fmt.Println("edge " + strconv.Itoa(e.From()) + ","+ strconv.Itoa(e.To()) + "has negative weight")
        }
    }

	for v := 0; v < G.V(); v++ {
		distTo[v] = math.MaxFloat64
	}
	distTo[s] = 0.0
	edgeTo[s] = Edge{w:-1}

	// relax vertices in order of distance from s

	source := &Node {
        label: s,
        distTo: distTo[s],
    }
	pq.Push(source)

	for len(pq) != 0 {
	    v := pq.Pop().(*Node).label
	    for _, e := range G.Adj(v) {
	    	relax(e)
	    }
	}

}

// relax edge e and update pq if changed
func relax(e Edge) {
    v, w := e.From(), e.To()
    if distTo[w] > distTo[v] + e.Weight() {
		distTo[w] = distTo[v] + e.Weight()
		edgeTo[w] = e

		node := &Node {
	        label: w,
	        distTo: distTo[w],
	    }

	    ifContains := false
	    for _, n := range pq {
	    	if n.label == w {
	    		pq.update(node, n.index)
	    		ifContains = true
	    	}
	    }

		if ifContains == false {
			pq.Push(node)
		}

	}
}

// Check if there is a path from s to v
func hasPathTo(v int) bool {
    return distTo[v] < math.MaxFloat64
}

// Returns a shortest path from the source vertex to vertex v
func pathTo(v int) Stack {
    if (!hasPathTo(v)) {
    	return []Edge{}
    }

    e := edgeTo[v]
    var path Stack
    for e.w != -1  {
    	path.Push(e)
		e = edgeTo[e.From()]
    }

    return path
}

func writeResult(protein1, protein2, temp string) {
	fileName := protein1 + "_to_" + protein2 + "interaction_path.txt"
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Fprintln(f, temp)
	fmt.Println("Successfully created the file ", fileName)
}

func RunDijkstra(num int, edges []Edge, gnMap, pnMap map[string]int, proteins []Protein) {
	distTo = make([]float64, num)
	edgeTo = make([]Edge, num)
	graph := CreateGraph(num, edges)

	// Read inquiry protein info from standard input
	readProtein := bufio.NewReader(os.Stdin)
	fmt.Print("Input the first protein name: ")
	protein1, _ := readProtein.ReadString('\n')
	fmt.Print("Input the second protein name: ")
	protein2, _ := readProtein.ReadString('\n')

	// Delete the tailing \n of the input
	protein1, protein2 = protein1[:6], protein2[:6]
	dijkstra(graph, pnMap[protein1])
	p := pathTo(pnMap[protein2]) // Return the shortest path between protein1 and protein2

	if len(p) == 0 {
		fmt.Println("The possible interaction path between ", protein1, "and ", protein2, "is not found!")
		return
	}

	// Extract and print out path from the stack
	var temp string
	var e Edge
	for len(p) != 0  {
		e = p.Pop()
		temp += (proteins[e.v].label + " -> ")
	}
	temp += proteins[e.w].label
	temp = "The possible protein interaction chain in this system is, \n" + temp

	fmt.Println(temp)

	writeResult(protein1, protein2, temp)
}
