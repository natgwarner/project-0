package managers

//Hp is the base value of the player's hitpoints; used for initialization.
var Hp int

//Money is the int value of the player's wallet; used for initialization.
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
