package utils

import (
	"bufio"
	"fmt"
	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
	"os"
	"strconv"
	"strings"
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
	filePath := ("../battery/" + path) //"./search/battery/Prueba1.txt"
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
