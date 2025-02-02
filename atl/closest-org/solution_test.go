package closest_org

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	root := &GroupNode{
		Name:      "company",
		Employees: make(Set[Employee]),
		SubGroups: Set[*GroupNode]{
			{
				Name: "HR",
				Employees: Set[Employee]{
					"Mona":    true,
					"Springs": true,
				},
			}: true,
			{
				Name:      "Engg",
				Employees: Set[Employee]{},
				SubGroups: Set[*GroupNode]{
					{
						Name: "BE",
						Employees: Set[Employee]{
							"Alice": true,
							"Bob":   true,
						},
					}: true,
					{
						Name: "FE",
						Employees: Set[Employee]{
							"Lisa":   true,
							"Marley": true,
						},
					}: true,
				},
			}: true,
		},
	}

	assert.Equal(t, "FE", GetCommonGroupForEmployees(root, Set[Employee]{"Lisa": true, "Marley": true}))
	assert.Equal(t, "company", GetCommonGroupForEmployees(root, Set[Employee]{"Alice": true, "Mona": true}))
}
