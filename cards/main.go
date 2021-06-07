package main

func main() {
	// deck := newDeck()

	// hand, remainingDeck := deal(deck, 5)

	// hand.print()
	// remainingDeck.print()
	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")

}
