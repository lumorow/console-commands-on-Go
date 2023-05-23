package Flags

type Flags struct {
	L bool
	M bool
	W bool
}

func New(l, m, w, ext bool) *Flags {
	return &Flags{
		L: l,
		M: m,
		W: w,
	}
}
