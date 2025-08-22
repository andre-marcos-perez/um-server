package sequence

import (
	"testing"
)

func TestLinkedListNew(t *testing.T) {
	list := NewLinkedList[int]()
	if got := list.Len(); got != 0 {
		t.Errorf("NewLinkedList[int]: expected %v, got %v", 0, list.Len())
	}
}

func TestLinkedListInsert(t *testing.T) {
	list := NewLinkedList[int]()

	list.Insert(1)
	expected := 1
	if got := list.Len(); got != expected {
		t.Errorf("NewLinkedList[int]: expected %v, got %v", expected, got)
	}

	list.Insert(2)
	expected = 2
	if got := list.Len(); got != expected {
		t.Errorf("NewLinkedList[int]: expected %v, got %v", expected, got)
	}

}

func TestLinkedListPeek(t *testing.T) {
	list := NewLinkedList[string]()

	if _, err := list.Peek(); err == nil {
		t.Errorf("NewLinkedList[string]: expected error, got nil")
	}

	list.Insert("1")
	expected := "1"
	if got, _ := list.Peek(); *got != expected {
		t.Errorf("NewLinkedList[string]: expected %v, got %v", expected, got)
	}

	list.Insert("2")
	expected = "2"
	if got, _ := list.Peek(); *got != expected {
		t.Errorf("NewLinkedList[string]: expected %v, got %v", expected, got)
	}

}

func TestLinkedListDelete(t *testing.T) {
	list := NewLinkedList[int]()
	list.Insert(1)
	list.Insert(2)

	expected := 2
	if got, _ := list.Delete(); *got != expected {
		t.Errorf("NewLinkedList[int]: expected %v, got %v", expected, got)
	}

	expected = 1
	if got, _ := list.Peek(); *got != expected {
		t.Errorf("NewLinkedList[int]: expected %v, got %v", expected, got)
	}

	expected = 1
	if got, _ := list.Delete(); *got != expected {
		t.Errorf("NewLinkedList[int]: expected %v, got %v", expected, got)
	}

	if _, err := list.Delete(); err == nil {
		t.Errorf("NewLinkedList[int]: expected error, got nil")
	}
}
