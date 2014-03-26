package save

type Saver interface {
	Save(v string) error
}
