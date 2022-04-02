package DerivativesOfLinkedList

import (
	D "dataStructuresAndAlgorithms/dataStructure"
)

func AddToQueue() {
	D.CreateEnd()
}

func RemoveFromQueue() {
	D.DeleteBegin()
}

func PrintQueue() {
	D.PrintLinkedList()
}

/*
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter the choice: ")
	var choice int

	var continueWithDS = true
	fmt.Println("Hello! Here you can play with Queue as below:")

	for continueWithDS {
		fmt.Println("0 - Exit")
		fmt.Println("1 - Add data to Queue")
		fmt.Println("2 - Remove data from Queue")
		fmt.Println("99 - Print the data in the Queue")

		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())

		switch choice {
		case 0:
			fmt.Println("Exiting ...")
			continueWithDS = false
		case 1:
			D.AddToQueue()
		case 2:
			D.RemoveFromQueue()
		case 99:
			D.PrintQueue()
		}
	}
}
*/
