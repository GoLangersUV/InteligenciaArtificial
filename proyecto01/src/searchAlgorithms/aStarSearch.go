package searchAlgorithms

import (
  "container/heap"
  "time"
)

type AStarSearch struct{}

func (a *AStarSearch) LookForGoal(env *Environment) SearchResult {
  startTime := time.Now()

  compare := func(a, b *Node) bool {
    return a.F < b.F
  }

  openList := &PriorityQueue{compare: compare}
  heap.Init(openList)
  closedList := make(map[State]bool)

  startNode := &Node{
    Position:             env.InitPosition,
    G:                    0,
    H:                    heuristic(&Node{Position: env.InitPosition}, env),
    F:                    0 + heuristic(&Node{Position: env.InitPosition}, env),
    Depth:                0,
    HasPickedUpPassenger: false,
  }

  heap.Push(openList, startNode)

  var expandedNodes int
  var maxDepth int

  for openList.Len() > 0 {
    currentNode := heap.Pop(openList).(*Node)
    expandedNodes++
    if currentNode.Depth > maxDepth {
      maxDepth = currentNode.Depth
    }

    if currentNode.Position == env.GoalPosition && currentNode.HasPickedUpPassenger {
      totalPath := reconstructPathI(currentNode)
      totalCost := currentNode.G
      timeExecuted := time.Since(startTime)

      return SearchResult{
        SolutionFound: true,
        ExpandedNodes: expandedNodes,
        TreeDepth:     maxDepth,
        Cost:          totalCost,
        TimeExecuted:  timeExecuted,
        Path:          totalPath,
      }
    }

    state := State{
      Position:             currentNode.Position,
      HasPickedUpPassenger: currentNode.HasPickedUpPassenger,
    }
    closedList[state] = true

    successors := getSuccessors(currentNode, env, true)

    for _, successor := range successors {
      state := State{
        Position:             successor.Position,
        HasPickedUpPassenger: successor.HasPickedUpPassenger,
      }
      if closedList[state] {
        continue
      }

      heap.Push(openList, successor)
    }
  }

  timeExecuted := time.Since(startTime)
  return SearchResult{
    SolutionFound: false,
    ExpandedNodes: expandedNodes,
    TreeDepth:     maxDepth,
    TimeExecuted:  timeExecuted,
  }
}

