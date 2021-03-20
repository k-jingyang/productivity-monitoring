package main

import (
	"fmt"
	"os"

	"github.com/adlio/trello"
	"github.com/joho/godotenv"
)

const PERSONAL_BOARD_TODAY = "PERSONAL_BOARD_TODAY"
const PERSONAL_BOARD_TOMORROW = "PERSONAL_BOARD_TOMORROW"
const PERSONAL_BOARD_THIS_WEEK = "PERSONAL_BOARD_THIS_WEEK"
const PERSONAL_BOARD_WAITING = "PERSONAL_BOARD_WAITING"
const PERSONAL_BOARD_INBOX = "PERSONAL_BOARD_INBOX"
const PERSONAL_BOARD_DONE = "PERSONAL_BOARD_DONE"

const WORK_BOARD_TODAY = "WORK_BOARD_TODAY"
const WORK_BOARD_TOMORROW = "WORK_BOARD_TOMORROW"
const WORK_BOARD_THIS_WEEK = "WORK_BOARD_THIS_WEEK"
const WORK_BOARD_WAITING = "WORK_BOARD_WAITING"
const WORK_BOARD_INBOX = "WORK_BOARD_INBOX"
const WORK_BOARD_DONE = "WORK_BOARD_DONE"

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	token := os.Getenv("TOKEN")

	client := trello.NewClient(apiKey, token)

	listEnvNames := []string{
		PERSONAL_BOARD_TODAY,
		PERSONAL_BOARD_TOMORROW,
		PERSONAL_BOARD_THIS_WEEK,
		PERSONAL_BOARD_WAITING,
		PERSONAL_BOARD_THIS_WEEK,
		PERSONAL_BOARD_DONE,
		WORK_BOARD_TODAY,
		WORK_BOARD_TOMORROW,
		WORK_BOARD_THIS_WEEK,
		WORK_BOARD_WAITING,
		WORK_BOARD_THIS_WEEK,
		WORK_BOARD_DONE,
	}

	listEnvNames = []string{PERSONAL_BOARD_TODAY}

	for _, listEnvName := range listEnvNames {
		list, _ := client.GetList(os.Getenv(listEnvName), trello.Defaults())
		cards, _ := list.GetCards(trello.Defaults())

		fmt.Println(listEnvName, "has", len(cards), "cards.")
	}

}
