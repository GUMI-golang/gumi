package gumi

type Information struct {
	Dt         int64
}

func DefaultInformation() *Information {
	return &Information{
		Dt: 0,
	}
}
