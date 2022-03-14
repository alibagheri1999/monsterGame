package interaction

import (
	"fmt"
	"os"
)

type RoundDate struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	fmt.Println("GAME OVER")
	fmt.Println("-------------------------")
	fmt.Printf("%v won !\n", winner)
}

func PrintRoundStatistics(roundData *RoundDate) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("player attacked monster for %v damge.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("playerperformed a strong attack aginst monster for %v damge.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("moster attackes player for  %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("player health : %v.\n", roundData.PlayerHealth)
	fmt.Printf("player health : %v.\n", roundData.MonsterHealth)

}

func WriteLogFile(roundData *[]RoundDate) {
	file, err := os.Create("gameLog.txt")

	if err != nil {
		fmt.Println("saving file failed")
		return
	}
	for index, value := range *roundData {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Heal value":     fmt.Sprint(value.PlayerHealValue),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)

		if err != nil {
			fmt.Println("writting file failed")
			continue
		}
	}

	file.Close()

	fmt.Println("Wrote data to log!!!")
}
