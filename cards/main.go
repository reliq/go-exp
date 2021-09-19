package main

func main() {
	// cards := newDeck()

	// save to file:
	// err := cards.saveToFile("my_deck.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// retrieve from file:
	// cards := newDeckFromFile("my_deck.txt")
	// cards.print()

	// deal:
	// fmt.Println("Dealing cards:")
	// hand, _ := deal(cards, 5)

	// fmt.Println("hand:", hand.toString())
	// fmt.Println("remaining deck:", remainingCards.print())

	// shuffle:
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
