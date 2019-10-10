package managers

import "fmt"

//Playername is the player's name.
var Playername string

//Hp is the base value of the player's hitpoints.
var Hp int

//Money is the int value of the player's wallet.
var Money int

//Playerhp is the string value of the player's hitpoints; used in the Player struct.
var Playerhp string

//Playermoney is the string value of the player's wallet; used in the Player struct.
var Playermoney string

//Player is the struct that holds all of the player's information.
type Player struct {
	Name      string
	Hitpoints string
	Money     string
}

//NewPlayer is a constructor that creates a player.
func NewPlayer(name string, hitpoints string, money string) Player {
	return Player{
		Name:      name,
		Hitpoints: hitpoints,
		Money:     money,
	}
}

//PlayerCheck allows the player to look at their current stats when they type the "check" command.
func PlayerCheck(p *Player) {
	fmt.Println("\nYou take a quick peek inside yourself (and your pockets).\n" + p.Name + "\n" + p.Hitpoints + "\n" + p.Money + "\n")
}
