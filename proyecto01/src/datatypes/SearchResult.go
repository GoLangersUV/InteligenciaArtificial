package datatypes

import "time"

type SearchResult struct {
	PathFound                []BoardCoordinate
	SolutionFound            bool
	ExpandenNodes, TreeDepth int
	Cost                     float32
	TimeExe                  time.Duration
}
