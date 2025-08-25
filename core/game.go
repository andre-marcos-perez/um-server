package core

import "fmt"

const (
	GameInitPlayerDeckSize  int = 7
	GameInitPlayerMinAmount int = 2
	GameInitPlayerMaxAmount int = 10
	GameInitDrawDeckSize    int = 112
)

type Game struct {
	DrawDeck    *Deck
	DiscardDeck *Deck
}

func NewGame(players []Player) *Game {
	game := &Game{
		DrawDeck:    initDrawDeck(),
		DiscardDeck: initDiscardDeck(),
	}
	if len(players) < GameInitPlayerMinAmount || len(players) > GameInitPlayerMaxAmount {
		panic(fmt.Sprintf("game must between %v and %v players", GameInitPlayerMinAmount, GameInitPlayerMaxAmount))
	}
	for _, player := range players {
		for i := 0; i < GameInitPlayerDeckSize; i += 1 {
			card, err := game.DrawDeck.Draw()
			if err != nil {
				panic(err)
			}
			player.deck.Place(*card)
		}
	}
	return game
}

func initDrawDeck() *Deck {

	cards := make([]Card, GameInitDrawDeckSize)

	suits := []CardSuit{Red, Green, Blue, Yellow}
	ranks := []CardRank{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Reverse, Skip, PlusTwo}
	for _, suit := range suits {
		for _, rank := range ranks {
			for j := 0; j < 2; j += 1 {
				cards = append(cards, NewCard(suit, rank))
			}
		}
	}

	suits = []CardSuit{Wild}
	ranks = []CardRank{Any, PlusFour}
	for _, rank := range ranks {
		for j := 0; j < 4; j += 1 {
			cards = append(cards, NewCard(suits[0], rank))
		}
	}

	return NewDeck(
		WithCards(cards),
		WithShuffle(),
	)
}

func initDiscardDeck() *Deck {
	return NewDeck()
}
