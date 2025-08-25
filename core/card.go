package core

import "fmt"

type CardSuit uint8

const (
	Red CardSuit = iota
	Blue
	Green
	Yellow
	Wild
)

type CardRank uint8

const (
	Zero CardRank = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Reverse
	Skip
	Any
	PlusTwo
	PlusFour
)

type Card struct {
	suit CardSuit
	rank CardRank
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.suit, c.rank)
}

func NewCard(suit CardSuit, rank CardRank) Card {
	return Card{
		suit: suit,
		rank: rank,
	}
}
