package color4tilix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FncTransformer func(string, []byte) (Paletter, error)

func CreateBatch(outDir string, files []string, fnc FncTransformer, ee chan error) (int, error) {
	fc := make(chan string)

	w := readAndParse(fnc, fc, ee)
	cont := writeFile(outDir, w, ee)

	for _, f := range files {
		fc <- f
	}
	close(fc)
	c := <-cont
	return c, nil
}

func writeFile(dir string, w <-chan *TilixColor, e chan error) <-chan int {

	cont := make(chan int)
	c := 0

	go func() {

		for tc := range w {

			m, err := json.MarshalIndent(tc, "", "  ")
			if err != nil {
				e <- fmt.Errorf("unable to create json for schema '%s': %v", tc.Name, err)
			}

			f, err := os.Create(dir + "/" + tc.Name + ".json")
			if err != nil {
				e <- fmt.Errorf("unable to create the output file '%s': %v", tc.Name, err)
			}
			_, err = f.Write(m)
			if err != nil {
				e <- fmt.Errorf("unable to write file '%s': %v", tc.Name, err)
			}
			c++
			f.Close()

		}

		cont <- c
	}()

	return cont

}

func readAndParse(fnc FncTransformer, files <-chan string, e chan error) <-chan *TilixColor {

	w := make(chan *TilixColor)

	go func() {

		for file := range files {

			data, err := ioutil.ReadFile(file)
			if err != nil {
				e <- fmt.Errorf("unable to open the file '%s': %v", file, err)
			}
			_, fileName := filepath.Split(file)
			splits := strings.Split(fileName, ".")
			//p, err := color4tilix.NewItermColorsTransformer(splits[0], data)
			p, err := fnc(splits[0], data)
			if err != nil {
				e <- fmt.Errorf("unable to create a trasfomer for file '%s': %v", file, err)
			}
			tc, err := NewTilixColor(p)
			if err != nil {
				e <- fmt.Errorf("unable to create Tilix color scheme for file'%s': %v", file, err)
			}
			w <- tc
		}
		close(w)
	}()
	return w

}
