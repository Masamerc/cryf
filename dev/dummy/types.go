package dummy

type Record interface {
	Preview()
}

type Generator interface {
	Generate() Record
}
