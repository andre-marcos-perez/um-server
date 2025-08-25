package core

import "fmt"

type cardSuit uint8

const (
	Red cardSuit = iota
	Blue
	Green
	Yellow
)

type cardRank uint8

const (
	Zero cardRank = iota
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
	Wild
	PlusTwo
	PlusFour
)

type Card struct {
	suit cardSuit
	rank cardRank
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.suit, c.rank)
}

func NewCard(suit cardSuit, rank cardRank) *Card {
	return &Card{
		suit: suit,
		rank: rank,
	}
}
