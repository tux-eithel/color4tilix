package github.com/tux-eithel/color4tilix

// TilixColor represents is the struct for https://github.com/gnunn1/tilix/blob/master/source/gx/tilix/colorschemes.d
type TilixColor struct {
	Name                     string       `json:"name"`
	Comment                  string       `json:"comment"`
	UseThemeColor            bool         `json:"use-theme-color"`
	ForegroundColor          string       `json:"foreground-color,omitempty"`
	BackgroundColor          string       `json:"background-color,omitempty"`
	UseHighlightColor        bool         `json:"use-highlight-color,omitempty"`
	HighlightForegroundColor string       `json:"highlight-foreground-color,omitempty"`
	HighlightBackgroundColor string       `json:"highlight-background-color,omitempty"`
	UseCursorColor           bool         `json:"use-cursor-color,omitempty"`
	CursorForegroundColor    string       `json:"cursor-foreground-color,omitempty"`
	CursorBackgroundColor    string       `json:"cursor-background-color,omitempty"`
	UseBadgeColor            bool         `json:"use-badge-color,omitempty"`
	BadgeColor               string       `json:"badge-color,omitempty"`
	UseBoldColor             bool         `json:"use-bold-color,omitempty"`
	BoldColor                string       `json:"bold-color,omitempty"`
	Palette                  []color.RGBA `json:"palette"`
}
