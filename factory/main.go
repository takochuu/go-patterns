package main

type Type int

const (
	TypeDisk = 1 << iota
)

type disk struct {
}

func newTypeDisk() *disk {
	return &disk{}
}

func newStore(t Type) *disk {
	switch t {
	case TypeDisk:
		return newTypeDisk()
	default:
		return newTypeDisk()
	}
}

func main() {
	newStore(TypeDisk)
}
