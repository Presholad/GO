package main

import "fmt"

func main() {
	var publisher, writer, artist, title string = "DizzyBooks Publishing Inc.", "Tracey Hatchet", "Jewel Tampson", "Mr. GoToSleep"
	var year, pageNumber int64 = 1997, 14
	var grade float32 = 6.5
	var cost float64

	if grade > 8 {
		cost = 100
	} else {
		cost = 50
	}

	if pageNumber > 100 {
		cost += 65
	} else {
		cost += 30
	}

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "Year Made:",
		year, "Condition Grade", grade, "Number of Pages",
		pageNumber, "The Price of this Comic is:", cost, "\n")

	title = "Epic Vol.1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0

	if grade > 8 {
		cost = 100
	} else {
		cost = 50
	}

	if pageNumber > 100 {
		cost += 65
	} else {
		cost += 30
	}

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "Year Made:", year, "Condition Grade", grade, "Number of Pages", pageNumber, "The Price of this Comic is:", cost)

}
