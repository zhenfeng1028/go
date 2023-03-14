package main

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

type HouseOption func(*House)

func WithConcrete() HouseOption {
	return func(h *House) {
		h.Material = "concrete"
	}
}

func WithoutFireplace() HouseOption {
	return func(h *House) {
		h.HasFireplace = false
	}
}

func WithFloors(floors int) HouseOption {
	return func(h *House) {
		h.Floors = floors
	}
}

// NewHouse now takes a slice of option as the rest arguments
func NewHouse(opts ...HouseOption) *House {
	const (
		defaultFloors       = 2
		defaultHasFireplace = true
		defaultMaterial     = "wood"
	)

	h := &House{
		Material:     defaultMaterial,
		HasFireplace: defaultHasFireplace,
		Floors:       defaultFloors,
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the stantiated
		// *House as the argument
		opt(h)
	}

	// return the modified house instance
	return h
}

func main() {
	_ = NewHouse(
		WithConcrete(),
		WithoutFireplace(),
		WithFloors(3),
	)
}
