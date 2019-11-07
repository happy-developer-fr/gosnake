package direction

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func Print(direction Direction) string {
	return [...]string{"North", "East", "South", "West"}[direction]
}
