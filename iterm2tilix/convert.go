// convert is an example cli tool to try color4tilix library used to convert
// .itermcolors files in tilix json schema files.
// In this example we use the itermcolors_parser, but users can implement their own parser.
// This can be used as example for other conversions.
package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/tux-eithel/color4tilix/lib"
)

func main() {

	dir := flag.String("d", "", "Directory to convert")
	flag.Parse()

	if dir == nil || *dir == "" {
		log.Fatalf("no directory given: use -h")
	}

	files, err := filepath.Glob(*dir + "/*.itermcolors")
	if err != nil {
		log.Fatalf("unable to reads files: %v", err)
	}
	fmt.Printf("Read %d files...\n", len(files))

	ee := make(chan error)

	go func() {
		for e := range ee {
			log.Printf("%v\n", e)
		}
	}()

	fnc := func(n string, data []byte) (color4tilix.Paletter, error) {
		return color4tilix.NewItermColorsTransformer(n, data)
	}

	c := color4tilix.CreateBatch(*dir, files, fnc, ee)
	fmt.Printf("Converted %d files\n", c)
	close(ee)

}
