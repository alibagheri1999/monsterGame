package interaction

import (
	"bufio"
	"fmt"
	"os"
	// "reflect"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if strings.TrimSpace(playerChoice) == "1" {
			return "ATTACK"
		} else if strings.TrimSpace(playerChoice) == "2" {
			return "HEAL"
		} else if strings.TrimSpace(playerChoice) == "3" && specialAttackIsAvailable {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("Fetching the user input failed. please try again!!!")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput, nil
}
