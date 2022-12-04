package main

func main() {
	cards := newDeck()
	//hand, reamingCards := deals(cards, 3)
	//hand.printCards()
	//reamingCards.printCards()
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.printCards()
}
