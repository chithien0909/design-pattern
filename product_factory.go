package main

import (
	"fmt"
)

type IProduct interface {
	makeJuice() IJuice
	makeJam() IJam
}

func GetProductFactory(name string) (IProduct, error) {
	if name == "apple" {
		return &Apple{}, nil
	}
	if name == "orange" {
		return &Orange{}, nil
	}

	return nil, fmt.Errorf("product of type %s not found", name)
}
