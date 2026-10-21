// Package algo реализует структуры LIFO, FIFO, BFS, DFS
package algo

type Graph map[string][]string

func BFS(g Graph, start, target string) ([]string, int) {
	queue := []string{start}
	visited := map[string]bool{start: true}
	parent := make(map[string]string)
	visitedCount := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visitedCount++

		if curr == target {
			return reconstructPath(parent, start, target), visitedCount
		}

		for _, neighbor := range g[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = curr
				queue = append(queue, neighbor)
			}
		}
	}

	return nil, visitedCount
}

func DFS(g Graph, start, target string) ([]string, int) {
	visited := make(map[string]bool)
	parent := make(map[string]string)
	visitedCount := 0

	var dfs func(curr string) bool
	dfs = func(curr string) bool {
		visitedCount++
		if curr == target {
			return true
		}
		visited[curr] = true
		for _, neighbor := range g[curr] {
			if !visited[neighbor] {
				parent[neighbor] = curr
				if dfs(neighbor) {
					return true
				}
			}
		}
		return false
	}

	if dfs(start) {
		return reconstructPath(parent, start, target), visitedCount
	}
	return nil, visitedCount
}

func reconstructPath(parent map[string]string, start, target string) []string {
	path := []string{}
	curr := target

	for curr != "" {
		path = append(path, curr)
		if curr == start {
			break
		}
		curr = parent[curr]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
