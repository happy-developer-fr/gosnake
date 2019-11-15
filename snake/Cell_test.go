package snake

import "testing"

func TestCell(t *testing.T) {
	cell := Cell{true}

	if cell.Food != true {
		t.Errorf("Food should be true but is %v", false)
	}
}
