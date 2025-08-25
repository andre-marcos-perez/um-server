package sequence

import (
	"reflect"
	"testing"
)

func TestLinkedListNew(t *testing.T) {
	t.Run("it should create a new linked list", func(t *testing.T) {
		list := NewLinkedList[int]()
		if got := list.Len(); got != 0 {
			t.Errorf("NewLinkedList[int]: expected %v, got %v", 0, list.Len())
		}
	})
}

func TestLinkedListIter(t *testing.T) {
	t.Run("it should iterate over all elements", func(t *testing.T) {
		list := NewLinkedList[int]()
		list.Insert(1)
		list.Insert(2)
		list.Insert(3)
		elements := list.Iter()
		for i := 0; i < 3; i++ {
			expected := 3 - i
			if got := elements[i]; got != expected {
				t.Errorf("NewLinkedList[int]: expected %v, got %v", expected, got)
			}
		}
	})
}

func TestLinkedListPeek(t *testing.T) {
	t.Run("it should always return the element at the head", func(t *testing.T) {
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
	})
}

func TestLinkedListSort(t *testing.T) {
	t.Run("it should randomly sort elements", func(t *testing.T) {
		list := NewLinkedList[int]()
		for i := 0; i < 100; i++ {
			list.Insert(i)
		}
		original := list.Iter()
		list.Sort()
		got := list.Iter()
		// very likely different order
		if reflect.DeepEqual(got, original) {
			t.Errorf("NewLinkedList[int]: expected %v, got %v", list, got)
		}
	})
}

func TestLinkedListInsert(t *testing.T) {
	t.Run("it should insert elements to the head", func(t *testing.T) {
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
	})
}

func TestLinkedListDelete(t *testing.T) {
	t.Run("it should delete elements from the head", func(t *testing.T) {
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
	})
}
