package main

/*
	Here we will be learning how to make patterns
*/

func main() {
	// Comment and uncomment based on the pattern numbers

	//Inputs given to print the pattern - 1, 2
	//rows := 5
	//cols := 4

	//Inputs given to print the pattern - 3, 4, 5, 6, 7
	//input := 3

	/****************************************************
	Pattern - 1 - Rectangle pattern

	*****
	*****
	*****
	*****

	*/
	/*
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 2 - Hollow rectangle

	****
	*  *
	*  *
	*  *
	****

	*/
	/*
		for i := 1; i <= rows; i++ {
			for j := 1; j <= cols; j++ {
				if i == 1 || j == 1 || i == rows || j == cols {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 3 - Inverted half pyramid

	*****
	****
	***
	**
	*

	*/
	/*
		for i := input; i > 0; i-- {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 4 - Half pyramid with frontal spaces

	    *
	   **
	  ***
	 ****
	*****

	*/
	/*
		for i := 1; i <= input; i++ {
			for j := (input - i); j > 0; j-- {
				fmt.Print(" ")
			}
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 5 - Half pyramid using numbers

	1
	22
	333
	4444
	55555

	*/
	/*
		for i := 1; i <= input; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print(i)
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 6 - Floyd's triangle

	01
	02 03
	04 05 06
	07 08 09 10
	11 12 13 14 15

	*/
	/*
		count := 1
		for i := 1; i <= input; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%02d ", count)
				count++
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 7 - Butterfly

	*      *
	**    **
	***  ***
	********
	********
	***  ***
	**    **
	*      *

	*/
	/*
		cols := 2 * input
		for i := 1; i <= input; i++ {
			for j := 1; j <= cols; j++ {
				if j <= i || j > (cols-i) {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		for i := input; i > 0; i-- {
			for j := 1; j <= cols; j++ {
				if j <= i || j > (cols-i) {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	****************************************************/
}
