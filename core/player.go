package core

type Player struct {
	Deck *Deck
}

func NewPlayer() *Player {
	return &Player{
		Deck: NewDeck(),
	}
}
