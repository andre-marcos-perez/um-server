package core

type Game struct {
	DrawPile    *Deck
	DiscardPile *Deck
}

func NewGame() *Game {
	return &Game{
		DrawPile:    initDrawPile(),
		DiscardPile: initDiscardPile(),
	}
}

func initDrawPile() *Deck {
	i := 0
	cards := make([]Card, 0, 112)
	suits := []CardSuit{Red, Green, Blue, Yellow}
	ranks := []CardRank{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Reverse, Skip, PlusTwo}
	for _, suit := range suits {
		for _, rank := range ranks {
			for j := 0; j < 2; j += 1 {
				cards[i] = *NewCard(suit, rank)
				i += 1
			}
		}
	}
	suits = []CardSuit{Wild}
	ranks = []CardRank{Any, PlusFour}
	for _, rank := range ranks {
		for j := 0; j < 4; j += 1 {
			cards[i] = *NewCard(suits[0], rank)
			i += 1
		}
	}
	return NewDeck(
		WithCards(cards),
		WithShuffle(),
	)
}

func initDiscardPile() *Deck {
	return NewDeck()
}
