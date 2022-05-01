package main

/*
	Here we will be learning how to make patterns
*/

func main() {
	// Comment and uncomment based on the pattern numbers

	//Inputs given to print the pattern - 1, 2
	//rows := 5
	//cols := 4

	//Inputs given to print the pattern - 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
	//input := 9

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

	/****************************************************
	Pattern - 8 - Inverted Numbers Pattern

	12345
	1234
	123
	12
	1

	*/
	/*
		for i := input; i > 0; i-- {
			for j := 1; j <= i; j++ {
				fmt.Print(j)
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 9 - 1_0 Pattern

	1
	01
	101
	0101
	10101

	*/
	/*
		for i := 1; i <= input; i++ {
			for j := 1; j <= i; j++ {
				if (i+j)%2 == 0 {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 10 - Rhombus

	    *****
	   *****
	  *****
	 *****
	*****

	*/
	/*
		for i := 1; i <= input; i++ {
			for j := 1; j <= (input - i); j++ {
				fmt.Print(" ")
			}
			for j := 1; j <= input; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 11 - triangle with numbers

	    1
	   1 2
	  1 2 3
	 1 2 3 4
	1 2 3 4 5

	*/
	/*
		for i := 1; i <= input; i++ {
			for j := 1; j <= (input - i); j++ {
				fmt.Print(" ")
			}
			for j := 1; j <= i; j++ {
				fmt.Printf("%d ", j)
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 12 - Palindromic pattern

	    1
	   212
	  32123
	 4321234
	543212345

	*/
	/*
		for i := 1; i <= input; i++ {
			for j := 1; j <= (input - i); j++ {
				fmt.Print(" ")
			}
			for j := i; j >= 1; j-- {
				fmt.Print(j)
			}
			for j := 2; j <= i; j++ {
				fmt.Print(j)
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 13 - Mirrored Triangle

	    *
	   ***
	  *****
	 *******
	*********
	*********
	 *******
	  *****
	   ***
	    *

	*/
	/*
		for i := 1; i <= input; i++ {
			for j := 1; j <= (input - i); j++ {
				fmt.Print(" ")
			}
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			for j := 1; j <= (i - 1); j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
		for i := input; i >= 1; i-- {
			for j := 1; j <= (input - i); j++ {
				fmt.Print(" ")
			}
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			for j := 1; j <= (i - 1); j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	****************************************************/

	/****************************************************
	Pattern - 14 - Zig-Zag

	input is no. of stars
	here no. of rows are fixed, i.e. 3

	  *   *
	 * * * *
	*   *   *

	*/
	/*
		for i := 1; i <= 3; i++ {
			for j := 1; j <= input; j++ {
				if ((i+j)%4 == 0) || (i == 2 && j%4 == 0) {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	****************************************************/

}
