package main

type IJuice interface {
	setName(name string)
	getName() string
	setPrice(price int)
	getPrice() int
}

type Juice struct {
	price int
	name  string
}

func (j *Juice) setName(name string) {
	j.name = name
}

func (j *Juice) getName() string {
	return j.name
}

func (j *Juice) setPrice(price int) {
	j.price = price
}

func (j *Juice) getPrice() int {
	return j.price
}
