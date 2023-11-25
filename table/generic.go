package table

type Generic[A any] interface {
	Set(xPos, yPos uint16, value A) error
	ValueAt(xPos, yPos uint16) (A, error)
}
