package engine

import (
	"errors"
	"github.com/andre-marcos-perez/um-server/core"
	"github.com/andre-marcos-perez/um-server/player"
)

const (
	GameInitPlayerDeckSize  int = 7
	GameInitPlayerMinAmount int = 2
	GameInitPlayerMaxAmount int = 10
	GameInitDrawDeckSize    int = 112
)

var (
	ErrGamePlayerInvalidAmount = errors.New("invalid amount of players")
)

type Game struct {
	DrawDeck    *core.Deck
	DiscardDeck *core.Deck
}

func NewGame(players []player.Player) (*Game, error) {

	if len(players) < GameInitPlayerMinAmount || len(players) > GameInitPlayerMaxAmount {
		return nil, ErrGamePlayerInvalidAmount
	}

	game := &Game{
		DrawDeck:    initDrawDeck(),
		DiscardDeck: initDiscardDeck(),
	}

	for _, player := range players {
		for i := 0; i < GameInitPlayerDeckSize; i += 1 {
			card, err := game.DrawDeck.Draw()
			if err != nil {
				panic(err)
			}
			player.Deck.Place(*card)
		}
	}

	return game, nil
}

func initDrawDeck() *core.Deck {

	cards := make([]core.Card, 0)

	suits := []core.CardSuit{core.Red, core.Green, core.Blue, core.Yellow}
	ranks := []core.CardRank{core.Zero, core.One, core.Two, core.Three, core.Four, core.Five, core.Six, core.Seven, core.Eight, core.Nine, core.Reverse, core.Skip, core.PlusTwo}
	for _, suit := range suits {
		for _, rank := range ranks {
			for j := 0; j < 2; j += 1 {
				cards = append(cards, core.NewCard(suit, rank))
			}
		}
	}

	suits = []core.CardSuit{core.Wild}
	ranks = []core.CardRank{core.Any, core.PlusFour}
	for _, rank := range ranks {
		for j := 0; j < 4; j += 1 {
			cards = append(cards, core.NewCard(suits[0], rank))
		}
	}

	return core.NewDeck(
		core.WithCards(cards),
		core.WithShuffle(),
	)
}

func initDiscardDeck() *core.Deck {
	return core.NewDeck()
}
