package main

// main.go contains code to generate SageMath graphs of the Jute orderings of
// various graphs. None of the core code for the Jute ordering algorithm is in
// this file. See graph.go, then addnode.go, then finally ordering.go.

import (
	"fmt"
	"strconv"
	"strings"
)

// RelativeOrdering sorts the graph using the supplied node as the tip, then
// prints the resulting ordering.
func (gn *GraphNode) RelativeOrdering() string {
	relativeOrdering := gn.relativeOrdering()
	s := fmt.Sprint(relativeOrdering[0].name)
	for i := 1; i < len(relativeOrdering); i++ {
		s = fmt.Sprint(s, "-")
		s = fmt.Sprint(s, relativeOrdering[i].name)
	}
	return s
}

// SageGen returns a string that can be fed into Sage to create a visualization
// of the longest chain and the votes for each edge in that chain.
func (g *Graph) SageGen(tip *GraphNode) string {
	s := fmt.Sprintln("G = DiGraph()")
	for edge, weight := range tip.relativeVoteGraph {
		// Parse the edge name into its components.
		nodes := strings.Split(string(edge), "-")
		s = fmt.Sprint(s, "G.add_edge("+nodes[0]+", "+nodes[1]+", "+strconv.Itoa(int(weight))+")\n")
	}
	relativeOrdering := tip.RelativeOrdering()
	s = fmt.Sprint(s, "H = G.plot(edge_labels=True, layout='acyclic', edge_color='grey')\n")
	s = fmt.Sprint(s, "H.show(title=\""+relativeOrdering+"\", figsize=(5,16))\n")
	filename := relativeOrdering + ".png"
	s = fmt.Sprintf("%sH.save(filename=\"/home/user/plots/%s\", title=\""+relativeOrdering+"\", figsize=(5,16))\n", s, filename)
	return s
}

