package color4tilix

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func TestCreateBatch(t *testing.T) {

	dir, err := ioutil.TempDir("", "TestCreateBatch")
	if err != nil {
		t.Fatalf("unable to write tmp dir: %v", err)
	}
	defer os.RemoveAll(dir) // clean up

	f, err := ioutil.TempFile(dir, "*.itermcolors")
	if err != nil {
		t.Fatalf("unable to write tmp file: %v", err)
	}

	_, err = f.WriteString(realXml)
	if err != nil {
		t.Fatalf("unable to write content inside tmp file: %v", err)
	}
	name := f.Name()
	f.Close()

	fncEmptyPalette := func(n string, data []byte) (Paletter, error) {
		return &PaletterTest{}, nil
	}
	fncErrorPalette := func(n string, data []byte) (Paletter, error) {
		return nil, fmt.Errorf("error palette function")
	}
	fncNilPalette := func(n string, data []byte) (Paletter, error) {
		return nil, nil
	}

	fncWhitePalette := func(n string, data []byte) (Paletter, error) {
		allwhite := "all white"
		var white *color.RGBA
		var dark *color.RGBA
		if c, ok := color.RGBAModel.Convert(color.White).(color.RGBA); ok {
			white = &c
		}
		if c, ok := color.RGBAModel.Convert(color.Black).(color.RGBA); ok {
			dark = &c
		}
		return &PaletterTest{
			NameT:                     allwhite,
			CommentT:                  allwhite + " comment",
			UseThemeColorT:            false,
			ForegroundColorT:          white,
			BackgroundColorT:          white,
			UseHighlightColorT:        true,
			HighlightForegroundColorT: white,
			HighlightBackgroundColorT: white,
			UseCursorColorT:           true,
			CursorForegroundColorT:    white,
			CursorBackgroundColorT:    white,
			UseBadgeColorT:            true,
			BadgeColorT:               white,
			UseBoldColorT:             true,
			BoldColorT:                white,
			PaletteT:                  []*color.RGBA{white, dark},
		}, nil
	}

	type args struct {
		outDir string
		files  []string
		fnc    FncTransformer
	}
	tests := []struct {
		name   string
		args   args
		want   int
		errors int
	}{
		{
			name: "empty dir",
			args: args{
				outDir: dir,
				files:  []string{},
				fnc:    fncEmptyPalette,
			},
			want:   0,
			errors: 0,
		},
		{
			name: "non existing file",
			args: args{
				outDir: dir,
				files:  []string{"notAValidFile"},
				fnc:    fncEmptyPalette,
			},
			want:   0,
			errors: 1,
		},
		{
			name: "error palette function",
			args: args{
				outDir: dir,
				files:  []string{name},
				fnc:    fncErrorPalette,
			},
			want:   0,
			errors: 1,
		},
		{
			name: "palette with error",
			args: args{
				outDir: dir,
				files:  []string{name},
				fnc:    fncNilPalette,
			},
			want:   0,
			errors: 1,
		},
		{
			name: "simple struct",
			args: args{
				outDir: dir,
				files:  []string{name},
				fnc:    fncWhitePalette,
			},
			want:   1,
			errors: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			numErrors := 0
			listError := make([]error, 0)

			ee := make(chan error)
			wait := &sync.WaitGroup{}

			wait.Add(1)
			go func() {
				for e := range ee {
					numErrors++
					listError = append(listError, e)
				}
				wait.Done()

			}()

			if got := CreateBatch(tt.args.outDir, tt.args.files, tt.args.fnc, ee); got != tt.want {
				t.Errorf("CreateBatch() = %v, want %v", got, tt.want)
				return
			}
			close(ee)
			wait.Wait()

			if numErrors != tt.errors {
				t.Errorf("CreateBatch(): check errors => %v, want %v = erros %v", numErrors, tt.errors, listError)
				return
			}

		})
	}
}
