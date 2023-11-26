package dummy

type Record interface{}

type Generator interface {
	Generate() Record
}
