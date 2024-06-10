package main

import f "fmt"

func main() {
	f.Println("Practice using the fmt package!")

	const percent int = 50
	f.Printf("LOOADDING...%d%%\n", percent)

	f.Println("Hey, how are you?\n-Type good bad or okay")

	var response string
	f.Scan(&response)
	greet := f.Sprintf("im %v too!!", response)
	f.Println(greet)

}
