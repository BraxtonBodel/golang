package findi

import "testing"

func TestAdd(t *testing.T) {
	find := New()
	find.Add(Item{})
	if len(find.Items) == 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	find := New()
	find.Add(Item{})
	results := find.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
