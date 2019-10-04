package main

import (
	"fmt"
	"strconv"
)

var playername string
var hp int
var money int

//Player is the struct that hold's all of the player's information.
type Player struct {
	Name      string
	Hitpoints string
	Money     string
}

func main() {
	//intitializes starting hp and money
	hp = 100
	money = 0
	//converts the hp and money vars to strings
	stringhp := strconv.Itoa(hp)
	stringmoney := strconv.Itoa(money)
	//allows the hp and money strings to be dispayed by the Player struct
	playerhp := stringhp + " hp"
	playermoney := stringmoney + " coins"

	fmt.Println("Hello. What is your name?")
	fmt.Scanf("%s", &playername)
	fmt.Println("I see. Your name is " + playername + ".")

	var player = Player{
		Name:      playername,
		Hitpoints: playerhp,
		Money:     playermoney,
	}

	fmt.Println("This is your information:\n" + player.Name + "\n" + player.Hitpoints + "\n" + player.Money + "\n")

	fmt.Println("You clumsily tripped and fell on your face, receiving 20 damage.")
	hp -= 20
	stringhp = strconv.Itoa(hp)
	playerhp = stringhp + " hp"
	fmt.Println("You now have " + playerhp + ".")
}
