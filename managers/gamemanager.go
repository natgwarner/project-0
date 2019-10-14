package managers

import (
	"bufio"
	"fmt"
	"io"
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
	io.WriteString(os.Stdout,
		`You awaken. The world around you is pitch black and cold, and you find yourself unable to move as you struggle against the void.
Slowly but surely the darkness around you begins to fade, heat flowing over you as you began to hear the sound of bustling streets, 
the lively crowd around you passing by without a second thought of your presence. You knew you were somewhere, but you couldn't see a single thing.
You thought to yourself, and you knew exactly what you had to do.`+"\n",
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
		case "check":
			PlayerCheck(p)
		case "backpack":
			InventoryCheck()
		case "search":
			fmt.Println("You tried to scan the area for anything useful, but the crowd made it impossible.")
		default:
			if attempts == 0 {
				fmt.Println("Just open your eyes.")
			} else if attempts == 1 {
				fmt.Println("Type 'open eyes' into the console.")
			} else if attempts == 2 {
				fmt.Println("Really?")
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
	io.WriteString(os.Stdout,
		`You opened your eyes and were partially blinded by the sheer intensity of the world around you. As your eyes adjusted, 
your mouth hung open with wonder at the sight of the most vibrant city you had ever seen. Tall skyscarpers were covered in colorful billboards, 
illuminating the world around you to such a degree that you could hardly tell it was late into the night. The people around you were dressed 
in an odd fashion, their style reminescent of some grungy, dystopian future you might find in a 1970s fiction. You looked down and noticed 
you were dressed very similar to them, and to your surprise, your left arm was completely cybernetic. You shuddered as you imagined what your hair might look like 
as you began to walk about the street, looking around as you did your best to dodge the crowd. You spotted several locations of interest but only a few 
really caught your eye: towering over the city hub was a skyscraper marked with a large '107', to the left of it you saw what appeared to be a nightclub, named in a 
language you couldn't understand, and to the right of it was a humble info kiosk. Each likely held answers as to where and when you were, but the kiosk didn't seem 
too terribly interesting. You made up your mind and you travelled to the...`+"\n",
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
		case "hint":
			fmt.Println("Your options are 'go to skyscraper', 'go to nightclub', and 'go to kiosk'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			InventoryCheck()
		case "search":
			if attempts < 1 {
				fmt.Println("Someone rudely bumps into you, scoffing at you as they walk away. You notice a small card fall from their bag and take it.")
				GetMoney(500, p)
				attempts++
			} else {
				fmt.Println("The crowd is too busy -- you can't find anything of value.")
			}
		case "go to skyscraper":
			selection = 3
			a++
		case "go to nightclub":
			selection = 4
			a++
		case "go to kiosk":
			selection = 5
			a++
		default:
			fmt.Println("You can't think of a reason why you would want to do that right now.")
		}
	}
	return selection
}

//Scene3 is the skyscraper area.
func Scene3(p *Player) int {
	io.WriteString(os.Stdout,
		`You approach the skyscraper and notice that there doesn't seem to be anyone entering or leaving. You find this curious, but that thought is quickly
pushed from your mind as you recoil in horror. To your dismay, the skyscraper has glass paneling, and your hair looks about ten times more awful than you could
have imagined. You sigh as you come to terms with your truly disgusting appearance before someone approaches you from the side. "Yo, you tryin' to get
in here?" You turned and saw a woman who's hair was somehow worse than your own, chewing gum and blowing a large bubble as it broke with a pop. She turned
to look into the building as she continued, "This building is owned by EdenTech, those bigwigs that control all the infrastructure that runs this city.
This is their HQ, and you ain't gettin' in unless you get someone inside to let you." She places her hand on the glass before clenching it into a fist, mimicking
a punch. "Can't break in neither, cause the drones'll getcha and chew ya up like a dog starved for weeks." You shot her a puzzled glance, and she seemed to
understand that you had no idea what she was talking about as she sighed. "Look, just stay away from here, 'less you wanna cause trouble. If you think you might,
go ahead and follow me. I know a place." She began to walk away and waved, and you considered following her. However, you could still visit the nightclub
or the info kiosk from here.`+"\n",
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
		case "hint":
			fmt.Println("Your options are 'follow', 'go to nightclub', and 'go to kiosk'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			InventoryCheck()
		case "search":
			if attempts < 1 {
				fmt.Println("You take a look around you and find a flyer for EdenTech on the ground, which you take.")
				InventoryAdd("EdenTech Flyer")
				attempts++
			} else {
				fmt.Println("You quickly scan the area for something interesting, but you didn't find anything.")
			}
		case "follow":
			selection = 6
			a++
		case "go to nightclub":
			selection = 4
			a++
		case "go to kiosk":
			selection = 5
			a++
		default:
			fmt.Println("You can't think of a reason why you would want to do that right now.")
		}
	}
	return selection
}

//Scene4 is the nightclub scene.
func Scene4(p *Player) int {
	io.WriteString(os.Stdout,
		`You decide to pay a visit to the nightclub, crossing the street and pushing your way through the crowd. You begin to hear the sound of a thumping 
bass as you approach, and oddly enough, you notice that there doesn't seem to be a queue -- or a door, for that matter. You head down the stairs and through the 
opening in the wall, and you're greeted by flashing lights and neon signs. The music was incredibly loud and drowned out almost everything else, a loud chatter 
filling the room as you once again pushed through another crowd. Cyborgs were dancing together on the floor, their metallic bodies covered in lights that flashed 
in time with the music as they raved, the sight almost blinding to you as you decided to keep to the edges of the club. As you're looking around, you accidentally bump 
into an incredibly large man, his attention immediately focused on you as he looked down at you, scowling. "Watch where you're goin', punk," he grunted at you 
as he slowly walked off, joined a group of thugs as they all shot you malicious glances. You figured you'd best stay away from them for now. You looked around the club 
for someone to talk to; maybe someone in here could give you some answers. You decided the bartender might know something, but you could also leave and go back to the info kiosk. 
You also supposed you could ignore caution and ask your new, threatening, friends. You decided to...`+"\n",
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
		case "hint":
			fmt.Println("Your options are 'talk to bartender', 'talk to thugs', and 'go to kiosk'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			InventoryCheck()
		case "search":
			if attempts < 1 {
				fmt.Println("You saw an unbroken glowstick on the ground and decided to pick it up. Maybe it would be useful later.")
				InventoryAdd("Glowstick")
				attempts++
			} else {
				fmt.Println("This area is a bit too crowded for thorough searching.")
			}
		case "talk to bartender":
			selection = 7
			a++
		case "talk to thugs":
			selection = 8
			a++
		case "go to kiosk":
			selection = 5
			a++
		default:
			fmt.Println("You thought about doing that, but the music was so loud it caused you to forget what it was.")
		}
	}
	return selection
}

//Scene5 is the kiosk scene.
func Scene5(p *Player) int {
	io.WriteString(os.Stdout,
		`You figured the safest thing to do would be to pay the info kiosk a visit. You took in the city a bit more as you approached the stand, 
still wondering how it was that you ended up in such a place. You arrived at the kiosk and noticed it was covered in stickers, mostly of cats, but some
seemed to be pictures of cartoonized food with cute faces. You tried to look for a pamphlet of some kind, but the only thing you could see was a service 
bell. You decided the only thing to do was ring it, and you gave it a quick smack. Almost immediately someone popped up from under the counter, and you had
never regretted a decision faster in your life. You were greeted by what appeared to be a short, crazy-eyed girl with pigtails, her work uniform a bit tattered. Her hands were 
behind her back for a moment, and you flinched away as she quickly thrust a large knife into the countertop, giggling at you as her wild eyes took you in, 
almost like she was trying to determine how hard you'd be to kill, her smile making you shiver. "I'm so happy to have someone to help! What can I do for you today? And don't you DARE leave 
before I've answered your questions!" You held up your hands and tried to stay back from the counter. You'd better get your questions out before this girl loses it, 
but you could always ignore her and try to run.`+"\n",
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
		case "hint":
			fmt.Println("Your options are 'where am i?', 'who are you?', and 'run away'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			InventoryCheck()
		case "search":
			fmt.Println("You'd better not make any sudden movements. You might catch that knife.")
		case "where am i?":
			selection = 9
			a++
		case "who are you?":
			if attempts < 1 {
				io.WriteString(os.Stdout, `"What a stupid, annoying question!" She let out a yell and pulled the knife up from the countertop, 
throwing it at you. You tried to dodge but the knife was faster, and it sank into your leg with a sickening shlick. You screamed in pain and pulled it out, 
noticing that you somehow weren't bleeding, but you were still in incredible pain. You take 50 damage, but keep the knife for yourself.`+"\n")
				InventoryAdd("Knife")
				TakeDamage(50, p)
				if Hp <= 0 {
					GameOver(5)
				}
				attempts++
			} else {
				fmt.Println("You must be crazier than her if you're going to ask that again.")
			}
		case "run away":
			fmt.Println("Bad move. You try to run, but the girl pulls out two more knives from under the counter and throws them your way, stabbing your vitals with deadly precision.")
			fmt.Println("You take 99999 damage.")
			TakeDamage(99999, p)
			if Hp <= 0 {
				GameOver(5)
			}
		default:
			fmt.Println("You decided against that. You wouldn't wanna incite this girl.")
		}
	}
	return selection
}
