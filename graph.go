package main


type Graph struct {
    edges []Edge
    adj map[int][]Edge
    numofV int
}

func CreateGraph(numofV int, edges []Edge) Graph {
	adj := edgesToAdj(edges)
    return Graph{edges, adj, numofV}
}

func edgesToAdj(edges []Edge) map[int][]Edge {
	adj := make(map[int][]Edge)

	for _, e := range edges {
		adj[e.From()] = append(adj[e.From()], e)
		/* Interchange from and to nodes, and add it to the adjacency list
		   As this is the bi-directional graph */
		e.w, e.v = e.v, e.w
		adj[e.From()] = append(adj[e.From()], e)
	}

	return adj
}

func (g *Graph) V() int {
    return g.numofV
}

func (g *Graph) Adj(v int) []Edge {
    return g.adj[v]
}

func (g *Graph) Edges() []Edge {
    return g.edges
}