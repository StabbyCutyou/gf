package main

import (
	"flag"
	"log"

	"github.com/StabbyCutyou/gf"
)

func main() {
	// Start parsing all of our flags
	_ = flag.String("a", "", "reserved for future use")
	_ = flag.String("b", "", "reserved for future use")
	_ = flag.String("c", "", "reserved for future use")
	dryRun := flag.Bool("d", false, "Dry run mode - create a Gofile without deleting other files (note, this will leave the package unable to compile)")
	_ = flag.String("e", "", "reserved for future use")
	_ = flag.String("f", "", "reserved for future use")
	_ = flag.String("g", "", "reserved for future use")
	_ = flag.String("h", "", "reserved for future use")
	_ = flag.String("i", "", "reserved for future use")
	_ = flag.String("j", "", "reserved for future use")
	_ = flag.String("k", "", "reserved for future use")
	_ = flag.String("l", "", "reserved for future use")
	_ = flag.String("m", "", "reserved for future use")
	_ = flag.String("n", "", "reserved for future use")
	organize := flag.Bool("o", false, "Organize package level fixtures at the top of the resulting Gofile")
	pathkage := flag.String("p", "", "The fully qualified path of the package to Gofilize")
	_ = flag.String("q", "", "reserved for future use")
	_ = flag.String("r", "", "reserved for future use")
	_ = flag.String("s", "", "reserved for future use")
	_ = flag.String("t", "", "reserved for future use")
	_ = flag.String("u", "", "reserved for future use")
	_ = flag.String("v", "", "reserved for future use")
	_ = flag.String("w", "", "reserved for future use")
	_ = flag.String("x", "", "reserved for future use")
	_ = flag.String("y", "", "reserved for future use")
	_ = flag.String("z", "", "reserved for future use")

	flag.Parse()
	if *pathkage == "" {
		log.Fatal("You must provide a package with the -p flag")
	}
	f := gf.GoFiler{
		Pathkage: *pathkage,
		DryRun:   *dryRun,
		Organize: *organize,
	}
	if err := f.BestPractices(); err != nil {
		log.Fatal(err)
	}
}
