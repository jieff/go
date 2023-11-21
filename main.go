package main

import "github.com/jieff/go/internal/entity"

type Car struct {
	Model string
	Color string
}

func main() {

	order, err := entity.NewOrder("1", -10, 1)
	if err != nil {
		println(err.Error())
	} else {
		println(order.Id)
	}
}
