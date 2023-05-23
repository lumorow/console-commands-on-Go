package Flags

type Flags struct {
	A bool
}

func New(a bool) *Flags {
	return &Flags{
		A: a,
	}
}
