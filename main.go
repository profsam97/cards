package main

func main() {
	// var card string = "Ace of shades" this is the long form of creating and assigning variables
	// card := newCard()
	// card := "Ace of shades" //golang does inference for use by infering the type

	// card = "Five set of diamonds" // here we are reassigning we should ignore the semicolon :

	// this is a slice, it just like an array. but it allows you to add or remove items, all items in an array must
	// be of the same type
	// cards := newDeckFromFile("my_cards")
	cards := newDeck()
	// hand, remainingHand := deal(cards, 5)

	// hand.print()
	// remainingHand.print()
	// print is a method on type deck
	// cards.print()

	// fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
}
