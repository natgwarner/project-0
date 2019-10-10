package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/natgwarner/project-0/managers"
)

func init() {
	//intitializes starting hp and money
	managers.Hp = 100
	managers.Money = 0
	//converts the hp and money vars to strings to be displayed by Player struct
	managers.Playerhp = strconv.Itoa(managers.Hp) + " hp"
	managers.Playermoney = strconv.Itoa(managers.Money) + " credits"
	//allows the user to skip the intro dialogue by inputting a name after the run command, if they so choose
	if len(os.Args) > 1 {
		managers.Playername = os.Args[1]
	}
}

func main() {
	if managers.Playername == "" {
		fmt.Println("Ready for adventure? Good. What's your name?")
		fmt.Scanf("%s", &managers.Playername)

		var confirm string
		a := 0
		confirmattempts := 0
		for {
			if a > 0 {
				break
			}
			if confirmattempts == 0 {
				fmt.Println("I see. Your name is " + managers.Playername + ". Is that correct? (Please answer 'yes' or 'no'.)")
			} else if confirmattempts < 3 {
				fmt.Println("So, your name's " + managers.Playername + ", right?")
			} else {
				fmt.Println("Don't be difficult. Is your name " + managers.Playername + "?")
			}
			fmt.Scanf("%s", &confirm)
			switch confirm {
			case "yes":
				fmt.Println("Excellent. Let's begin.")
				a++
			case "no":
				fmt.Println("Okay, tell me again. What is your name?")
				fmt.Scanf("%s", &managers.Playername)
				confirmattempts++
			case "check":
				fmt.Println("You tried to look inside yourself, but you don't exist yet.")
				confirmattempts++
			default:
				fmt.Println("Just type in yes or no, please.")
				confirmattempts += 2
			}
		}
	}

	player := managers.NewPlayer(managers.Playername, managers.Playerhp, managers.Playermoney)
	/*
		managers.PlayerCheck(&player)

		fmt.Println("You clumsily trip and fall on your face, receiving 20 damage.")
		managers.TakeDamage(20, &player)
		fmt.Println("However, while you writhe against the concrete, you spot a card worth 2000 credits, lifting your spirits!")
		managers.GetMoney(2000, &player)

		managers.PlayerCheck(&player)
		managers.InventoryCheck()
		managers.InventoryAdd("Sword")
		managers.InventoryCheck()
	*/
	if managers.Scene1(&player) == 2 {
		managers.Scene2(&player)
	} else {
		panic("It's borked.")
	}
}
