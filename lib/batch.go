package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FncTransformer is a function that given a name and a slice of bye
// returns a Palletter (or eventually an error).
// It is used by CreateBatch function to transform the content of a file in a Palletter.
type FncTransformer func(string, []byte) (Paletter, error)

// CreateBatch reads a list of files and try to use fnc function to transform these files into
// Tilix schema, saving the result on outDir. In case of some errors, it propagates the error using the channel.
// files should have a list of file to iterate, with relative path.
// outDir can be the same directory where input files exist.
func CreateBatch(outDir string, files []string, fnc FncTransformer, ee chan error) int {
	fc := make(chan string)

	w := readAndParse(fnc, fc, ee)
	cont := writeFile(outDir, w, ee)

	for _, f := range files {
		fc <- f
	}
	close(fc)
	c := <-cont
	return c
}

// writeFile write a TilixColor struct inside a file. If some error occurs, it will propagate using e channel.
// It returns a chan of int where will be printed the number of files written without errors.
// After the result has been sent down the channel, it closes the channel.
func writeFile(dir string, w <-chan *TilixColor, e chan error) <-chan int {

	cont := make(chan int)
	c := 0

	go func() {

		for tc := range w {

			if tc == nil {
				e <- fmt.Errorf("nil schema received")
				continue
			}

			m, err := json.MarshalIndent(tc, "", "  ")
			if err != nil {
				e <- fmt.Errorf("unable to create json for schema '%s': %v", tc.Name, err)
				continue
			}

			f, err := os.Create(dir + "/" + tc.Name + ".json")
			if err != nil {
				e <- fmt.Errorf("unable to create the output file '%s': %v", tc.Name, err)
				continue
			}
			_, err = f.Write(m)
			if err != nil {
				e <- fmt.Errorf("unable to write file '%s': %v", tc.Name, err)
				continue
			}
			c++
			f.Close()

		}

		cont <- c
		close(cont)

	}()

	return cont

}

// readAndParse tries to read a file and using fnc convert it to a TilixColor struct.
// It returns a channel where results will be sent. After all results has been sent, it closes the channel.
func readAndParse(fnc FncTransformer, files <-chan string, e chan error) <-chan *TilixColor {

	w := make(chan *TilixColor)

	go func() {

		for file := range files {

			data, err := ioutil.ReadFile(file)
			if err != nil {
				e <- fmt.Errorf("unable to open the file '%s': %v", file, err)
				continue
			}
			_, fileName := filepath.Split(file)
			splits := strings.Split(fileName, ".")
			p, err := fnc(splits[0], data)
			if err != nil {
				e <- fmt.Errorf("unable to create a trasfomer for file '%s': %v", file, err)
				continue
			}
			tc, err := NewTilixColor(p)
			if err != nil {
				e <- fmt.Errorf("unable to create Tilix color scheme for file'%s': %v", file, err)
				continue
			}
			w <- tc
		}
		close(w)
	}()
	return w

}
