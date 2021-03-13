package main

func main()  {
	cards := newDeck()
	cards.shuffle().print()
	// fmt.Println(cards.deal(3))
	// fmt.Println(cards.toString())
	// cards.saveToFile("test_files")
	// deck, _ := newDeckFromFile("test_files");
	// deck.print()
}