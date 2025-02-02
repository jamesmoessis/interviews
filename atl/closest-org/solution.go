package closest_org

import "math"

/*
            Company
          /           \
        HR               Engg
      /   \          /          \
     Mona  Springs  BE          FE
                    /  \        / \
                  Alice Bob  Lisa Marley


Example calls from above structure :

getCommonGroupForEmployees (
  currentGroup = Company,
  targetEmployees = listOf(Lisa,Marley)
) -> FE
getCommonGroupForEmployees (
  currentGroup = Company,
  targetEmployees = listOf(Alice, Marley)
) -> Engg
getCommonGroupForEmployees (
  currentGroup = Company,
  targetEmployees = listOf(Mona,Lisa)
) -> Company


data class GroupNode (
    val groupName: String,
    val subGroups : Set<GroupNode> = emptySet(),
    var employeesInGroup : Set<Employee> = emptySet(),
)
*/

type Set[T comparable] map[T]bool

type GroupNode struct {
	Name      string
	SubGroups Set[*GroupNode]
	Employees Set[Employee]
}

type Employee string

func GetCommonGroupForEmployees(root *GroupNode, targets Set[Employee]) string {
	// DFS or BFS would work
	// Find paths of all targets as slice of nodes and then take the common start of them all
	// Start by searching for each one to make it easy
	paths := make([][]*GroupNode, 0, len(targets))
	for target := range targets {
		path := dfs(root, target)
		if path == nil {
			return "" // couldn't find
		}
		paths = append(paths, path)
	}

	commonStart := findCommonStart(paths)
	return commonStart[len(commonStart)-1].Name
}

func findCommonStart[T comparable](arrs [][]T) []T {
	lens := make([]int, 0, len(arrs))
	for _, path := range arrs {
		lens = append(lens, len(path))
	}
	shortest := minOf(lens...)

	if shortest == 0 {
		return []T{}
	}

	for j := 0; j < shortest; j++ {
		for i := 0; i < len(arrs)-1; i++ {
			if arrs[i][j] != arrs[i+1][j] {
				return arrs[i][:j]
			}
		}
	}
	return arrs[0][:shortest]
}

func minOf(nums ...int) int {
	minimum := math.MaxInt
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

func dfs(root *GroupNode, target Employee) []*GroupNode {
	path := []*GroupNode{root}
	// base case, check employee in this group, if it is return this node
	_, ok := root.Employees[target]
	if ok {
		return path
	}

	// check child groups
	for child := range root.SubGroups {
		subPath := dfs(child, target)
		if subPath != nil {
			return append(path, subPath...)
		}
	}

	// not found
	return nil
}
