// Package color4tilix is a library to help to translate terminal color schema to Tilix color schema.
// It provides and maintains the struct necessary to create the right json for Tilix color schema.
//
// It defines an interface, Paletter, which can be used to define a new struct to translate an existing schema
// to a Tilix json schema. The struct, will be used by NewTilixColor function which is in charge to build a correct TilixColor struct.
// As example, the library provides ItermSchema struct used to translate .itermcolors files (iTerm color schema) into Tilix color schema.
//
// The library provides also a function, CreateBatch, which can be used to read files, translate and save it.
// Generally a conversion process requires some input files, an output dir and a set of instruction to translate input files from one format to another one.
// For this reasons, the library has the FncTransformer type, which is used by CreateBatch function.
//
// If you'd like to transform another terminal schema into Tilix color schemes, (and you're fine with CreateBatch workflow) you should do:
//  - define a new struct which implements Paletter interface
//  - provide a compatible FncTransformer function
package color4tilix

import (
	"fmt"
	"image/color"
)

// TilixColor represents is the struct for https://github.com/gnunn1/tilix/blob/master/source/gx/tilix/colorschemes.d .
type TilixColor struct {
	Name                     string   `json:"name"`
	Comment                  string   `json:"comment"`
	UseThemeColor            bool     `json:"use-theme-colors"`
	ForegroundColor          string   `json:"foreground-color,omitempty"`
	BackgroundColor          string   `json:"background-color,omitempty"`
	UseHighlightColor        bool     `json:"use-highlight-color,omitempty"`
	HighlightForegroundColor string   `json:"highlight-foreground-color,omitempty"`
	HighlightBackgroundColor string   `json:"highlight-background-color,omitempty"`
	UseCursorColor           bool     `json:"use-cursor-color,omitempty"`
	CursorForegroundColor    string   `json:"cursor-foreground-color,omitempty"`
	CursorBackgroundColor    string   `json:"cursor-background-color,omitempty"`
	UseBadgeColor            bool     `json:"use-badge-color,omitempty"`
	BadgeColor               string   `json:"badge-color,omitempty"`
	UseBoldColor             bool     `json:"use-bold-color,omitempty"`
	BoldColor                string   `json:"bold-color,omitempty"`
	Palette                  []string `json:"palette"`
}

// Paletter is an interface for parsers.
// It's a big interface, but in this way we can enforce the necessary fields for the json Tilix schema,
// delegating the actual implementation to the parser (color schemas between terminal are pretty vary).
type Paletter interface {
	Name() string
	Comment() string
	UseThemeColor() bool
	ForegroundColor() *color.RGBA
	BackgroundColor() *color.RGBA
	UseHighlightColor() bool
	HighlightForegroundColor() *color.RGBA
	HighlightBackgroundColor() *color.RGBA
	UseCursorColor() bool
	CursorForegroundColor() *color.RGBA
	CursorBackgroundColor() *color.RGBA
	UseBadgeColor() bool
	BadgeColor() *color.RGBA
	UseBoldColor() bool
	BoldColor() *color.RGBA
	Palette() []*color.RGBA
}

// hexOrEmpty return the hex form of the color if is not nil,
// otherwise returns an empty string.
func hexOrEmpty(c *color.RGBA) string {
	if c != nil {
		return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	}
	return ""
}

// NewTilixColor creates a new TilixColor struct starting from a parser.
// Its main role is to guarantee the correctness of the struct.
// This struct can be marshal to a json.
func NewTilixColor(parser Paletter) (*TilixColor, error) {

	if parser == nil {
		return nil, fmt.Errorf("empty parser given")
	}

	tc := TilixColor{Name: parser.Name(), Comment: parser.Comment()}

	p := parser.Palette()

	if p == nil {
		return nil, fmt.Errorf("no palette defined")
	}
	if tc.Name == "" {
		return nil, fmt.Errorf("no Name defined")
	}

	tc.UseThemeColor = parser.UseThemeColor()

	tc.ForegroundColor = hexOrEmpty(parser.ForegroundColor())
	tc.BackgroundColor = hexOrEmpty(parser.BackgroundColor())

	tc.UseHighlightColor = parser.UseHighlightColor()
	tc.HighlightForegroundColor = hexOrEmpty(parser.HighlightForegroundColor())
	tc.HighlightBackgroundColor = hexOrEmpty(parser.HighlightBackgroundColor())

	tc.UseCursorColor = parser.UseCursorColor()
	tc.CursorForegroundColor = hexOrEmpty(parser.CursorForegroundColor())
	tc.CursorBackgroundColor = hexOrEmpty(parser.CursorBackgroundColor())

	tc.UseBadgeColor = parser.UseBadgeColor()
	tc.BadgeColor = hexOrEmpty(parser.BadgeColor())

	tc.UseBoldColor = parser.UseBoldColor()
	tc.BoldColor = hexOrEmpty(parser.BoldColor())

	tc.Palette = make([]string, len(p))

	var hex string
	for k, v := range p {
		hex = hexOrEmpty(v)
		if hex == "" {
			return nil, fmt.Errorf("the color for schema '%s' at position %d seems empty", tc.Name, k)
		}
		tc.Palette[k] = hex
	}

	return &tc, nil
}
