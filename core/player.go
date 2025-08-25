package core

type Player struct {
	deck *Deck
}

func NewPlayer() *Player {
	return &Player{
		deck: NewDeck(),
	}
}
