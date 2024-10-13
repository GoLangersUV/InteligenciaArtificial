package searchAlgorithms

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetMatrix() ([][]int, error) {
	// Open the file
	file, err := os.Open("./search/battery/Prueba1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a 2D slice to store the matrix
	var matrix [][]int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by spaces to get individual string numbers
		stringValues := strings.Fields(line)

		// Convert string numbers to integers and append to a row
		var row []int
		for _, stringValue := range stringValues {
			num, err := strconv.Atoi(stringValue)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return nil, err
			}
			row = append(row, num)
		}

		// Append the row to the matrix
		matrix = append(matrix, row)
	}

	return matrix, nil
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
