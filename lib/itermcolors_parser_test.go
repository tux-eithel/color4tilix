package lib

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

const xmlColors = `

	<key>Ansi 0 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 1 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 10 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 11 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 12 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 13 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 14 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 15 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 2 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 3 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 4 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 5 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 6 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 7 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 8 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Ansi 9 Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
`

const emptyPaletteXml = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Background Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Badge Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Bold Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Cursor Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Cursor Guide Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Cursor Text Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Foreground Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Link Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Selected Text Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Selection Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Tab Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Underline Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
</dict>
</plist>`

const realXml = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>` + xmlColors + `
	<key>Background Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Badge Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Bold Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Cursor Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Cursor Guide Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Cursor Text Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Foreground Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Link Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Selected Text Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Selection Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Tab Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
	<key>Underline Color</key>
	<dict>
		<key>Alpha Component</key>
		<real>1</real>
		<key>Blue Component</key>
		<real>1</real>
		<key>Color Space</key>
		<string>sRGB</string>
		<key>Green Component</key>
		<real>1</real>
		<key>Red Component</key>
		<real>1</real>
	</dict>
</dict>
</plist>`

func TestNewItermColorsTransformer(t *testing.T) {

	var white *color.RGBA
	//var dark *color.RGBA
	if c, ok := color.RGBAModel.Convert(color.White).(color.RGBA); ok {
		white = &c
	}
	/*
		if c, ok := color.RGBAModel.Convert(color.Black).(color.RGBA); ok {
			dark = &c
		}
	*/

	type args struct {
		name string
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *PaletterTest
		wantErr bool
	}{
		{
			name: "bad data",
			args: args{
				name: "file xml",
				data: []byte("hey, i'm a bad string"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "real data",
			args: args{
				name: "clovis color schema",
				data: []byte(realXml),
			},
			want: &PaletterTest{
				NameT:                     "clovis color schema",
				CommentT:                  fmt.Sprintf(descriptionItermSchema, "clovis color schema"),
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
				PaletteT:                  []*color.RGBA{white, white, white, white, white, white, white, white, white, white, white, white, white, white, white, white},
			},
			wantErr: false,
		},
		{
			name: "empty palette",
			args: args{
				name: "clovis without palette",
				data: []byte(emptyPaletteXml),
			},
			want: &PaletterTest{
				NameT:                     "clovis without palette",
				CommentT:                  fmt.Sprintf(descriptionItermSchema, "clovis without palette"),
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
				PaletteT:                  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewItermColorsTransformer(tt.args.name, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewItermColorsTransformer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && tt.wantErr {
				return
			}

			if got == nil && !tt.wantErr {
				t.Errorf("NewItermColorsTransformer() = %v, want %v", got, tt.want)
				return
			}

			if got.Name() != tt.want.Name() {
				t.Errorf("NewItermColorsTransformer(): wrong Name => %v, want %v", got.Name(), tt.want.Name())
				return
			}

			if got.Comment() != tt.want.Comment() {
				t.Errorf("NewItermColorsTransformer(): wrong Comment => %v, want %v", got.Comment(), tt.want.Comment())
				return
			}

			if got.UseThemeColor() != tt.want.UseThemeColor() {
				t.Errorf("NewItermColorsTransformer(): wrong UseThemeColor => %v, want %v", got.UseThemeColor(), tt.want.UseThemeColor())
				return
			}

			if !reflect.DeepEqual(got.ForegroundColor(), tt.want.ForegroundColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong ForegroundColor => %v, want %v", got.ForegroundColor(), tt.want.ForegroundColor())
				return
			}
			if !reflect.DeepEqual(got.BackgroundColor(), tt.want.BackgroundColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong BackgroundColor => %v, want %v", got.BackgroundColor(), tt.want.BackgroundColor())
				return
			}

			if got.UseHighlightColor() != tt.want.UseHighlightColor() {
				t.Errorf("NewItermColorsTransformer(): wrong UseHighlightColor => %v, want %v", got.UseHighlightColor(), tt.want.UseHighlightColor())
				return
			}
			if !reflect.DeepEqual(got.HighlightForegroundColor(), tt.want.HighlightForegroundColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong HighlightForegroundColor => %v, want %v", got.HighlightForegroundColor(), tt.want.HighlightForegroundColor())
				return
			}
			if !reflect.DeepEqual(got.HighlightBackgroundColor(), tt.want.HighlightBackgroundColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong HighlightBackgroundColor => %v, want %v", got.HighlightBackgroundColor(), tt.want.HighlightBackgroundColor())
				return
			}

			if got.UseCursorColor() != tt.want.UseCursorColor() {
				t.Errorf("NewItermColorsTransformer(): wrong UseCursorColor => %v, want %v", got.UseCursorColor(), tt.want.UseCursorColor())
				return
			}
			if !reflect.DeepEqual(got.CursorForegroundColor(), tt.want.CursorForegroundColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong CursorForegroundColor => %v, want %v", got.CursorForegroundColor(), tt.want.CursorForegroundColor())
				return
			}
			if !reflect.DeepEqual(got.CursorBackgroundColor(), tt.want.CursorBackgroundColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong CursorBackgroundColor => %v, want %v", got.CursorBackgroundColor(), tt.want.CursorBackgroundColor())
				return
			}

			if got.UseBadgeColor() != tt.want.UseBadgeColor() {
				t.Errorf("NewItermColorsTransformer(): wrong UseBadgeColor => %v, want %v", got.UseBadgeColor(), tt.want.UseBadgeColor())
				return
			}
			if !reflect.DeepEqual(got.BadgeColor(), tt.want.BadgeColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong BadgeColor => %v, want %v", got.BadgeColor(), tt.want.BadgeColor())
				return
			}

			if got.UseBoldColor() != tt.want.UseBoldColor() {
				t.Errorf("NewItermColorsTransformer(): wrong UseBoldColor => %v, want %v", got.UseBoldColor(), tt.want.UseBoldColor())
				return
			}
			if !reflect.DeepEqual(got.BoldColor(), tt.want.BoldColor()) {
				t.Errorf("NewItermColorsTransformer(): wrong BoldColor => %v, want %v", got.BoldColor(), tt.want.BoldColor())
				return
			}

			if !reflect.DeepEqual(got.Palette(), tt.want.Palette()) {
				t.Errorf("NewItermColorsTransformer(): wrong Palette => %v, want %v", got.Palette(), tt.want.Palette())
				return
			}

		})
	}
}
