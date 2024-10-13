package searchAlgorithms

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"

  "github.com/Krud3/InteligenciaArtificial/src/datatypes"
)

func coordinateType(x int, y int, coordinateValue int, coordinateMap map[string]datatypes.BoardCoordinate) {
  if coordinateValue == 6 {
    coordinateMap["goal"] = datatypes.BoardCoordinate{X: x, Y: y}
  } else if coordinateValue == 5 {
    coordinateMap["passenger"] = datatypes.BoardCoordinate{X: x, Y: y}
  } else if coordinateValue == 2 {
    coordinateMap["init"] = datatypes.BoardCoordinate{X: x, Y: y}
  }
}

func GetMatrix(path string) (datatypes.ScannedMatrix, error) {
  // Open the file
  filePath := path //"./search/battery/Prueba1.txt"
  file, err := os.Open(filePath)
  if err != nil {
    fmt.Printf("Error opening file: %s; error: %s", filePath, err)
    var zero datatypes.ScannedMatrix
    return zero, err
  }
  defer file.Close()

  // Create a 2D slice to store the matrix
  var matrix datatypes.Matrix

  // Read the file line by line
  scanner := bufio.NewScanner(file)

  //For hold importants coordinates
  mainCoordinates := make(map[string]datatypes.BoardCoordinate)

  for lineIndex := 0; scanner.Scan(); lineIndex++ {
    line := scanner.Text()
    // Split the line by spaces to get individual string numbers
    stringValues := strings.Fields(line)

    // Convert string numbers to integers and append to a row
    var row []int
    for valueIndex, stringValue := range stringValues {
      num, err := strconv.Atoi(stringValue)
      if err != nil {
        fmt.Println("Error converting to integer:", err)
        var zero datatypes.ScannedMatrix
        return zero, err
      }
      coordinateType(lineIndex, valueIndex, num, mainCoordinates)
      row = append(row, num)
    }

    // Append the row to the matrix
    matrix = append(matrix, row)
  }
  return datatypes.ScannedMatrix{Matrix: matrix, MainCoordinates: mainCoordinates}, nil
}

func FromPosToPath(path []Position) [][]int {
  result := make([][]int, len(path))
  for i, pos := range path {
    result[i] = []int{pos.X, pos.Y}
  }
  return result
}

func ValidateMatrix(matrix [][]int) ([10][10]int, error) {
  var result [10][10]int

  if len(matrix) != 10 {
    return result, fmt.Errorf("La matriz debe tener 10 filas, pero tiene %d", len(matrix))
  }

  for i := 0; i < 10; i++ {
    if len(matrix[i]) != 10 {
      return result, fmt.Errorf("La fila %d debe tener 10 columnas, pero tiene %d", i, len(matrix[i]))
    }
    for j := 0; j < 10; j++ {
      result[i][j] = matrix[i][j]
    }
  }

  return result, nil
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

func reconstructPathI(node *Node) []Position {
  var path []Position
  current := node
  for current != nil {
    path = append([]Position{current.Position}, path...)
    current = current.Parent
  }
  return path
}

func getSuccessors(node *Node, env *Environment, useGCost bool) []*Node {
  var successors []*Node

  directions := []struct {
    dx, dy int
  }{
    {-1, 0}, // Arriba
    {0, 1},  // Derecha
    {1, 0},  // Abajo
    {0, -1}, // Izquierda
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
          Depth:                node.Depth + 1,
          HasPickedUpPassenger: hasPickedUpPassenger,
        }

        newNode.H = heuristic(newNode, env)

        if useGCost {
          newNode.G = node.G + cost
          newNode.F = newNode.G + newNode.H
        }

        successors = append(successors, newNode)
      }
    }
  }

  return successors
}

func getCellCost(cellValue int) float32 {
    switch cellValue {
    case 0: // Tráfico liviano
        return 1
    case MIDCOST: // Tráfico medio
        return 4
    case HEAVYCOST: // Tráfico pesado
        return 7
    default:
        return 1 // Default cost for other types
    }
}
