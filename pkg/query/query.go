package query

type frog struct {
	id   int
	name string
}

func (f frog) withID(id int) frog {
	f.id = id
	return f
}

func (f frog) withName(name string) frog {
	f.name = name
	return f
}

func newFrog() frog {
	return frog{}
}
