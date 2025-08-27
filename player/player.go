package player

import "github.com/andre-marcos-perez/um-server/core"

type Player struct {
	Deck *core.Deck
}

func NewPlayer() *Player {
	return &Player{
		Deck: core.NewDeck(),
	}
}
