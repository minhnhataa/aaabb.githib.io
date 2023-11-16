package models

import "fmt"

type Cataloger interface {
	Create()
}

type Catalog struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t Catalog) Create() {
	fmt.Printf("hi: %v", t.Name)
}

func main() {
	
}
