package dummy

// Enum type of data generated
const (
	SALES = iota
)

// Record is the interface for all records that are generated
type Record interface {
	Preview()
}

// Generator is the interface for all generators
type Generator interface {
	Generate() Record
}
