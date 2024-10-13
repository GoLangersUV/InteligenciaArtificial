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
