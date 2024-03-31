package main

type Orange struct {
}

func (o *Orange) makeJuice() IJuice {
	return &OrangeJuice{
		Juice: Juice{
			name:  "Orange Juice",
			price: 50,
		},
	}
}

func (o *Orange) makeJam() IJam {
	return &Jam{
		name:  "Orange Jam",
		price: 100,
	}
}
