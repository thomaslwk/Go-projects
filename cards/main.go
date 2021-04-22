package main

func main() {
	// cards := newdeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
