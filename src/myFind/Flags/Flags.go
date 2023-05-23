package Flags

type Flags struct {
	Sl  bool
	D   bool
	F   bool
	Ext bool
}

func New(sl, d, f, ext bool) *Flags {
	return &Flags{
		Sl:  sl,
		D:   d,
		F:   f,
		Ext: ext,
	}
}
