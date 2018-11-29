package main

import (
	"fmt"
	"log"
	"os"

	"github.com/StabbyCutyou/gf"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("You must provide the fully qualifed path to your package")
	}
	f := gf.GoFiler{Path: os.Args[1]}
	s, _ := f.BestPractices()
	fmt.Println(s)
}
