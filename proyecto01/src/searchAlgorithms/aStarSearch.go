package searchAlgorithms

import (
  "time"
  "container/heap"
)

type AStarSearch struct {}

type Node struct {
  Position             Position
  Parent               *Node
  G                    float32 // Cost from start to current node
  H                    float32 // Heuristic cost to goal
  F                    float32 // Total cost (F = G + H)
  Depth                int
  HasPickedUpPassenger bool
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
  node := x.(*Node)
  *pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  node := old[n-1]
  *pq = old[0 : n-1]
  return node
}

func manhattanDistance(a, b Position) float32 {
  return float32(abs(a.X-b.X) + abs(a.Y-b.Y))
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

func heuristic(node *Node, env *Environment) float32 {
  minCost := float32(1)
  if node.HasPickedUpPassenger {
    return minCost * manhattanDistance(node.Position, env.GoalPosition)
  } else {
    toPassenger := manhattanDistance(node.Position, env.DogPosition)
    passengerToGoal := manhattanDistance(env.DogPosition, env.GoalPosition)
    return minCost * (toPassenger + passengerToGoal)
  }
}

func (a *AStarSearch) LookForGoal(env *Environment) SearchResult {
  startTime := time.Now()

  openList := &PriorityQueue{}
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
      totalPath := reconstructPathU(currentNode)
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

    successors := getSuccessors(currentNode, env)

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

type State struct {
  Position             Position
  HasPickedUpPassenger bool
}

func getSuccessors(node *Node, env *Environment) []*Node {
  var successors []*Node

  directions := []struct {
    dx, dy int
  }{
    {-1, 0}, // Up
    {0, 1},  // Right
    {1, 0},  // Down
    {0, -1}, // Left
  }

  for _, dir := range directions {
    newX := node.Position.X + dir.dx
    newY := node.Position.Y + dir.dy

    if newX >= 0 && newX < 10 && newY >= 0 && newY < 10 {
      cellValue := env.Matrix[newX][newY]
      if cellValue != WALL {
        cost := getCellCost(cellValue)

        hasPickedUpPassenger := node.HasPickedUpPassenger
        if !hasPickedUpPassenger && (newX == env.DogPosition.X && newY == env.DogPosition.Y) {
          hasPickedUpPassenger = true
        }

        newNode := &Node{
          Position:             Position{X: newX, Y: newY},
          Parent:               node,
          G:                    node.G + cost,
          Depth:                node.Depth + 1,
          HasPickedUpPassenger: hasPickedUpPassenger,
        }

        newNode.H = heuristic(newNode, env)
        newNode.F = newNode.G + newNode.H

        successors = append(successors, newNode)
      }
    }
  }

  return successors
}

func getCellCost(cellValue int) float32 {
  switch cellValue {
    case 0: 
    return 1
    case MIDCOST: 
    return 4
    case HEAVYCOST: 
    return 7
  default:
    return 1 
  }
}

func reconstructPathU(node *Node) []Position {
  var path []Position
  current := node
  for current != nil {
    path = append([]Position{current.Position}, path...)
    current = current.Parent
  }
  return path
}
