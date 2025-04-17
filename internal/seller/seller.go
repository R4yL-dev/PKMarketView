package seller

type Seller struct {
	name        string
	level       int
	nationality string
}

func NewSeller(name string, level int, nationality string) *Seller {
	return &Seller{
		name:        name,
		level:       level,
		nationality: nationality,
	}
}
