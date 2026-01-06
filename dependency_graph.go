package main

type DependencyGraph struct {
	graph map[string]map[string]bool
	roots map[string]bool
}

func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		graph: make(map[string]map[string]bool),
		roots: make(map[string]bool),
	}
}

func (dg *DependencyGraph) AddCall(caller, callee string) {
	if dg.graph[caller] == nil {
		dg.graph[caller] = make(map[string]bool)
	}
	dg.graph[caller][callee] = true
}

func (dg *DependencyGraph) MarkRoot(funcName string) {
	dg.roots[funcName] = true
}

func (dg *DependencyGraph) GetReachable() map[string]bool {
	reachable := make(map[string]bool)
	visited := make(map[string]bool)

	var dfs func(string)
	dfs = func(funcName string) {
		if visited[funcName] {
			return
		}
		visited[funcName] = true
		reachable[funcName] = true

		for callee := range dg.graph[funcName] {
			dfs(callee)
		}
	}

	for root := range dg.roots {
		dfs(root)
	}

	return reachable
}
