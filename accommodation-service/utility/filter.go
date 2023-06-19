package utility

type Range struct {
	Min float64
	Max float64
}

type Filter struct {
	Range                *Range
	Benefits             []string
	Distinguished        bool
	DistinguishedHostIds []string
}
