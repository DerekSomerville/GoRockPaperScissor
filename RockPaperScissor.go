package main

import (
	"fmt"
)
func determineWinner(player int, computer int) (string){
	var result string
	if player == computer {
		result = "Draw"
	} else if (player +1)%3 == computer {
		result = "Player wins"
	} else if (computer +1)%3 == player {
		result = "Computer wins"
	}
	return result
}

func print(message string){
	fmt.Println(message)
}

func getInputInt(message string)(string){
	var input string
	print(message)
	fmt.Scanln(&input)
	return input
}

func generateRequest(listOfItems []string) (string){
	var result string = "Please select"
	for counter := range listOfItems {
		result += fmt.Sprint(" ",counter," ",listOfItems[counter])
	}
	return result
}

func main() {
	print(determineWinner(0,1))
	var listOfItems = make([]string,3)
	listOfItems[0] = "apple"
	listOfItems[1] = "grape"
	var request = generateRequest(listOfItems)
	print(getInputInt(request))
}
