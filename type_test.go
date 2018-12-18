package color4tilix

import (
	"image/color"
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

	emptyName := "empty"
	allwhite := "all white"
	whiteColorString := "#ffffff"
	darkColorString := "#000000"
	emptyTilixColor := &TilixColor{Name: emptyName}
	allWhiteTilixColor := &TilixColor{
		Name:                     allwhite,
		Comment:                  allwhite + " comment",
		UseThemeColor:            false,
		ForegroundColor:          whiteColorString,
		BackgroundColor:          whiteColorString,
		UseHighlightColor:        true,
		HighlightForegroundColor: whiteColorString,
		HighlightBackgroundColor: whiteColorString,
		UseCursorColor:           true,
		CursorForegroundColor:    whiteColorString,
		CursorBackgroundColor:    whiteColorString,
		UseBadgeColor:            true,
		BadgeColor:               whiteColorString,
		UseBoldColor:             true,
		BoldColor:                whiteColorString,
		Palette:                  []string{whiteColorString, darkColorString},
	}

	var white *color.RGBA
	var dark *color.RGBA
	if c, ok := color.RGBAModel.Convert(color.White).(color.RGBA); ok {
		white = &c
	}
	if c, ok := color.RGBAModel.Convert(color.Black).(color.RGBA); ok {
		dark = &c
	}

	tests := []struct {
		name    string
		args    Paletter
		want    *TilixColor
		wantErr bool
	}{
		{
			name:    "empty parser",
			args:    nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty name",
			args:    &PaletterTest{PaletteT: make([]*color.RGBA, 0)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty palette",
			args:    &PaletterTest{},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "valid but empty",
			args:    &PaletterTest{NameT: emptyName, PaletteT: make([]*color.RGBA, 0)},
			want:    emptyTilixColor,
			wantErr: false,
		},
		{
			name:    "wrong color",
			args:    &PaletterTest{NameT: emptyName, PaletteT: []*color.RGBA{dark, nil}},
			want:    &TilixColor{Name: emptyName},
			wantErr: true,
		},
		{
			name: allwhite,
			args: &PaletterTest{
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
			},
			want:    allWhiteTilixColor,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTilixColor(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTilixColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.wantErr {
				return
			}

			if got == nil && !tt.wantErr {
				t.Errorf("NewTilixColor() = %v, want %v", got, tt.want)
				return
			}

			if got.Name != tt.want.Name {
				t.Errorf("NewTilixColor(): wrong Name => %v, want %v", got.Name, tt.want.Name)
				return
			}
			if got.Comment != tt.want.Comment {
				t.Errorf("NewTilixColor(): wrong Comment => %v, want %v", got.Comment, tt.want.Comment)
				return
			}

			if got.UseThemeColor != tt.want.UseThemeColor {
				t.Errorf("NewTilixColor() wrong UseThemeColor => %v, want %v", got.UseThemeColor, tt.want.UseThemeColor)
				return
			}
			if got.ForegroundColor != tt.want.ForegroundColor {
				t.Errorf("NewTilixColor() wrong ForegroundColor => %v, want %v", got.ForegroundColor, tt.want.ForegroundColor)
				return
			}
			if got.BackgroundColor != tt.want.BackgroundColor {
				t.Errorf("NewTilixColor() wrong BackgroundColor => %v, want %v", got.BackgroundColor, tt.want.BackgroundColor)
				return
			}

			if got.UseHighlightColor != tt.want.UseHighlightColor {
				t.Errorf("NewTilixColor() wrong UseHighlightColor => %v, want %v", got.UseHighlightColor, tt.want.UseHighlightColor)
				return
			}
			if got.HighlightForegroundColor != tt.want.HighlightForegroundColor {
				t.Errorf("NewTilixColor() wrong HighlightForegroundColor => %v, want %v", got.HighlightForegroundColor, tt.want.HighlightForegroundColor)
				return
			}
			if got.HighlightBackgroundColor != tt.want.HighlightBackgroundColor {
				t.Errorf("NewTilixColor() wrong HighlightBackgroundColor => %v, want %v", got.HighlightBackgroundColor, tt.want.HighlightBackgroundColor)
				return
			}

			if got.UseCursorColor != tt.want.UseCursorColor {
				t.Errorf("NewTilixColor() wrong UseCursorColor => %v, want %v", got.UseCursorColor, tt.want.UseCursorColor)
				return
			}
			if got.CursorForegroundColor != tt.want.CursorForegroundColor {
				t.Errorf("NewTilixColor() wrong CursorForegroundColor => %v, want %v", got.CursorForegroundColor, tt.want.CursorForegroundColor)
				return
			}
			if got.CursorBackgroundColor != tt.want.CursorBackgroundColor {
				t.Errorf("NewTilixColor() wrong CursorBackgroundColor => %v, want %v", got.CursorBackgroundColor, tt.want.CursorBackgroundColor)
				return
			}

			if got.UseBadgeColor != tt.want.UseBadgeColor {
				t.Errorf("NewTilixColor() wrong UseBadgeColor => %v, want %v", got.UseBadgeColor, tt.want.UseBadgeColor)
				return
			}
			if got.BadgeColor != tt.want.BadgeColor {
				t.Errorf("NewTilixColor() wrong BadgeColor => %v, want %v", got.BadgeColor, tt.want.BadgeColor)
				return
			}

			if got.UseBoldColor != tt.want.UseBoldColor {
				t.Errorf("NewTilixColor() wrong UseBoldColor => %v, want %v", got.UseBoldColor, tt.want.UseBoldColor)
				return
			}
			if got.UseBoldColor != tt.want.UseBoldColor {
				t.Errorf("NewTilixColor() wrong UseBoldColor => %v, want %v", got.UseBoldColor, tt.want.UseBoldColor)
				return
			}

			if len(got.Palette) != len(tt.want.Palette) {
				t.Errorf("NewTilixColor() wrong Palette length => %v, want %v", len(got.Palette), len(tt.want.Palette))
				return
			}

			for k, v := range got.Palette {
				if v != tt.want.Palette[k] {
					t.Errorf("NewTilixColor() wrong Palette color at %d position => %v, want %v", k, v, tt.want.Palette[k])
					return
				}
			}

		})
	}
}
