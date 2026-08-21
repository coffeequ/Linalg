package linaltype

type Vector struct {
	Data [2]float64
}

func New(data [2]float64) *Vector {
	return &Vector{Data: data}
}
