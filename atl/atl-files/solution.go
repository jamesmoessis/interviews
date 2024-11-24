package main

import "sort"

/**
Imagine we have a system that stores files, and these files can be grouped into collections.
We are interested in knowing where our resources are being taken up.

For this system we would like to generate a report that lists:
The total size of all files stored; and
The top N collections (by file size) where N can be a user-defined value

file1.txt (size: 100)
file2.txt (size: 200) in collection "collection1"
file3.txt (size: 200) in collection "collection1"
file4.txt (size: 300) in collection "collection2"
file5.txt (size: 10)
*/

func main() {

}

type file struct {
	name       string
	size       int
	collection *string
}

type report struct {
	totalSize int
	topNCols  []string
}

// generateReport returns the total size of all files stored; and
// The top N collections (by file size)
func generateReport(files []*file, n int) *report {
	if n > len(files) {
		n = len(files)
	}
	if n < 0 {
		n = 0
	}

	totalSize := 0
	colToSize := make(map[string]int, len(files))

	for _, f := range files {
		totalSize += f.size
		if f.collection != nil {
			colToSize[*f.collection] = colToSize[*f.collection] + f.size
		}
	}
	cols := make([]string, 0, len(files))
	for col := range colToSize {
		cols = append(cols, col)
	}

	// note: can change this from O(nlog(n)) to O(nlog(topN)) if we did some kind of insertion sort
	sort.Slice(cols, func(i, j int) bool {
		return colToSize[cols[i]] > colToSize[cols[j]]
	})
	return &report{
		totalSize: totalSize,
		topNCols:  cols[0:n],
	}
}
