package reader

type IReader interface {
	Read() ([]string, error)
	CountLines() error
	CountWords() error
	CountBytes() error
}
