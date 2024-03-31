package main

type Apple struct {
}

func (a *Apple) makeJuice() IJuice {
	return &AppleJuice{
		Juice{
			price: 100,
			name:  "Apple Juice",
		},
	}
}

func (a *Apple) makeJam() IJam {
	return &AppleJam{
		Jam{
			price: 200,
			name:  "Apple Jam",
		},
	}
}
