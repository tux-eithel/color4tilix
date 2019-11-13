package color4tilix

import (
	"encoding/xml"
	"fmt"
	"image/color"
	"strings"
)

// MainXml is the main xml node of a .itermcolors file.
type MainXml struct {
	XMLName xml.Name `xml:"plist"`
	Data    Dict     `xml:"dict"`
}

// Dict is the direct child of "dict" node on the xml.
// The right interpretation for this struct is the follow:
// given an index, the Keys[index] is the property name, the DictValues[index] is the "color" value.
type Dict struct {
	Keys       []string     `xml:"key"`
	DictValues []DictValues `xml:"dict"`
}

// DictValues is the actual color associate to a property.
// .itermolors files save the color in 3 (or 4) separate values:
//   - Blue Component
//   - Green Component
//   - Red Component
//   - Alpha Component
type DictValues struct {
	Keys   []string  `xml:"key"`
	Values []float64 `xml:"real"`
}

// ItermColor is a simple struct to better handle Dict with DictValues.
// It's used to merge Dict and DictValues structs in a single, user friendly, struct.
type ItermColor struct {
	Name  string
	Color color.RGBA
}

// descriptionItermSchema is the default comment string.
// .itermcolors files doesn't support comments inside, so we set un a default.
const descriptionItermSchema = "Automatic conversion for theme '%s' made by color4tilix"

// ItermSchema should represents the color schema used by iTerm2.
// To elaborate this, I've used this as reference https://raw.githubusercontent.com/Clovis-team/clovis-open-code-extracts/master/utils/Clovis-iTerm2-Color-Scheme.itermcolors .
type ItermSchema struct {
	name        string
	description string
	objs        map[string]color.RGBA
}

func (it *ItermSchema) Name() string {
	return it.name
}
func (it *ItermSchema) Comment() string {
	return it.description
}
func (it *ItermSchema) UseThemeColor() bool {
	return false // It's better to set it as false, otherwise Tilix doesn't display the theme's foreground/background colors.
}

func (it *ItermSchema) ForegroundColor() *color.RGBA { return colorOrNil(it, "Foreground Color") }
func (it *ItermSchema) BackgroundColor() *color.RGBA { return colorOrNil(it, "Background Color") }
func (it *ItermSchema) UseHighlightColor() bool      { return colorOrNil(it, "Selection Color") != nil }
func (it *ItermSchema) HighlightForegroundColor() *color.RGBA {
	return colorOrNil(it, "Selected Text Color")
}
func (it *ItermSchema) HighlightBackgroundColor() *color.RGBA {
	return colorOrNil(it, "Selection Color")
}
func (it *ItermSchema) UseCursorColor() bool               { return colorOrNil(it, "Cursor Color") != nil }
func (it *ItermSchema) CursorForegroundColor() *color.RGBA { return colorOrNil(it, "Cursor Text Color") }
func (it *ItermSchema) CursorBackgroundColor() *color.RGBA { return colorOrNil(it, "Cursor Color") }
func (it *ItermSchema) UseBadgeColor() bool                { return colorOrNil(it, "Badge Color") != nil }
func (it *ItermSchema) BadgeColor() *color.RGBA            { return colorOrNil(it, "Badge Color") }
func (it *ItermSchema) UseBoldColor() bool                 { return colorOrNil(it, "Bold Color") != nil }
func (it *ItermSchema) BoldColor() *color.RGBA             { return colorOrNil(it, "Bold Color") }
func (it *ItermSchema) Palette() []*color.RGBA {
	p := make([]*color.RGBA, 0, 16)
	baseName := "Ansi %d Color"
	for i := 0; i < 16; i++ {
		c := colorOrNil(it, fmt.Sprintf(baseName, i))
		if c == nil {
			return nil
		}
		p = append(p, c)
	}
	return p
}

// NewItermColorsTransformer creates a ItermSchema.
// It accepts a name, which will be theme name, and a slice of byte as a xml data to parse.
func NewItermColorsTransformer(name string, data []byte) (*ItermSchema, error) {

	x := MainXml{}

	err := xml.Unmarshal(data, &x)
	if err != nil {
		return nil, fmt.Errorf("unable to parse xml: %s", err)
	}

	allColors := make([]ItermColor, 0, len(x.Data.Keys))

	for k, v := range x.Data.Keys {
		tmp := ItermColor{
			Name:  v,
			Color: color.RGBA{},
		}

		k1 := 0 // sometimes there are string next to "key", we increment separatly only if recognize the "key"
		for _, v1 := range x.Data.DictValues[k].Keys {
			switch {
			case strings.Contains(v1, "Blue"):
				tmp.Color.B = uint8(x.Data.DictValues[k].Values[k1] * 255)
				k1++
			case strings.Contains(v1, "Red"):
				tmp.Color.R = uint8(x.Data.DictValues[k].Values[k1] * 255)
				k1++
			case strings.Contains(v1, "Green"):
				tmp.Color.G = uint8(x.Data.DictValues[k].Values[k1] * 255)
				k1++
			case strings.Contains(v1, "Alpha"):
				tmp.Color.A = uint8(x.Data.DictValues[k].Values[k1] * 255)
				k1++
			}

		}
		allColors = append(allColors, tmp)

	}

	it := &ItermSchema{
		name:        name,
		description: fmt.Sprintf(descriptionItermSchema, name),
		objs:        make(map[string]color.RGBA),
	}
	for _, v := range allColors {
		it.objs[v.Name] = v.Color
	}

	return it, nil
}

// colorOrNil is an utility function to check if a property is available and
// if it is present, it returns the color.RGBA associate.
func colorOrNil(it *ItermSchema, k string) *color.RGBA {
	v, ok := it.objs[k]
	if !ok {
		return nil
	}
	return &v
}
