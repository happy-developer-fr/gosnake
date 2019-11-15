package snake

import (
	"fmt"
	"testing"
)

func TestCell(t *testing.T) {
	cell := Cell{true}

	if cell.Food != true {
		t.Errorf("Food should be true but is %v", cell.Food)
	}
	fmt.Println(cell.Food)
	// Output: toto
}

func BenchmarkCell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cell := Cell{true}
		fmt.Sprintf("cell is : %v", cell)
	}
}
