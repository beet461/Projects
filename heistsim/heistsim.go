package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("You have passed the guards, now you may commence the heist.")
	} else {
		isHeistOn = false
		fmt.Println("You disappointment, you couldn't even make it past the first step!")
	}
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Stop wasting time you moron and get the money and get out of there!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Stop trying to open it and get out of there or it's jail for you!")
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("You have ruined it, get out.")
		case 1:
			isHeistOn = false
			fmt.Println("You have ruined it, get out.")
		case 2:
			isHeistOn = false
			fmt.Println("You have ruined it, get out.")
		case 3:
			isHeistOn = false
			fmt.Println("You have ruined it, get out.")
		default:
			isHeistOn = true
			fmt.Println("YOU HAVE DONE IT!")
		}
	}

	amtStolen := 10000 + rand.Intn(1000000)
	if isHeistOn {
		fmt.Println("You have stolen a grand total of:", amtStolen)
	} else {
		fmt.Println("So close noob.")
	}

	if amtStolen >= 50000 {
		fmt.Println("You get to keep 20000")
	} else {
		fmt.Println("I'm calling the cops. That's not enough money")
	}

	fmt.Println(isHeistOn)
}
