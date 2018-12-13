package color4tilix

import (
	"image/color"
	"reflect"
	"testing"
)

type PaletterTest struct {
	NameT                     string
	CommentT                  string
	UseThemeColorT            bool
	ForegroundColorT          *color.RGBA
	BackgroundColorT          *color.RGBA
	UseHighlightColorT        bool
	HighlightForegroundColorT *color.RGBA
	HighlightBackgroundColorT *color.RGBA
	UseCursorColorT           bool
	CursorForegroundColorT    *color.RGBA
	CursorBackgroundColorT    *color.RGBA
	UseBadgeColorT            bool
	BadgeColorT               *color.RGBA
	UseBoldColorT             bool
	BoldColorT                *color.RGBA
	PaletteT                  []*color.RGBA
}

func (it *PaletterTest) Name() string {
	return it.NameT
}
func (it *PaletterTest) Comment() string {
	return it.CommentT
}
func (it *PaletterTest) UseThemeColor() bool {
	return it.UseThemeColorT
}

func (it *PaletterTest) ForegroundColor() *color.RGBA          { return it.ForegroundColorT }
func (it *PaletterTest) BackgroundColor() *color.RGBA          { return it.BackgroundColorT }
func (it *PaletterTest) UseHighlightColor() bool               { return it.UseHighlightColorT }
func (it *PaletterTest) HighlightForegroundColor() *color.RGBA { return it.HighlightForegroundColorT }
func (it *PaletterTest) HighlightBackgroundColor() *color.RGBA { return it.HighlightForegroundColorT }
func (it *PaletterTest) UseCursorColor() bool                  { return it.UseCursorColorT }
func (it *PaletterTest) CursorForegroundColor() *color.RGBA    { return it.CursorForegroundColorT }
func (it *PaletterTest) CursorBackgroundColor() *color.RGBA    { return it.CursorBackgroundColorT }
func (it *PaletterTest) UseBadgeColor() bool                   { return it.UseBadgeColorT }
func (it *PaletterTest) BadgeColor() *color.RGBA               { return it.BadgeColorT }
func (it *PaletterTest) UseBoldColor() bool                    { return it.UseBoldColorT }
func (it *PaletterTest) BoldColor() *color.RGBA                { return it.BoldColorT }
func (it *PaletterTest) Palette() []*color.RGBA                { return it.PaletteT }

func TestNewTilixColor(t *testing.T) {

	tests := []struct {
		name        string
		args        Paletter
		want        *TilixColor
		wantErr     bool
		stringError string
	}{
		{
			name:        "empty parser",
			args:        nil,
			want:        nil,
			wantErr:     true,
			stringError: "empty parser given",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTilixColor(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTilixColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && tt.wantErr && err.Error() != tt.stringError {
				t.Errorf("NewTilixColor() error = %v, stringError %v", err, tt.stringError)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTilixColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
