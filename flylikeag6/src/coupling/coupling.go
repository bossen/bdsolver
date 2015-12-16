package coupling

import (
	"log"
    "utils"
)

type Node struct {
	S, T       int
	Visited    bool
	Adj        [][]*Edge
	Succ       []*Node
	BasicCount int
}

type Edge struct {
	To    *Node
	Prob  float64
	Basic bool
}

type Coupling struct {
	Nodes []*Node
}

func New() Coupling {
	c := Coupling{}
	c.Nodes = make([]*Node, 0)
	return c
}

func FindNode(u, v int, c *Coupling) *Node {
	for _, n := range c.Nodes {
		if n.S == u && n.T == v {
			return n
		}
	}
	return nil
}

func Reachable(root *Node) []*Node {
	// Using slices might be slow. If we got performance problems we might
	// implement using lists instead.
	var reachables []*Node
	
	// Adding itself to reachables
	reachables = append(reachables, root)

	if root.Adj == nil {
		return reachables
	}
	
	root.Visited = true

	// Find all reachables from the  u,v node
	reachables = visit(root, reachables)

	for _, n := range reachables {
		n.Visited = false
	}

	return reachables
}

func visit(root *Node, results []*Node) []*Node {
	// log.Printf("%s, %s", root.S, root.T)
	if root.Adj == nil {
		return results
	}
	
	for i := range root.Adj {
		for j := range root.Adj[0] {
			edge := root.Adj[i][j]
			toVisit := edge.To
			
			if !edge.Basic || toVisit.Visited {
				continue
			} else if !utils.ApproxEqual(edge.Prob, 0) {
				toVisit.Visited = true
				results = append(results, toVisit)
				results = visit(toVisit, results)
			}	
		}
	}

	return results
}

func IsNodeInSlice(n *Node, nodes []*Node) bool {
	for _, succNode := range nodes {
		if succNode == n {
			return true
		}
	}
	return false
}

func DeleteNodeInSlice(n *Node, nodes *[]*Node) {
	for i := 0; i < len(*nodes); i++ {
		if (*nodes)[i] == n {
			// https://github.com/golang/go/wiki/SliceTricks
			(*nodes)[i] = (*nodes)[len(*nodes)-1]
			(*nodes)[len(*nodes)-1] = nil
			(*nodes) = (*nodes)[:len(*nodes)-1]
			break
		}
	}
	return
}



func RecoverBasicNodes(node *Node) {
	// the first basic node we have not traversed yet, such that we try and connect it with the other basic nodes
	firstbasic := findFirstBasic(node)
	
	// find all traversable to start the basic node recovery from
	t := findAllTraversableBasic(node, firstbasic, []utils.IntPair{})
	
	recoverBasicNodesRecursive(node, t)
}

func recoverBasicNodesRecursive(node *Node, traversed []utils.IntPair) {
	// create a copy such that we avoid out of bounds errors as we iterate
	t := traversed
	
	for _, pair := range t {
		// find index immediately right to the current index as i and j
		i, j := pair.I, (pair.J + 1) % len(node.Adj[0])
		
		if node.Adj[i][j].Basic {
			continue
		}
		
		// assuming index (i,j) is basic, we see what we can reach if it was basic
		tprime := findAllTraversableBasic(node, utils.IntPair{i, j}, traversed)
		
		if len(tprime) == len(traversed) {
			// if we could not more more basic cells, we do not improve anything by setting (i,j) to basic
			continue
		}
		
		// otherwise, we set it to basic and make appropiate updates, especially updating traversed
		log.Printf("setting cell with index (%v,%v) in the matching for node (%v,%v)", i, j, node.S, node.T)
		node.Adj[i][j].Basic = true
		node.BasicCount++
		tprime = append(tprime, utils.IntPair{i, j})
		traversed = tprime
		
		if len(traversed) == len(node.Adj) + (len(node.Adj[0]) - 1) {
			// if we have the correct number of nodes, we can now traverse all basic nodes, and we terminate
			return
		}
	}
	recoverBasicNodesRecursive(node, traversed)
	return
}


func findFirstBasic(node *Node) utils.IntPair {
    for i, row := range node.Adj {
        for j, edge := range row {
            if edge.Basic {
                return utils.IntPair{i, j}
            }
        }
    }
    return utils.IntPair{-1, -1}
}

func findAllTraversableBasic(node *Node, curr utils.IntPair, basicsfound []utils.IntPair) []utils.IntPair {
    if node.Adj[curr.I][curr.J].Basic {
        // if the current node is basic, add it to the result set
        // depending on where we call this function, we may or may not want to add the current node, this handles that
        basicsfound = append(basicsfound, curr)
    }
    
    basicsfound = traverseHorizontal(node, basicsfound, curr.I)
    basicsfound = traverseVertical(node, basicsfound, curr.J)
    
    return basicsfound
}

func traverseHorizontal(node *Node, basicsfound []utils.IntPair, i int) []utils.IntPair {
    for j := range node.Adj[0] {
        pair := utils.IntPair{i, j}
        
        if node.Adj[i][j].Basic && !utils.IsIntPairInSlice(pair, basicsfound) {
            basicsfound = append(basicsfound, pair)
        
            basicsfound = traverseVertical(node, basicsfound, j)
        }
    }
    
    return basicsfound
}

func traverseVertical(node *Node, basicsfound []utils.IntPair, j int) []utils.IntPair {
    for i := range node.Adj {
        pair := utils.IntPair{i, j}
        
        if node.Adj[i][j].Basic && !utils.IsIntPairInSlice(pair, basicsfound) {
            basicsfound = append(basicsfound, pair)
        
            basicsfound = traverseHorizontal(node, basicsfound, i)
        }
    }
    
    return basicsfound
}
