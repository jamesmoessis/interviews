package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	col1 := "collection1"
	col2 := "collection2"
	col3 := "collection3"
	files := []*file{
		&file{
			"file1.txt",
			1,
			&col1,
		},
		&file{
			"file2.txt",
			3,
			&col1,
		},
		&file{
			"file3.txt",
			3,
			&col2,
		},
		&file{
			"file4.txt",
			1,
			&col3,
		},
		&file{
			"file5.txt",
			2,
			&col2,
		},
	}

	assert.Equal(t, report{
		totalSize: 10,
		topNCols:  []string{col2, col1},
	}, *generateReport(files, 2))
}
