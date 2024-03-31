package main

type IJam interface {
	getPrice() int
	getName() string
	setPrice(int)
	setName(string)
}

type Jam struct {
	price int
	name  string
}

func (j *Jam) setName(name string) {
	j.name = name
}

func (j *Jam) getName() string {
	return j.name
}

func (j *Jam) setPrice(price int) {
	j.price = price
}

func (j *Jam) getPrice() int {
	return j.price
}
