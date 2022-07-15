package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number of rows in matrix: ")
	scanner.Scan()
	rows, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Please enter the number of columns in matrix: ")
	scanner.Scan()
	cols, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Please enter the matrix: ")
	matrix := make([][]int64, 0)
	for i := 0; i < rows; i++ {
		row := make([]int64, 0)
		for j := 0; j < cols; j++ {
			input, _ := reader.ReadString(' ')
			inputInt, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
			row = append(row, inputInt)
		}
		matrix = append(matrix, row)
	}

	fmt.Println(matrix)
	fmt.Println("Please enter the search value:")
	scanner.Scan()
	inpu, _ := strconv.Atoi(scanner.Text())
	for outerIndex, ro := range matrix {
		for innerIndex, value := range ro {
			if value == int64(inpu) {
				fmt.Printf("row: %d, column:%d", outerIndex, innerIndex)
				break
			}
		}
	}
}
