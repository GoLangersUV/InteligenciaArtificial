package searchAlgorithms

import (
	"math"
	"time"

	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
)

func (a *DepthSearch) LookForGoal(e *enviroment) datatypes.SearchResult {
	var stack []datatypes.AgentStep
	var visited = make(map[datatypes.BoardCoordinate]bool)
	var parentNodes []datatypes.AgentStep
	expandedNodes := 0
	totalCost := float32(0)

	initialPosition := e.agent.position
	stack = append(stack, initialPosition)

	start := time.Now()

	pathToPassenger := []datatypes.BoardCoordinate{}
	pathToGoal := []datatypes.BoardCoordinate{}

	//fmt.Println("Iniciando búsqueda en profundidad...")
	//fmt.Printf("Posición inicial del agente: [%d, %d]\n", initialPosition.CurrentPosition.X, initialPosition.CurrentPosition.Y)

	for len(stack) > 0 {
		//fmt.Println("\nEstado actual de la pila:", stack)

		currentStep := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//fmt.Printf("Nodo desapilado: [%d, %d]\n", currentStep.CurrentPosition.X, currentStep.CurrentPosition.Y)

		if visited[currentStep.CurrentPosition] {
			//fmt.Printf("Nodo [%d, %d] ya visitado, continuando...\n", currentStep.CurrentPosition.X, currentStep.CurrentPosition.Y)
			continue
		}
		visited[currentStep.CurrentPosition] = true
		parentNodes = append(parentNodes, currentStep)

		// Calculate the cost of the current cell

		cellValue := e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y]
		cellCost := getCellCost(cellValue)
		totalCost += cellCost
		//fmt.Printf("Pasando por la casilla [%d, %d] con valor %d, costo acumulado: %.2f\n", currentStep.CurrentPosition.X, currentStep.CurrentPosition.Y, cellValue, totalCost)

		// Phase 1: Find the passenger
		if !e.agent.passenger && e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == 5 {
			//fmt.Println("Pasajero encontrado en la posición:", currentStep.CurrentPosition)

			e.passengerPosition = datatypes.BoardCoordinate{
				X: currentStep.CurrentPosition.X,
				Y: currentStep.CurrentPosition.Y,
			}
			e.agent.passenger = true
			pathToPassenger = reconstructPath(parentNodes, currentStep)

			// reload stack after finding the passenger
			stack = nil
			parentNodes = nil
			visited = make(map[datatypes.BoardCoordinate]bool)
			stack = append(stack, datatypes.AgentStep{
				PreviousPosition: datatypes.BoardCoordinate{
					X: math.MaxInt,
					Y: math.MaxInt,
				},
				CurrentPosition: currentStep.CurrentPosition,
				Depth:           currentStep.Depth + 1,
			})
			//fmt.Println("Reiniciando la búsqueda desde el pasajero.")
			continue
		}

		// Phase 2: Find the goal
		if e.agent.passenger && e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == 6 {
			//fmt.Println("Destino encontrado en la posición:", currentStep.CurrentPosition)
			end := time.Now()
			pathToGoal = reconstructPath(parentNodes, currentStep)
			combinedPath := append(pathToPassenger, pathToGoal...)

			//fmt.Println("Camino combinado desde el pasajero hasta el destino:", combinedPath)
			//fmt.Printf("Tiempo de ejecución: %v\n", end.Sub(start))

			return datatypes.SearchResult{
				PathFound:     combinedPath,
				SolutionFound: true,
				ExpandenNodes: expandedNodes,
				TreeDepth:     len(combinedPath),
				Cost:          totalCost - 1,
				TimeExe:       end.Sub(start),
			}
		}

		// Expand the neighbors of the current node
		e.agent.position = currentStep
		agentPerception := Percept(e.agent, e.board)
		expandedNodes++

		//fmt.Printf("Expandiendo vecinos del nodo [%d, %d]: %v\n", currentStep.CurrentPosition.X, currentStep.CurrentPosition.Y, agentPerception)

		for _, perception := range agentPerception {
			if !visited[perception.Coordinate] && perception.Coordinate != currentStep.PreviousPosition {
				//fmt.Printf("Apilando nodo: [%d, %d]\n", perception.Coordinate.X, perception.Coordinate.Y)
				stack = append(stack, datatypes.AgentStep{
					Action:           perception.Direction,
					Depth:            currentStep.Depth + 1,
					CurrentPosition:  perception.Coordinate,
					PreviousPosition: currentStep.CurrentPosition,
				})
			} else {
				//fmt.Printf("Vecino [%d, %d] ya visitado o es el nodo previo, no apilado.\n", perception.Coordinate.X, perception.Coordinate.Y)
			}
		}
	}

	//fmt.Println("No se encontró una solución.")
	return datatypes.SearchResult{
		SolutionFound: false,
		ExpandenNodes: expandedNodes,
		TreeDepth:     0,
		Cost:          totalCost,
		TimeExe:       time.Since(start),
	}
}
