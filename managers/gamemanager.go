package managers

import (
	"bufio"
	"fmt"
	"os"
)

/*
This file contains every encounter in the game. Each encounter is labelled with a number, and each returns a number.
This number corresponds to a different, unique branch of the story, similar to pages in an adventure book.
Therefore, whichever int the function returns determines which encounter will come next.
i.e if Scene1 returns a 2, in main.go, Scene2 will be executed.
*/

//Scene1 is the first encounter in the game.
func Scene1(p *Player) int {
	fmt.Println("You awaken. The world around you is pitch black and cold, and you find yourself unable to move as you struggle against the void." +
		" Slowly but surely the darkness around you begins to fade, heat flowing over you as you began to hear the sound of bustling streets," +
		" the lively crowd around you passing by without a second thought of your presence. You knew you were somewhere, but you couldn't see a single thing." +
		" You thought to yourself, and you knew exactly what you had to do.",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "open eyes":
			a++
			selection = 2
		default:
			if attempts == 0 {
				fmt.Println("Just open your eyes.")
			} else if attempts == 1 {
				fmt.Println("Come on, really? Open your eyes.")
			} else if attempts == 2 {
				fmt.Println("Type 'open eyes' into the console.")
			} else if attempts == 3 {
				fmt.Println("You're kidding, right? If you don't type 'open eyes' this time, I'm actually going to explode.")
			} else {
				panic("The narrator's face became a deep red and he promptly exploded, killing you in the process. Maybe you should've just opened your eyes. Game over.")
			}
			attempts++
		}

	}
	return selection
}

//Scene2 is the second encounter in the game, branching off of Scene1.
func Scene2(p *Player) int {
	fmt.Println("You opened your eyes and were partially blinded by the sheer intensity of the world around you. As your eyes adjusted" +
		"your mouth hung open with wonder at the sight of the most vibrant city you had ever seen. Tall skyscarpers were covered in colorful billboards" +
		"illuminating the world around you to such a degree that you could hardly tell it was late into the night. The people around you were dressed" +
		"in an odd fashion, their style reminescent of some grungy, dystopian future you might find in a 1970s fiction. You looked down and noticed" +
		"you were dressed very similar to them, and to your surprise, your left arm was completely cybernetic. You shuddered to think what your hair might look like" +
		"as you began to walk about the street, looking around as you did your best to dodge the crowd. You spotted sseveral locations of interest but only a few" +
		"really caught your eye. Towering over the square was a skyscraper marked with a large '109', to the left of it you saw what appeared to be a nightclub," +
		"named in a language you couldn't understand, and to the right of it was a humble info kiosk. Each likely held answers as to where and when you were, but" +
		"the kiosk didn't seem too terribly interesting. You made up your mind and you travelled to the...",
	)

	return 3
}