// Build some graphs, and then print some code that can be used to generate the
// graphs in SageMath.
func main() {
	// Diamond Graph
	gDiamond := NewGraph()
	d1 := gDiamond.CreateNode(gDiamond.GenesisNode())
	d2 := gDiamond.CreateNode(gDiamond.GenesisNode())
	d3 := gDiamond.CreateNode(d1, d2)
	fmt.Printf("\n# Diamond Graph\n%s", gDiamond.SageGen(d3))

	// Pentagon Graph
	gPentagon := NewGraph()
	p1 := gPentagon.CreateNode(gPentagon.GenesisNode())
	p2 := gPentagon.CreateNode(gPentagon.GenesisNode())
	p3 := gPentagon.CreateNode(p1)
	p4 := gPentagon.CreateNode(p2, p3)
	fmt.Printf("\n# Pentagon Graph\n%s", gPentagon.SageGen(p4))

	// Double Diamond Graph
	gDDiamond := NewGraph()
	dd1 := gDDiamond.CreateNode(gDDiamond.GenesisNode())
	dd2 := gDDiamond.CreateNode(gDDiamond.GenesisNode())
	dd3 := gDDiamond.CreateNode(dd1, dd2)
	dd4 := gDDiamond.CreateNode(dd2)
	dd5 := gDDiamond.CreateNode(dd3, dd4)
	fmt.Printf("\n# Double Diamond Graph\n%s", gDDiamond.SageGen(dd5))

	// Nested Diamond Graph
	ndg := NewGraph()
	nd1 := ndg.CreateNode(ndg.GenesisNode())
	nd2 := ndg.CreateNode(ndg.GenesisNode())
	nd3 := ndg.CreateNode(nd1)
	nd4 := ndg.CreateNode(nd1, nd2)
	nd5 := ndg.CreateNode(nd2)
	nd6 := ndg.CreateNode(nd3, nd4)
	nd7 := ndg.CreateNode(nd4, nd5)
	nd8 := ndg.CreateNode(nd6, nd7)
	nd9 := ndg.CreateNode(nd8)
	nd10 := ndg.CreateNode(nd9)
	nd11 := ndg.CreateNode(nd10)
	fmt.Printf("\n# Nested Diamond Graph\n%s", ndg.SageGen(nd11))

	// Ongoing Graph.
	og := NewGraph()
	o1 := og.CreateNode(og.GenesisNode())
	o2 := og.CreateNode(og.GenesisNode())
	o3 := og.CreateNode(og.GenesisNode())
	o4 := og.CreateNode(og.GenesisNode())
	o5 := og.CreateNode(og.GenesisNode())
	o11 := og.CreateNode(o1, o2)
	o12 := og.CreateNode(o1, o2, o3)
	o13 := og.CreateNode(o2, o3, o4)
	o14 := og.CreateNode(o3, o4, o5)
	o15 := og.CreateNode(o4, o5, o5)
	o21 := og.CreateNode(o11, o12)
	o22 := og.CreateNode(o11, o12, o13)
	o23 := og.CreateNode(o12, o13, o14)
	o24 := og.CreateNode(o13, o14, o15)
	o25 := og.CreateNode(o14, o15, o15)
	o31 := og.CreateNode(o21, o22)
	o32 := og.CreateNode(o21, o22, o23)
	o33 := og.CreateNode(o22, o23, o24)
	o34 := og.CreateNode(o23, o24, o25)
	o35 := og.CreateNode(o24, o25, o25)
	o41 := og.CreateNode(o31, o32)
	o42 := og.CreateNode(o31, o32, o33)
	o43 := og.CreateNode(o32, o33, o34)
	o44 := og.CreateNode(o33, o34, o35)
	o45 := og.CreateNode(o34, o35, o35)
	o51 := og.CreateNode(o41, o42, o43, o44, o45)
	fmt.Printf("\n# Ongoing Graph\n%s", og.SageGen(o51))

	// Impossibility Proof Graph
	ip := NewGraph()
	ip1 := ip.CreateNode(ip.GenesisNode())
	ip2 := ip.CreateNode(ip.GenesisNode())
	ip3 := ip.CreateNode(ip1)
	ip4 := ip.CreateNode(ip2)
	ip5 := ip.CreateNode(ip3)
	ip6 := ip.CreateNode(ip4, ip5)
	ip7 := ip.CreateNode(ip4)
	ip8 := ip.CreateNode(ip7)
	ip9 := ip.CreateNode(ip8)
	ip10 := ip.CreateNode(ip9)
	ip11 := ip.CreateNode(ip10, ip6)
	fmt.Printf("\n# Impossibility Proof Graph\n%s", ip.SageGen(ip11))

	// Abstain Graph
	a := NewGraph()
	a1 := a.CreateNode(a.GenesisNode())
	a2 := a.CreateNode(a1)
	a3 := a.CreateNode(a2)
	a4 := a.CreateNode(a3)
	a5 := a.CreateNode(a4)
	a6 := a.CreateNode(a.GenesisNode())
	a7 := a.CreateNode(a6)
	a8 := a.CreateNode(a7)
	a9 := a.CreateNode(a8)
	a10 := a.CreateNode(a9)
	a11 := a.CreateNode(a10)
	a12 := a.CreateNode(a11)
	a13 := a.CreateNode(a12)
	a14 := a.CreateNode(a13)
	a15 := a.CreateNode(a14)
	a16 := a.CreateNode(a15)
	a17 := a.CreateNode(a16)
	a18 := a.CreateNode(a17, a5)
	fmt.Printf("\n# Abstain Graph\n%s", a.SageGen(a18))

	// Leech Graph
	l := NewGraph()
	l1 := l.CreateNode(l.GenesisNode())
	l2 := l.CreateNode(l1)
	l3 := l.CreateNode(l2)
	l4 := l.CreateNode(l3)
	l5 := l.CreateNode(l4)
	l6 := l.CreateNode(l5)
	l7 := l.CreateNode(l6)
	l8 := l.CreateNode(l7)
	l9 := l.CreateNode(l8)
	la := l.CreateNode(l9)
	lb := l.CreateNode(la)
	lc := l.CreateNode(lb)
	ld := l.CreateNode(l.GenesisNode())
	le := l.CreateNode(ld, l2)
	lf := l.CreateNode(le, l3)
	lg := l.CreateNode(lf, l5)
	lh := l.CreateNode(lg, l7)
	li := l.CreateNode(lh, l8)
	lj := l.CreateNode(li, l9)
	lk := l.CreateNode(lj, lb)
	lm := l.CreateNode(lk, lc)
	fmt.Printf("\n# Leech Graph\n%s", l.SageGen(lm))

	// Low Latency Adversary Graph
	lla := NewGraph()
	lla1 := lla.CreateNode(lla.GenesisNode())
	lla2 := lla.CreateNode(lla.GenesisNode())
	lla3 := lla.CreateNode(lla.GenesisNode())
	lla4 := lla.CreateNode(lla.GenesisNode())
	lla5 := lla.CreateNode(lla.GenesisNode())
	lla6 := lla.CreateNode(lla5)
	lla7 := lla.CreateNode(lla1, lla2, lla3, lla4, lla6)
	lla8 := lla.CreateNode(lla6)
	lla9 := lla.CreateNode(lla8)
	lla10 := lla.CreateNode(lla7, lla9)
	fmt.Printf("\n# Low Latenct Adversary Graph\n%s", lla.SageGen(lla10))
}
