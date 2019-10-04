package managers

import "strconv"

//HPRefresh is able to modify the hp value, allowing the Player struct to display updated values.
func HPRefresh(i int) string {
	Playerhp = strconv.Itoa(i) + " hp"
	return Playerhp
}
