package main

import (
	"github.com/alibagheri1999/monsterGame/action"
	"github.com/alibagheri1999/monsterGame/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundDate{}

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""
	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialRound)

	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerHealth int
	var monsterHealth int
	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = action.AttackMoster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = action.HealPlayer()
	} else {
		playerAttackDmg = action.AttackMoster(true)
	}

	monsterAttackDmg = action.AttackPlayer()

	playerHealth, monsterHealth = action.GetHealthAmounts()

	roundData := interaction.RoundDate{
		Action:           userChoice,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}

	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}
	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
