package main

import "fmt"

type Cataloger interface {
	Create()
}

type Catalog struct {
	ID   int
	Name string
}

func (ca *Catalog) Create() {

	fmt.Printf("send: %v", ca)
}

func Contro(c Cataloger) {
	c.Create()
}

func main() {
	var data = Catalog{
		ID: 1,
		Name: "hieu",
	}
	Contro(&data)
}
