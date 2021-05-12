package main

func main() {
	cards := newDeck()

	// cards.print()
	// fmt.Println("Creating hand")

	// hand, _ := deal(cards, 4)
	// hand.print()

	cards.shuffle()
	cards.print()
	// remainingDeck.print()

	// fmt.Println(hand.toString())
	// hand.saveToFile("/Users/amanduggal/gits/GoPractice/helloworld/savedFile")

	// handFromFile := readFile("/Users/amanduggal/gits/GoPractice/helloworld/savedFile")
	// handFromFile.print()
}
