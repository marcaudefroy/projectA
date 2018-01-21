package main

import (
	"fmt"

	bar "github.com/marcaudefroy/projectB/models"
	beez "github.com/marcaudefroy/projectC/models"
)

func main() {
	_ = bar.Bar{}
	_ = beez.Beez{}

	fmt.Println("project A")
}
