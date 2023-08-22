package main

import (
	"Programa4/pkg"
	"fmt"
)

func main() {
	// CI CD
	pkg.ListFiles("./")
	fmt.Println()
	pkg.ListDirectories("./")
}
