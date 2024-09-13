package search

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
