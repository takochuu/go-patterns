package creational

type Type int

const (
	TypeDisk = 1 << iota
	TypeMem
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

func NewStore(t Type) Store {
	switch t {
	case TypeDisk:
		return newTypeDisk()
	case TypeMem:
		return newTypeMem()
	}
}

func main() {
	s, _ := data.NewStore(Typemem)
	f, _ := s.Open("file")

	n, _ := f.Write([]byte("data"))
	defer f.Close()
}
