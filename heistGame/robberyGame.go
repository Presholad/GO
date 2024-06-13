package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	isHeistOn := true
	eludedGuards := rand.Intn(100)
	rand.Seed(time.Now().UnixNano())

	//check if we eluded the guards
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		fmt.Println("YOU FAILED TO SNEAK PAST THE GUARDS, get a better disguise next time")
	}

	//check if vault is open
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("THE VAULT IS OPEN, GRAB THE CASH AND GO GO GO!!!!!")
	} else if !isHeistOn {
		fmt.Println("THE VAULT WAS TOO TOUGH TO CRACK, THE COPS ARE ON THE WAY CLEAR OUT.##########")
	}

	//check escape
	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("WE COULDNT ESCAPE WITH THE CASH HEIST FAILED!!!")

		case 1:
			isHeistOn = false
			fmt.Println("Turns out the vault doors dont open from the inside.....")

		case 2:
			isHeistOn = false
			fmt.Println("Cops caught up to us, Damn i guess we'll be spending christmas in jail")

		case 3:
			isHeistOn = false
			fmt.Println("The Cameras saw us leaving they know where we are, Gooduck it's everyman for himself now......")

		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("The total take from the heist is: %d\n", amtStolen)
	}

	//check heist status
	fmt.Println(isHeistOn)
}
