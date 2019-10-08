package main

import (
	"fmt"
	"strconv"

	"github.com/natgwarner/project-0/managers"
)

func main() {
	var playername string
	//intitializes starting hp and money
	managers.Hp = 100
	managers.Money = 0
	//converts the hp and money vars to strings to be displayed by Player struct
	managers.Playerhp = strconv.Itoa(managers.Hp) + " hp"
	managers.Playermoney = strconv.Itoa(managers.Money) + " coins"

	fmt.Println("Ready for adventure? Good. What's your name?")
	fmt.Scanf("%s", &playername)

	var confirm string
	a := 0
	confirmattempts := 0
	for {
		if a > 0 {
			break
		}
		if confirmattempts == 0 {
			fmt.Println("I see. Your name is " + playername + ". Is that correct? (Please answer 'yes' or 'no'.)")
		} else if confirmattempts < 3 {
			fmt.Println("So, your name's " + playername + ", right?")
		} else {
			fmt.Println("Don't be difficult. Is your name " + playername + "?")
		}
		fmt.Scanf("%s", &confirm)
		switch confirm {
		case "yes":
			fmt.Println("Excellent. Let's begin.")
			a++
		case "no":
			fmt.Println("Okay, tell me again. What is your name?")
			fmt.Scanf("%s", &playername)
			confirmattempts++
		case "status":
			fmt.Println("You tried to look inside yourself, but you don't exist yet.")
			confirmattempts++
		default:
			fmt.Println("Just type in yes or no, please.")
			confirmattempts += 2
		}
	}

	var player = managers.Player{
		Name:      playername,
		Hitpoints: managers.Playerhp,
		Money:     managers.Playermoney,
	}

	fmt.Println("\nYou take a quick peek inside yourself (and your pockets).\n" + player.Name + "\n" + player.Hitpoints + "\n" + player.Money + "\n")

	fmt.Println("You clumsily tripped and fell on your face, receiving 20 damage.")
	managers.TakeDamage(20, &player)
	fmt.Println("You now have " + managers.Playerhp + ".")
	fmt.Println(player.Hitpoints)
}
